package controllers

import (
	"rustdesk-api-server/app/dto"
	"rustdesk-api-server/app/services"
	"rustdesk-api-server/global"
	"rustdesk-api-server/utils/beegoHelper"
)

type UserController struct {
	BaseController
}

// current user information
func (ctl *UserController) CurrentUser() {
	ctl.JSON(beegoHelper.H{
		"name": ctl.loginUserInfo.Username,
	})
}

// registered user
func (ctl *UserController) Reg() {
	req := dto.UserRegReq{}
	req.Username = ctl.GetString("username")
	req.Password = ctl.GetString("password")
	req.AuthKey = ctl.GetString("auth_key")
	if len(req.Username) < 4 || len(req.Username) > 20 {
		ctl.JSON(beegoHelper.H{
			"error": "Username between 4-20 characters",
		})
	}

	if len(req.AuthKey) == 0 {
		ctl.JSON(beegoHelper.H{
			"error": "Please enter the authorization code",
		})
	}

	// Determine whether the registration key is legal
	if req.AuthKey != global.ConfigVar.App.AuthKey {
		ctl.JSON(beegoHelper.H{
			"error": "authorization code error",
		})
	}

	// to register an account
	if services.User.Reg(req.Username, req.Password) {
		ctl.JSON(beegoHelper.H{
			"msg": "registration success",
		})
	} else {
		ctl.JSON(beegoHelper.H{
			"error": "registration failed",
		})
	}
}

// 修改用户密码
func (ctl *UserController) SetPwd() {
	req := dto.UserSetPwdReq{}
	req.Username = ctl.GetString("username")
	req.Password = ctl.GetString("password")
	req.AuthKey = ctl.GetString("auth_key")
	if len(req.Username) < 4 || len(req.Username) > 20 {
		ctl.JSON(beegoHelper.H{
			"error": "用户名在4-20位之间",
		})
	}

	if len(req.AuthKey) == 0 {
		ctl.JSON(beegoHelper.H{
			"error": "请输入授权码",
		})
	}

	// 判断注册密钥是否合法
	if req.AuthKey != global.ConfigVar.App.AuthKey {
		ctl.JSON(beegoHelper.H{
			"error": "授权码错误",
		})
	}

	// 去注册账号
	if services.User.ResetPassword(req.Username, req.Password) {
		ctl.JSON(beegoHelper.H{
			"msg": "修改成功",
		})
	} else {
		ctl.JSON(beegoHelper.H{
			"error": "修改失败",
		})
	}
}
