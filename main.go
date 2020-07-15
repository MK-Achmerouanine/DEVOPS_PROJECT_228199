package main

import (
	"api_app/httpd/handler"
	"api_app/platform/newsfeed"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	port := ":5000"
	feed := newsfeed.New()

	r := chi.NewRouter()

	r.Get("/newsfeed", handler.NewsfeedGet(feed))
	r.Post("/newsfeed", handler.NewsfeedPost(feed))

	fmt.Println("Serving on " + port)
	http.ListenAndServe(port, r)
}
