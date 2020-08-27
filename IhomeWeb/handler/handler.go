package handler

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/afocus/captcha"
	"github.com/julienschmidt/httprouter"
	"image"
	"image/png"
	"net/http"
	"regexp"
	"sss/IhomeWeb/utils"
	"sss/IhomeWeb/models"
	"time"

	"github.com/micro/go-micro/v2/client"
	IhomeWeb "path/to/service/proto/IhomeWeb"
	GA "path/to/service/proto/GetArea"
	GI "path/to/service/proto/GetImageCd"
)

func IhomeWebCall(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	IhomeWebClient := IhomeWeb.NewUserService("go.micro.service.IhomeWeb", client.DefaultClient)
	rsp, err := IhomeWebClient.Call(context.TODO(), &IhomeWeb.Request{
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}


func GetArea(w http.ResponseWriter, r *http.Request,_ httprouter.Params){
	beego.Info("GetArea call")
	// decode the incoming request as json
	//var request map[string]interface{}
	//if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	//	http.Error(w, err.Error(), 500)
	//	return
	//}

	// call the backend service
	GAClient := GA.NewGetAreaService("go.micro.service.GetArea", client.DefaultClient)
	rsp, err := GAClient.GetArea(context.TODO(), &GA.Request{})
	if err != nil {
		beego.Info(err)
		http.Error(w, err.Error(), 500)
		return
	}

	//返回数据
	areaList := []models.Area{}
	for _, value := range rsp.Data {
		temp := models.Area{Id:int(value.Aid),Name:value.Aname}
		areaList = append(areaList, temp)
	}
	// we want to augment the response
	response := map[string]interface{}{
		"errno": rsp.Error,
		"errmsg": rsp.Errmsg,
		"data":areaList,
	}
	w.Header().Set("Content-Type","application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func GetIndex(w http.ResponseWriter, r *http.Request,_ httprouter.Params){

	beego.Info("GetIndex call")
	// we want to augment the response
	response := map[string]interface{}{
		"errno": utils.RECODE_OK,
		"errmsg": utils.RecodeText(utils.RECODE_OK),
	}
	w.Header().Set("Content-Type","application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func GetSession(w http.ResponseWriter, r *http.Request,_ httprouter.Params){
	// we want to augment the response
	beego.Info("GetSession call")
	response := map[string]interface{}{
		"errno": utils.RECODE_LOGINERR,
		"errmsg": utils.RecodeText(utils.RECODE_LOGINERR),
	}
	w.Header().Set("Content-Type","application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

//获取验证码图片
func GetImageCd(w http.ResponseWriter, r *http.Request,ps httprouter.Params){
	// we want to augment the response
	beego.Info("GetImageCd call")

	//获取uuid
	uudi := ps.ByName("uuid")

	GIClient := GI.NewGetImageCdService("go.micro.service.GetImageCd", client.DefaultClient)
	rsp, err := GIClient.GetImageCode(context.TODO(), &GI.Request{
		Uuid:uudi,
	})

	if err != nil {
		beego.Info(err)
		http.Error(w, err.Error(), 500)
		return
	}

	image := rspToImage(rsp)
	png.Encode(w,image)



	//response := map[string]interface{}{
	//	"errno": utils.RECODE_LOGINERR,
	//	"errmsg": utils.RecodeText(utils.RECODE_LOGINERR),
	//}
	//w.Header().Set("Content-Type","application/json")
	//// encode and write the response as json
	//if err := json.NewEncoder(w).Encode(response); err != nil {
	//	http.Error(w, err.Error(), 500)
	//	return
	//}
}

//获取验短信验证码
func GetSmscd(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {
	beego.Info("获取短信验证码 GetSmscd api/v1.0/smscode/:mobile ")
	//通过传入参数URL下Query获取前端的在url里的带参
	//beego.Info(r.URL.Query())
	//map[text:[9346] id:[474494b0-18eb-4eb7-9e68-a5ecf3c8317b]]
	//获取参数
	test:=r.URL.Query()["text"][0]
	id:=r.URL.Query()["id"][0]
	mobile:=ps.ByName("mobile")


	//通过正则进行手机号的判断
	//通过条件判断字符串是否匹配规则 返回正确或失败
	bl :=verifyMobileNum(mobile)
	//如果手机号不匹配那就直接返回错误不调用服务
	if bl == false{
		// 创建返回数据的map
		response := map[string]interface{}{
			"error": utils.RECODE_MOBILEERR ,
			"errmsg": utils.RecodeText(utils.RECODE_MOBILEERR),
		}

		//设置返回数据的格式
		w.Header().Set("Content-Type","application/json")

		// 发送数据
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	}

	// 调用服务
	//GIClient := GI.NewGetImageCdService("go.micro.service.GetImageCd", client.DefaultClient)
	//	rsp, err := GIClient.GetImageCode(context.TODO(), &GI.Request{
	//		Uuid:uudi,
	//	})
	GIClient := GI.NewGetImageCdService("go.micro.service.GetImageCd", client.DefaultClient)
	rsp, err := GIClient.GetSmsCode(context.TODO(), &GI.SmsRequest{
		Mobile:mobile,
		Imagestr:test,
		Uuid:id,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 创建返回数据的map
	response := map[string]interface{}{

		"errno": rsp.Error,
		"errmsg": rsp.Errmsg,

	}



	//设置返回数据的格式
	w.Header().Set("Content-Type","application/json")

	// 发送数据
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func rspToImage(rsp *GI.Response) captcha.Image{
	/*
	Error  string `protobuf:"bytes,1,opt,name=Error,proto3" json:"Error,omitempty"`
		Errmsg string `protobuf:"bytes,2,opt,name=Errmsg,proto3" json:"Errmsg,omitempty"`
		//图片结构元素
		Pix []byte `protobuf:"bytes,3,opt,name=Pix,proto3" json:"Pix,omitempty"`
		//图片的跨度
		Stride               int64           `protobuf:"varint,4,opt,name=Stride,proto3" json:"Stride,omitempty"`
		Min                  *Response_Point `protobuf:"bytes,5,opt,name=Min,proto3" json:"Min,omitempty"`
		Max                  *Response_Point `protobuf:"byte
	*/
	var Image captcha.Image
	if rsp == nil{
		return Image
	}


	var img image.RGBA
	img.Stride = int(rsp.Stride)
	img.Pix = rsp.Pix
	img.Rect.Max.X = int(rsp.Max.X)
	img.Rect.Min.X = int(rsp.Min.X)
	img.Rect.Max.Y = int(rsp.Max.Y)
	img.Rect.Min.Y = int(rsp.Min.Y)


	Image.RGBA = &img
	return Image
}

func verifyMobileNum(mobile string) bool{
	mobile_reg:=regexp.MustCompile(`0?(13|14|15|17|18|19)[0-9]{9}`)
	//通过条件判断字符串是否匹配规则 返回正确或失败
	bl :=mobile_reg.MatchString(mobile)
	//如果手机号不匹配那就直接返回错误不调用服务
	return  bl
}

//登陆
func PostRet(w http.ResponseWriter, r *http.Request,_ httprouter.Params) {
	beego.Info("PostRet  注册 /api/v1.0/users")
	//接收post发送过来的数据
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//验证数据
	if request["mobile"].(string) ==""||request["password"].(string)==""||request["sms_code"].(string)==""{
		//准备回传数据
		response := map[string]interface{}{
			"errno": utils.RECODE_DATAERR,
			"errmsg": utils.RecodeText(utils.RECODE_DATAERR),
		}
		//设置返回数据的格式
		w.Header().Set("Content-Type","application/json")
		//发送给前端
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		return

	}



	GIClient := GI.NewGetImageCdService("go.micro.service.GetImageCd", client.DefaultClient)
	rsp, err := GIClient.PostRet(context.TODO(), &GI.PostRetRequest{
		Mobile:request["mobile"].(string),
		Password:request["password"].(string),
		SmsCode:request["sms_code"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//读取cookie   统一cookie   userlogin，返回cookie
	//func (r *Request) Cookie(name string) (*Cookie, error)

	cookie,err :=r.Cookie("userlogin")
	if err!=nil || ""==cookie.Value{
		//创建1个cookie对象
		cookie:= http.Cookie{Name:"userlogin",Value:rsp.SessionId,Path:"/",MaxAge:3600}
		//对浏览器的cookie进行设置
		http.SetCookie(w,&cookie)
	}

	// 创建返回数据的map
	response := map[string]interface{}{

		"errno": rsp.Error,
		"errmsg": rsp.Errmsg,

	}
	//设置返回数据的格式
	w.Header().Set("Content-Type","application/json")

	// 发送数据
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
