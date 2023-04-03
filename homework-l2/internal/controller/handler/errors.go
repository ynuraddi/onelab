package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	message, _ := json.Marshal(struct {
		Err string `json:"error"`
	}{
		Err: "page not found",
	})

	w.WriteHeader(http.StatusNotFound)
	w.Write(message)
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	message, _ := json.Marshal(struct {
		Err string `json:"error"`
	}{
		Err: "method not allowed",
	})

	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write(message)
}

func InternalServerError(w http.ResponseWriter, err error) {
	message := "Internal server error: " + err.Error()

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(message))

	log.Println(message)
}

func BadRequest(w http.ResponseWriter, err error) {
	message := "Bad request: " + err.Error()

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(message))

	log.Println(message)
}
