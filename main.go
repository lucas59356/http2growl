package main

import (
	"github.com/lucas59356/http2growl/growl"
	"github.com/lucas59356/http2growl/http"
)

const bind = ":6969"
const growlServer = "localhost:23053"

func main() {
	println("Iniciando...")
	gntp := growlhandler.NewGNTP(growlServer)
	h := httphandler.NewHTTP(gntp.Notify)
	h.ListenAddr = bind
	err := h.Listen()
	panic(err)
}
