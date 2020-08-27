package handler

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/afocus/captcha"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
	"image/color"
	"math/rand"
	"reflect"
	models "sss/GetImageCd/models"
	"sss/GetImageCd/utils"
	"strconv"
	"submail_go_sdk/submail/sms"
	"time"

	log "github.com/micro/go-micro/v2/logger"

	GI "sss/GetImageCd/proto/GetImageCd"
)

type GetImageCd struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *GetImageCd) GetImageCode(ctx context.Context, req *GI.Request, rsp *GI.Response) error {
	beego.Info("获取验证码图片 GetImageCd /api/v1.0/imagecode/:uuid")

	/*生成验证码图片*/
	//创建图片句柄
	cap := captcha.New()

	//设置字体
	if err := cap.SetFont("comic.ttf"); err != nil {
		panic(err.Error())
	}


	//设置图片大小
	cap.SetSize(91, 41)
	//设置干扰强度
	cap.SetDisturbance(captcha.NORMAL)
	//设置前景色
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	//设置背景色
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})

	//生存随即的验证码图片
	img, str := cap.Create(4, captcha.NUM)
	//缓存图片
	bm,_ := connectRedis()
	bm.Put(req.Uuid,str,time.Second*300)

	//图片解引用
	img1 :=*img
	img2 :=*img1.RGBA
	//返回错误信息
	rsp.Error= utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Error)
	//返回图片拆分
	rsp.Pix = []byte(img2.Pix)
	rsp.Stride = int64(img2.Stride)

	rsp.Max = &GI.Response_Point{X:int64(img2.Rect.Max.X),Y:int64(img2.Rect.Max.Y)}
	rsp.Min = &GI.Response_Point{X:int64(img2.Rect.Min.X),Y:int64(img2.Rect.Min.Y)}
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *GetImageCd) Stream(ctx context.Context, req *GI.StreamingRequest, stream GI.GetImageCd_StreamStream) error {
	log.Infof("Received GetImageCd.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&GI.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *GetImageCd) PingPong(ctx context.Context, stream GI.GetImageCd_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&GI.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

func connectRedis() (cache.Cache,error){

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
		return bm,err
	}
	return bm,nil
}

func (e *GetImageCd) GetSmsCode(ctx context.Context, req *GI.SmsRequest, rsp *GI.SmsResponse) error {
	beego.Info("获取短信验证码 GetImageCd.GetSmscd api/v1.0/smscode/:mobile ")

	//初始化返回值
	rsp.Error = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Error)

	/*验证手机号是否存在*/
	//创建数据库orm句柄
	o:=orm.NewOrm()
	//使用手机号作为查询条件
	user := models.User{Mobile:req.Mobile}

	err := o.Read(&user)
	//如果不报错就说明查找到了
	//查找到就说明手机号存在
	if err==nil{

		beego.Info("用户以存在")
		rsp.Error = utils.RECODE_MOBILEERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}

	/*验证图片验证码是否正确*/
	//连接redis
	//配置缓存参数
	redis_conf := map[string]string{
		"key":utils.G_server_name,
		//127.0.0.1:6379
		"conn":utils.G_redis_addr+":"+utils.G_redis_port,
		"dbNum":utils.G_redis_dbnum,
	}
	beego.Info(redis_conf)

	//将map进行转化成为json
	redis_conf_js,_ :=json.Marshal(redis_conf)

	//创建redis句柄
	bm ,err :=cache.NewCache("redis",string(redis_conf_js))
	if err!=nil{
		beego.Info("redis连接失败",err)
		rsp.Error = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}

	value:=bm.Get(req.Uuid)
	if value ==nil{
		beego.Info("redis获取失败",err)
		rsp.Error = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}
	//reflect.TypeOf(value)会返回当前数据的变量类型
	beego.Info(reflect.TypeOf(value),value)
	//格式转换
	value_str,_:= redis.String(value,nil)

	if value_str != req.Imagestr{
		beego.Info("数据不匹配 图片验证码值错误")
		rsp.Error = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}

	//创建随机数
	r :=rand.New(rand.NewSource(time.Now().UnixNano()))
	size := r.Intn(9999)+1001
	beego.Info("验证码",size)
	SMSSend(strconv.Itoa(size))
	err=bm.Put(req.Mobile,size,time.Second*300)
	if err!=nil {
		beego.Info("redis创建失败",err)
		rsp.Error = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}

	return nil
}


