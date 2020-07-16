package main

import (
	"os"
	"net"
	"api_app/httpd/handler"
	"api_app/platform/newsfeed"

	"net/http"
	"log"

	"github.com/go-chi/chi"
)

func main() {
	

	port := ":5000"
	feed := newsfeed.New()

	r := chi.NewRouter()

	r.Get("/newsfeed", handler.NewsfeedGet(feed))
	r.Post("/newsfeed", handler.NewsfeedPost(feed))
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("Not found"))
		log.Println("404 error: Not found ", r.URL.Path[1:])
	})
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
				log.Println("Serving on " +ipnet.IP.String()+ port)
			}
		}
	}

	log.Println("GET : /newsfeed  | Get newsfeeds")
	log.Println("POST: /newsfeed  | Add newsfeed {title,post}")
	http.ListenAndServe(port, r)
}
