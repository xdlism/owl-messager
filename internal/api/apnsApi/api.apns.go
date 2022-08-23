package apnsApi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/db/repo"
)

type ApnsInfoResp struct {
	Id         int    `json:"id,omitempty"`
	Mode       int    `json:"mode"`
	MessageId  int    `json:"messageId,omitempty"` // message
	BundleId   string `json:"bundleId,omitempty"`
	Params     string `json:"params,omitempty"`
	Status     int    `json:"status,omitempty"`
	CreateTime string `json:"createTime,omitempty"`
	UpdateTime string `json:"updateTime,omitempty"`
}

type RespWrapper struct {
	app.Response
	ApnsInfoResp
}

func GetByMessage(ctx iris.Context) {
	log.Debug("get apns")
	var resp RespWrapper
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		log.Debug("id must be a int value")
		resp.Response.Code = tool.RespCodeNotFound
		resp.Message = tool.RespMsgIdNum
		tool.ResponseJSON(ctx, resp)
		return
	}
	log.Debug("id:%d", id)
	ms, err := repo.GetApnsByMessageId(id)
	if err != nil {
		log.Debug("get apns failed")
		log.Debug(err)
		resp.Response.Code = tool.RespCodeNotFound
		resp.Message = tool.RespMsgNotFount
		tool.ResponseJSON(ctx, resp)
		return
	}

	var tmpInfo = ApnsInfoResp{
		Id:         ms.Id,
		MessageId:  ms.MessageId,
		BundleId:   ms.BundleId,
		Params:     ms.Params,
		Status:     ms.Status,
		CreateTime: tool.FormatTime(ms.CreateTime),
		UpdateTime: tool.FormatTime(ms.UpdateTime),
	}
	resp.ApnsInfoResp = tmpInfo
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
