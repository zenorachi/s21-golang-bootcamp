package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct{}

type kek struct {
	Money int64  `json:"money"`
	Type  string `json:"candyType"`
	Count int64  `json:"candyCount"`
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Hello\n"))
	case http.MethodPost:
		fmt.Println(r.RequestURI)
		//var data []byte
		var k kek
		json.NewDecoder(r.Body).Decode(&k)
		d, _ := json.Marshal(k)
		w.Write(d)
	}
}
