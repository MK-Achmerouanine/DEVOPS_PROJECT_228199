package handler

import (
	"api_app/platform/newsfeed"
	"encoding/json"
	"net/http"
	"log"
)

func NewsfeedPost(feed newsfeed.Adder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		json.NewDecoder(r.Body).Decode(&request)

		feed.Add(newsfeed.Item{
			Title: request["title"],
			Post:  request["post"],
		})

		w.Write([]byte("Good job! you added " + request["title"]))
		log.Println("Newsfeed "+request["title"]+" added successfully")
	}
}
