package main

import (
	"./log"
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
	logger := log.NewLogger("GNTP")
	gntpClient.AppName = "Android"
	gntpClient.Server = *gntpServer
	err := gntpClient.Register(
		[]gntp.Notification{{
			DisplayName: "android",
			Enabled:     true,
			Event:       "android",
		}})
	if err != nil {
		logger.Panic(err)
	} else {
		logger.Info("Aplicativo registrado no growl com sucesso")
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
	logger := log.NewLogger("GNTP")
	toGNTP := gntp.Message{}
	toGNTP.DisplayName = gntpDisplayName
	toGNTP.Event = gntpEvent
	toGNTP.Title = n.Title
	toGNTP.Text = n.Text
	toGNTP.Sticky = n.Sticky
	toGNTP.Icon = n.Icon
	err := gntpClient.Notify(&toGNTP)
	if err != nil {
		logger.Error(err.Error())
	} else {
		logger.Info("Uma notificação foi enviada. Título: " + n.Title)
	}
}
