package users

import (
	"log"
	"net/http"

	"github.com/takokun778/samplegovercel/internal/handler"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL)

	switch r.Method {
	case http.MethodGet:
		handler.GetUser(w, r)
	case http.MethodPost:
		handler.PostUser(w, r)
	case http.MethodPut:
		handler.PutUser(w, r)
	case http.MethodDelete:
		handler.DeleteUser(w, r)
	case http.MethodOptions:
		w.Write([]byte(""))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
