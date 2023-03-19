package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/takokun778/samplegovercel/internal/openapi"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := getID(w, r)

	log.Printf("get user: %s", id)

	if id == "" {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	res := openapi.V1GetUsersResponse{
		User: openapi.User{
			Id:   id,
			Name: "村山美羽",
		},
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	id := uuid.NewString()

	log.Printf("post user: %s", id)

	var req openapi.V1PutUsersRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	name := req.Name
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	res := openapi.V1PostUsersResponse{
		User: openapi.User{
			Id:   id,
			Name: name,
		},
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
}

func PutUser(w http.ResponseWriter, r *http.Request) {
	id := getID(w, r)

	log.Printf("put user: %s", id)

	if id == "" {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	var req openapi.V1PutUsersRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	name := req.Name
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	res := openapi.V1PostUsersResponse{
		User: openapi.User{
			Id:   id,
			Name: name,
		},
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := getID(w, r)

	log.Printf("delete user: %s", id)

	if id == "" {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	w.Write([]byte("OK"))
}

func getID(w http.ResponseWriter, r *http.Request) string {
	return r.URL.Query().Get("id")
}
