package handler

import (
	"day04/candy-server/config"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type PostJSON struct {
	Money int64  `json:"money"`
	Type  string `json:"candyType"`
	Count int64  `json:"candyCount"`
}

type CorrectRequest struct {
	Change int64  `json:"change"`
	Thanks string `json:"thanks"`
}

type IncorrectRequest struct {
	Error string `json:"error"`
}

type RequestNotEnoughMoney struct {
	Error string `json:"error"`
}

var candies map[string]int64

func init() {
	var cfg config.Candies
	e := cfg.FillModels()
	if e != nil {
		log.Fatalln(e)
	}
	candies = cfg.GetMapModels()
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if r.RequestURI == "/buy_candy" {
			var post PostJSON
			json.NewDecoder(r.Body).Decode(&post)
			if _, ok := candies[post.Type]; !ok || post.Count <= 0 || post.Money <= 0 {
				badResp, _ := json.Marshal(IncorrectRequest{
					Error: "error: invalid input data",
				})
				w.WriteHeader(http.StatusBadRequest)
				w.Write(badResp)
			} else if post.Money < candies[post.Type]*post.Count {
				badRespMoney, _ := json.Marshal(IncorrectRequest{
					Error: fmt.Sprintf("You need %d more money!", candies[post.Type]*post.Count-post.Money),
				})
				w.WriteHeader(http.StatusPaymentRequired)
				w.Write(badRespMoney)
			} else {
				okResp, _ := json.Marshal(CorrectRequest{
					Change: post.Money - candies[post.Type]*post.Count,
					Thanks: "Thank you!",
				})
				w.WriteHeader(http.StatusCreated)
				w.Write(okResp)
			}
			w.Write([]byte("\n"))
		} else {
			http.Error(w, "error: path not found", http.StatusNotFound)
		}
	default:
		http.Error(w, "error: http method not allowed", http.StatusMethodNotAllowed)
	}
}
