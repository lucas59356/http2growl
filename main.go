package main

import (
	"flag"
	"github.com/lucas59356/go-logger"
	"time"
)

var (
	cmd = flag.NewFlagSet("cmd", flag.ExitOnError)
)

func main() {
	// Configurando o logger
	log := logger.New("CORE")
	log.Info("Iniciando...")
	cmd.Parse(flag.Args())
	if cmd.Parsed() {
		loadPlugins() // Iniciar os plugins
	}
	for {
		time.Sleep(10 * time.Second)
	}
}

func loadPlugins() {
	gntpInit() // gntp
	// pbInit()   // pushbullet
	httpInit() // http
}
