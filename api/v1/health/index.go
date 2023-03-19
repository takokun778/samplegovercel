package user

import (
	"net/http"

	"github.com/takokun778/samplegovercel/internal/handler"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	handler.Health(w, r)
}
