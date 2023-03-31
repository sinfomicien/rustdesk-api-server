package controllers

import "rustdesk-api-server/utils/common"

type IndexController struct {
	BaseController
}

func (ctl *IndexController) Index() {
	ctl.JSON(common.JsonResult{
		Code:  1,
		Msg:   "welcome, author github: https://github.com/xiaoyi510",
		Data:  nil,
		Count: 0,
	})
}