func (e *GetImageCd) PostRet(ctx context.Context, req *GI.PostRetRequest, rsp *GI.PostRetResponse) error {
	beego.Info("注册 GetImageCd.PostRet ")

	//初始化返回值
	rsp.Error = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(rsp.Error)

	//redis操作
	//配置缓存参数
	redis_conf := map[string]string{
		"key":utils.G_server_name,
		//127.0.0.1:6379
		"conn":utils.G_redis_addr+":"+utils.G_redis_port,
		"dbNum":utils.G_redis_dbnum,
	}
	beego.Info(redis_conf)

	//将map进行转化成为json
	redis_conf_js,_ :=json.Marshal(redis_conf)

	//创建redis句柄
	bm ,err :=cache.NewCache("redis",string(redis_conf_js))
	if err!=nil{
		beego.Info("redis连接失败",err)
		rsp.Error = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}

	//通过手机号获取到短信验证码
	sms_code:=bm.Get(req.Mobile)
	if  sms_code ==nil {
		beego.Info("获取数据失败")
		rsp.Error = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}

	//短信验证码对比
	sms_code_str,_:=redis.String(sms_code,nil)

	if sms_code_str != req.SmsCode{
		beego.Info("短信验证码错误")
		rsp.Error = utils.RECODE_SMSERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}

	/*将数据存入数据库*/
	o:=orm.NewOrm()
	user:=models.User{Mobile:req.Mobile,Password_hash: Md5String(req.Password) ,Name:req.Mobile}

	id ,err:= o.Insert(&user)
	if err!=nil{
		beego.Info("注册数据失败")
		rsp.Error = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(rsp.Error)
		return nil
	}
	beego.Info("user_id",id)

	/*创建sessionid  （唯一的随即码）*/
	sessionid:=Md5String(req.Mobile+req.Password)

	rsp.SessionId = sessionid

	/*以sessionid为key的一部分创建session*/
	//name //名字暂时使用手机号
	bm.Put(sessionid+"name",user.Mobile,time.Second*3600)
	//user_id
	bm.Put(sessionid+"user_id",id,time.Second*3600)
	//手机号
	bm.Put(sessionid+"mobile",user.Mobile,time.Second*3600)


	return nil
}


//加密函数
func Md5String(s string )string{
	//创建1个md5对象
	h:=md5.New()
	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil))
}

func SMSSend(code string){
	// SMS 短信服务配置 appid & appkey 请前往：https://www.mysubmail.com/chs/sms/apps 获取
	config := make(map[string]string)
	config["appid"]="29672"
	config["appkey"]="89d90165cbea8cae80137d7584179bdb"
	// SMS 数字签名模式 normal or md5 or sha1 ,normal = 明文appkey鉴权 ，md5 和 sha1 为数字签名鉴权模式
	config["signType"]="normal"

	//创建 短信 Send 接口
	submail := sms.CreateSend(config)

	//设置联系人 手机号码
	submail.SetTo("18320857647")

	//设置短信正文，请注意：国内短信需要强制添加短信签名，并且需要使用全角大括号 “【签名】”标识，并放在短信正文的最前面
	submail.SetContent("【SUBMAIL】您的验证码是："+code+"，请在30分钟输入")

	//执行 Send 方法发送短信
	send := submail.Send()
	fmt.Println("短信 Send 接口:",send)
}

