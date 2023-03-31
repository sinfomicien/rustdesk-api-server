package controllers

import "rustdesk-api-server/utils/common"

type HBController struct {
        BaseController
}

func (ctl *HBController) Aliveness() {
        ctl.JSON(common.JsonResult{
                Code:  0,
                Msg:   "Server is alive",
                Data:  nil,
                Count: 0,
        })
}