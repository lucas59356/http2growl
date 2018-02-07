package main

import (
	"github.com/lucas59356/http2growl/listeners/http"
)

func main() {
	println("Iniciando...")
	handler, err := GetHandler()
	if err != nil {
		panic(err)
	}
	h := httphandler.NewHTTP(handler)
	h.ListenAddr = HTTPBind
	err = h.Listen()
	panic(err)
}
