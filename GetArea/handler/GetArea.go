package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"
	"time"

	log "github.com/micro/go-micro/v2/logger"

	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
	models "sss/GetArea/models"
	GA "sss/GetArea/proto/GetArea"
	utils "sss/GetArea/utils"
)

type GetArea struct{}


func (c *GetArea) GetArea(ctx context.Context, req *GA.Request, out *GA.Response) error {
	//req := c.c.NewRequest(c.name, "GetArea.GetArea", in)
	out.Error = utils.RECODE_OK
	out.Errmsg = utils.RecodeText(out.Error)
	//err := c.c.Call(ctx, req, out, opts...)
	//if err != nil {
	//	return nil, err
	//}

		//1.从缓存中获取数据
		redisConf := map[string]string{
			"key":utils.G_server_name,
			//127.0.0.1:6379
			"conn":utils.G_redis_addr+":"+utils.G_redis_port,
			"dbNum":utils.G_redis_dbnum,
		}
		beego.Info(redisConf)
		redisConfJson,err := json.Marshal(redisConf)
		bm,err := cache.NewCache("redis",string(redisConfJson))
		if err != nil{
			beego.Info("连接失败")
			out.Error = utils.RECODE_DBERR
			out.Errmsg = utils.RecodeText(out.Error)
		}
		//获取数据需要定制一个key，用于缓存或者读取“area_info”
		areaValue:=bm.Get("area_info")
		if areaValue != nil{
			beego.Info("获取到数据")
			areaMap :=[]map[string]interface{}{}
			err := json.Unmarshal(areaValue.([]byte),&areaMap)
			if err != nil{
				return err
			}

			for _, value := range areaMap {
				temp := GA.Response_Areas{Aid:int32(value["aid"].(float64)),Aname: value["aname"].(string)}
				out.Data  = append(out.Data,&temp)
			}
			return nil
		}

		//2.如果没有数据就从mysql中查找

		//4、把数据发送给前端

	o := orm.NewOrm()
	var area []models.Area
	qs :=o.QueryTable("area")
	num ,err :=qs.All(&area)
	if err != nil{
		beego.Info("查询失败")
		out.Error = utils.RECODE_DATAERR
		out.Errmsg = utils.RecodeText(out.Error)
		return err
	}else if num == 0{
		beego.Info("没有数据")
		out.Error = utils.RECODE_NODATA
		out.Errmsg = utils.RecodeText(out.Error)
		return errors.New(out.Errmsg)
	}
	//3.将查找的数据存到缓存中
	areaJson,_ := json.Marshal(area)
	err = bm.Put("area_info",areaJson,3600*time.Second)
	if err != nil{
		beego.Info("缓存失败")
		beego.Info(err)

	}
	for _, value := range area {
		temp := GA.Response_Areas{Aid:int32(value.Id),Aname: value.Name}
		out.Data  = append(out.Data,&temp)
	}

	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *GetArea) Stream(ctx context.Context, req *GA.StreamingRequest, stream GA.GetArea_StreamStream) error {
	log.Infof("Received GetArea.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&GA.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *GetArea) PingPong(ctx context.Context, stream GA.GetArea_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&GA.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
