package controllers

import (
	"rustdesk-api-server/app/dto"
	"rustdesk-api-server/app/services"
	"rustdesk-api-server/utils/beegoHelper"
	"rustdesk-api-server/utils/common"
	"strings"
)

var Login = new(LoginController)

type LoginController struct {
	BaseController
}

// Login
func (ctl *LoginController) Login() {
	if ctl.Ctx.Input.IsPost() {
		// Get request parameters
		var req dto.LoginReq
		if err := ctl.BindJSON(&req); err != nil {
			ctl.JSON(common.JsonResult{
				Error: err.Error(),
			})
		}
		req.Username = strings.TrimSpace(req.Username)
		if len(req.Username) == 0 {
			ctl.JSON(common.JsonResult{
				Code:  -1,
				Error: "Username can not be empty",
			})
		}
		req.Password = strings.TrimSpace(req.Password)
		if len(req.Password) == 0 {
			ctl.JSON(common.JsonResult{
				Code:  -1,
				Error: "password can not be blank",
			})
		}
		req.ClientId = strings.TrimSpace(req.ClientId)
		if len(req.ClientId) == 0 {
			ctl.JSON(common.JsonResult{
				Code:  -1,
				Error: "Client ID cannot be empty",
			})
		}

		// Check if password correspond to user
		token, err := services.Login.UserLogin(req.Username, req.Password, req.ClientId, req.Uuid, ctl.Ctx)
		if err != nil {
			ctl.JSON(common.JsonResult{
				Code:  -1,
				Error: err.Error(),
			})
		}

		ctl.JSON(beegoHelper.H{
			"type": "access_token",
			"access_token": token,
			"user": beegoHelper.H{
				"name": req.Username,
				"grp": "default",
				"is_admin": true,
			},
		})

	}

}
