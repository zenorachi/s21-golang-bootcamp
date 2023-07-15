package handler

import (
	"log"
	"net/http"
)

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("METHOD: [%v] URI: [%v]", r.Method, r.RequestURI)
	handlePost(w, r)
}
