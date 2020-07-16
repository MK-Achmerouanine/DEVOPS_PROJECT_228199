package handler

import (
	"api_app/platform/newsfeed"
	"encoding/json"
	"net/http"
	"log"
	"strconv"
)

func NewsfeedGet(feed newsfeed.Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items := feed.GetAll()
		json.NewEncoder(w).Encode(items)
		
		log.Println("Get newsfeeds => Total:"+ strconv.Itoa(len(items)))
	}
}
