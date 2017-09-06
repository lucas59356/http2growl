package main

import (
	"./log"
	"errors"
	"github.com/kariudo/gopushbullet"
	"time"
)

var (
	pbKey        = cmd.String("pbk", "", "Chave de cliente do pushbullet")
	pbClient     = pbRegister()
	lastPushTime = float32(0.2)
)

func pbRegister() *pushbullet.Client {
	if *pbKey == "" {
		logger := log.NewLogger("PUSHBULLET-REGISTER")
		cmd.Usage()
		logger.Panic(errors.New("Chave inválida"))
	}
	return pushbullet.ClientWithKey(*pbKey)
}

func pbInit() { // Vai rodano o loop
	logger := log.NewLogger("PUSHBULLET")
	logger.Info("[!] Iniciando módulo pushbullet")
	logger.Info("Estou pronto")
	for {
		time.Sleep(time.Second * 2)
		go pbFetch()
	}
}

func pbFetch() {
	logger := log.NewLogger("PUSHBULLET-FETCH")
	pushes, err := pbClient.GetPushHistory(lastPushTime)
	if err != nil {
		logger.Error(err.Error())
	}
	for _, push := range pushes {
		if lastPushTime < push.Created {
			pbNotify(push)
			lastPushTime = push.Created
		}
	}
}

func pbNotify(m pushbullet.PushMessage) {
	logger := log.NewLogger("PUSHBULLET-NOTIFY")
	logger.Debug("Enviando " + m.Type + " para notificação.")
	n := gntpNotification{}
	if m.Title != "" {
		n.Title = m.Title
	} else {
		n.Title = "Pushbullet"
	}
	n.Text = m.Body
	n.Sticky = false
	n.Icon = "https://pushbullet.com/favicon.ico"
	n.Notify()
}
