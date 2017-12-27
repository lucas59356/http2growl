package httphandler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func LogRequests(w http.ResponseWriter, r *http.Request) {
	f, err := json.Marshal(r.Form)
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("%s %s FROM %s -> %v\n", r.Method, r.RequestURI, r.RemoteAddr, string(f))
}
