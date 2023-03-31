package controllers

import (
	"github.com/beego/beego/v2/client/orm"
	"log"
	"rustdesk-api-server/utils/beegoHelper"

	//beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/server/web"
	"rustdesk-api-server/app/models"
	"rustdesk-api-server/utils"
	"strings"
)

type BaseController struct {
	web.Controller
	controllerName string
	actionName     string
	loginUserInfo  *models.User
}

func (ctl *BaseController) Prepare() {
    var accessLog string
    accessLog = "IP:" + ctl.Ctx.Input.IP() + " Method:"+ctl.Ctx.Input.Method() + " " + 
               "URL:" + ctl.Ctx.Input.URL() + " " + "UserAgent" + ctl.Ctx.Input.UserAgent()
    log.Println(accessLog)
	controllerName, actionName := ctl.GetControllerAndAction()
	ctl.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	ctl.actionName = strings.ToLower(actionName)
	log.Println("Request", ctl.controllerName, ctl.actionName)
	// 获取token
	token := ctl.Ctx.Input.Header("Authorization")
	if ctl.controllerName != "login" && ctl.controllerName != "index" && !(ctl.controllerName == "user" && (ctl.actionName == "reg" || ctl.actionName == "setpwd")) {
		if token == "" {
			ctl.JSON(beegoHelper.H{
				"error": "User authorization verification failed",
			})
		} else {
			// 校验用户登录
			if !ctl.CheckLogin() {
				ctl.JSON(beegoHelper.H{
					"error": "User authorization information error",
				})
			}
		}

	}
}

type JsonResult struct {
	Code  int         `json:"code"`  // 响应编码：0成功 401请登录 403无权限 500错误
	Msg   string      `json:"msg"`   // 消息提示语
	Data  interface{} `json:"data"`  // 数据对象
	Count int64       `json:"count"` // 记录总数
}

func (this *BaseController) JSON(obj interface{}) {
	this.Data["json"] = obj
	//对json进行序列化输出
	this.ServeJSON()
	this.StopRun()
}

// 校验登录信息
func (ctl *BaseController) CheckLogin() bool {
	token := ctl.Ctx.Input.Header("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")
	// 解密token
	parseToken, err := utils.ParseToken(token)
	if err != nil {
		return false
	}

	// 查找用户登录信息
	var loginTokenInfo models.Token
	err = orm.NewOrm().QueryTable(new(models.Token)).
		Filter("uid", parseToken.UserId).
		Filter("client_id", parseToken.ClientId).
		Filter("access_token", parseToken.AccessToken).
		One(&loginTokenInfo)

	if err != nil {
		return false
	}

	var loginInfo models.User
	err = orm.NewOrm().QueryTable(new(models.User)).
		Filter("id", loginTokenInfo.Uid).
		One(&loginInfo)

	if err != nil {
		return false
	}
	// 判断用户是否被禁用
	ctl.loginUserInfo = &loginInfo
	if ctl.loginUserInfo.Status != 1 {
		ctl.JSON(beegoHelper.H{
			"error": "user is disabled",
		})
		return false
	}

	return true
}
