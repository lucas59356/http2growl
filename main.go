package main

import (
	"./log"
	"flag"
	"time"
)

var (
	cmd = flag.NewFlagSet("cmd", flag.ExitOnError)
)

func main() {
	// Configurando o logger
	logger := log.NewLogger("CORE")
	logger.Info("[CORE] Iniciando...")
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
