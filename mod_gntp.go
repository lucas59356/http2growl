package main

import (
	logger "github.com/lucas59356/go-logger"
	"github.com/mattn/go-gntp"
)

const (
	gntpDisplayName string = "android"
	gntpEvent       string = "android"
)

var (
	gntpClient = gntp.NewClient()
	gntpServer = cmd.String("g", "localhost:23053", "Servidor gntp (growl) onde será enviado as notificações")
)

func gntpInit() {
	cmd.Parse(cmd.Args())
	log := logger.New("GNTP")
	gntpClient.AppName = "Android"
	gntpClient.Server = *gntpServer
	err := gntpClient.Register(
		[]gntp.Notification{{
			DisplayName: "android",
			Enabled:     true,
			Event:       "android",
		}})
	if err != nil {
		log.Panic(err.Error())
	} else {
		log.Info("Aplicativo registrado no growl com sucesso")
	}
}

// Notification Objeto de notificação para uso interno da aplicação
type gntpNotification struct {
	Text   string
	Title  string
	Sticky bool
	Icon   string
}

func (n gntpNotification) Notify() {
	log := logger.New("GNTP")
	toGNTP := gntp.Message{}
	toGNTP.DisplayName = gntpDisplayName
	toGNTP.Event = gntpEvent
	toGNTP.Title = n.Title
	toGNTP.Text = n.Text
	toGNTP.Sticky = n.Sticky
	toGNTP.Icon = n.Icon
	err := gntpClient.Notify(&toGNTP)
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Uma notificação foi enviada. Título: " + n.Title)
	}
}
