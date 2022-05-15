package sender

import (
	"context"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/db/model"
	"github.com/lishimeng/owl/internal/db/repo"
	"github.com/lishimeng/owl/internal/plugins/loader"
	"github.com/lishimeng/owl/internal/provider/sms"
)

type Sms interface {
	Send(model.SmsMessageInfo) (err error)
}

type smsSender struct {
	ctx       context.Context
	provider  *sms.ProviderManager
	maxWorker int
}

func NewSmsSender(ctx context.Context) (m Sms, err error) {
	m = &smsSender{
		ctx:       ctx,
		provider:  sms.New(),
		maxWorker: 1,
	}
	return
}

func (m *smsSender) Send(p model.SmsMessageInfo) (err error) {
	// sender info
	log.Info("send sms:%d", p.Id)
	//si, err := repo.GetSmsSenderById(p.Sender)
	si, err := repo.GetDefaultSmsSender(0)
	if err != nil {
		log.Info("sms sender not exist")
		log.Info(err)
		return
	}

	tpl, err := repo.GetSmsTemplateById(p.Template)
	if err != nil {
		log.Info("sms template not exist:%d", p.Template)
		return
	}

	provider := loader.Get(si.Vendor)

	var req = sms.Request{
		Template:  tpl.SenderTemplateId,
		Sign:      p.Signature,
		Params:    p.Params,
		Receivers: p.Receivers,
	}
	resp, err := provider.Send(req)

	if err != nil {
		log.Info(err)
		return
	}
	log.Info(resp)
	return
}
