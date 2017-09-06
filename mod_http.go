package main

import (
	"./log"
	"net/http"
)

var (
	mux          = http.NewServeMux()
	httpListenOn = cmd.String("h", "0.0.0.0:6969", "Onde bindar o servidor da API")
)

func httpInit() {
	logger := log.NewLogger("HTTP")
	cmd.Parse(cmd.Args())
	logger.Info("[!] Iniciando módulo http")
	mux.HandleFunc("/notify", httpHandleRequest)
	logger.Debug("[HTTP] Escutando em " + *httpListenOn)
	err := http.ListenAndServe(*httpListenOn, mux)
	if err != nil {
		logger.Panic(err)
	}
}

func httpHandleRequest(w http.ResponseWriter, req *http.Request) {
	logger := log.NewLogger("HTTP-HANDLER")
	req.ParseForm()
	var text = req.Form.Get("text")
	var title = req.Form.Get("title")
	var sticky = req.Form.Get("sticky")
	var icon = req.Form.Get("icon")
	var n = gntpNotification{}
	if text != "" { // handle do parametro text
		n.Text = text
	} else {
		w.Write([]byte("500 empty text"))
		logger.Error("Erro na requisição: Text vazio")
	}

	if title != "" { // handle do parametro title
		n.Title = title
	} else {
		w.Write([]byte("500 empty title"))
		logger.Error("Erro na requisição: Title vazio")
	}

	switch sticky { // handle do parametro sticky
	case "True":
		n.Sticky = true
	case "true":
		n.Sticky = true
	case "1":
		n.Sticky = true
	case "False":
		n.Sticky = false
	case "false":
		n.Sticky = false
	case "0":
		n.Sticky = false
	case "":
		break
	default:
		w.Write([]byte("500 sticky not valid"))
		logger.Error("Erro na requisição: Sticky inválido")
	}

	if icon != "" { // handle do parametro sticky
		n.Icon = icon
	}
	n.Notify()
}
