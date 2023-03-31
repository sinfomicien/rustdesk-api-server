package controllers

import (
	"encoding/json"
	"rustdesk-api-server/app/dto"
	"rustdesk-api-server/app/services"
	"rustdesk-api-server/utils/beegoHelper"
	"strings"
	"time"
)

var Address = new(AddressBookController)

type AddressBookController struct {
	BaseController
}

// View Address Spectrum List
func (ctl *AddressBookController) List() {
	ack := dto.AbGetAck{}
	ack.Tags = []string{}
	// 查询 tags
	tags := services.Tags.FindTags(ctl.loginUserInfo.Id)
	for _, item := range tags {
		ack.Tags = append(ack.Tags, item.Tag)
	}

	// 查询 peers
	ack.Peers = []dto.AbGetPeer{}
	peerDbs := services.Peers.FindPeers(ctl.loginUserInfo.Id)
	for _, item := range peerDbs {
		ack.Peers = append(ack.Peers, dto.AbGetPeer{
			Id:       item.ClientId,
			Username: item.Username,
			Hostname: item.Hostname,
			Alias:    item.Alias,
			Platform: item.Platform,
			Tags:     strings.Split(item.Tags, ","),
		})
	}

	// Query the list of all logged-in accounts
	tokens := services.Token.FindTokens(ctl.loginUserInfo.Id)
	for _, item := range *tokens {
		ist := false
		for _, bookItem := range ack.Peers {
			if bookItem.Id == item.ClientId {
				ist = true
				break
			}
		}
		if !ist {
			ack.Peers = append(ack.Peers, dto.AbGetPeer{
				Id:       item.ClientId,
				Username: item.Username,
				Hostname: item.ClientId,
				Alias:    "Owner of:" + item.ClientId,
				Platform: "none",
				Tags:     strings.Split("", ","),
			})
		}
	}

	jdata, _ := json.Marshal(ack)

	ctl.JSON(beegoHelper.H{
//		"error":     false,
		"data":      string(jdata),
		"update_at": time.Now().Format("2006-01-02 15:04:05"),
	})
}

// 更新地址谱
func (ctl *AddressBookController) Update() {
	req := dto.AbUpdateReq{}

	if err := ctl.BindJSON(&req); err != nil {
		ctl.JSON(beegoHelper.H{
			"error": "Request parameter exception",
		})
		return
	}

	// 解析数据
	reqSub := &dto.AbUpdateSub{}
	err := json.Unmarshal([]byte(req.Data), reqSub)
	if err != nil {
		ctl.JSON(beegoHelper.H{
			"error": "Request data exception",
		})
	}

	// 批量删除tags
	services.Tags.DeleteAll(ctl.loginUserInfo.Id)
	// 批量删除Peers
	services.Peers.DeleteAll(ctl.loginUserInfo.Id)

	// 开始批量插入tags
	if !services.Tags.BatchAdd(ctl.loginUserInfo.Id, reqSub.Tags) {
		ctl.JSON(beegoHelper.H{
			"error": "Failed to import tags",
		})
	}
	// 开始批量插入peers
	if !services.Peers.BatchAdd(ctl.loginUserInfo.Id, reqSub.Peers) {
		ctl.JSON(beegoHelper.H{
			"error": "Import address book failed",
		})
	}

	ctl.JSON(beegoHelper.H{
		"data": "success",
	})

}
