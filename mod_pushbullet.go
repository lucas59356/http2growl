package main

import (
	"github.com/kariudo/gopushbullet"
	logger "github.com/lucas59356/go-logger"
	"time"
)

var (
	pbKey        = cmd.String("pbk", "", "Chave de cliente do pushbullet")
	pbClient     = pbRegister()
	lastPushTime = float32(0.2)
)

func pbRegister() *pushbullet.Client {
	if *pbKey == "" {
		log := logger.New("PUSHBULLET-REGISTER")
		cmd.Usage()
		log.Panic("Chave inválida")
	}
	return pushbullet.ClientWithKey(*pbKey)
}

func pbInit() { // Vai rodano o loop
	log := logger.New("PUSHBULLET")
	log.Info("[!] Iniciando módulo pushbullet")
	log.Info("Estou pronto")
	for {
		time.Sleep(time.Second * 2)
		go pbFetch()
	}
}

func pbFetch() {
	log := logger.New("PUSHBULLET-FETCH")
	pushes, err := pbClient.GetPushHistory(lastPushTime)
	if err != nil {
		log.Error(err)
	}
	for _, push := range pushes {
		if lastPushTime < push.Created {
			pbNotify(push)
			lastPushTime = push.Created
		}
	}
}

func pbNotify(m pushbullet.PushMessage) {
	log := logger.New("PUSHBULLET-NOTIFY")
	log.Debug("Enviando " + m.Type + " para notificação.")
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
