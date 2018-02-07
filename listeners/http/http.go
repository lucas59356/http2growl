package httphandler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucas59356/http2growl/utils"
)

const bind = ":6969"

// HTTP Handler
type HTTP struct {
	Router               *mux.Router
	ListenAddr           string
	NotificationCallback func(*utils.Message) error
}

// NewHTTP Generate new http object
func NewHTTP(cb func(*utils.Message) error) *HTTP {
	return &HTTP{
		Router:               mux.NewRouter(),
		ListenAddr:           bind,
		NotificationCallback: cb,
	}
}

// Listen http.ListenAndServe
func (h *HTTP) Listen() error {
	h.Router.HandleFunc("/notify", h.Handle)
	println("Escutando em " + h.ListenAddr)
	return http.ListenAndServe(h.ListenAddr, h.Router)
}

// Handle http handler
func (h *HTTP) Handle(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	LogRequests(w, r)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	m := utils.Message{
		Title:   r.Form.Get("title"),
		Message: r.Form.Get("text"),
		Icon:    r.Form.Get("icon"),
		Sticky:  false,
		Event:   r.Form.Get("from"),
		From:    r.Form.Get("from"),
	}
	err = h.NotificationCallback(&m)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, err.Error())
		println(err.Error())
	}
}
