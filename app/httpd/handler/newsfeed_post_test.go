package handler

import (
	"api_app/platform/mock_http"
	"api_app/platform/newsfeed"
	"net/http"
	"testing"
)

func TestNewsfeedPost(t *testing.T) {
	feed := newsfeed.New()

	headers := http.Header{}
	headers.Add("content-type", "application/json")

	w := &mock_http.ResponseWriter{}
	r := &http.Request{
		Header: headers,
	}

	r.Body = mock_http.RequestBody(map[string]string{
		"title": "hello",
		"post":  "world",
	})

	handler := NewsfeedPost(feed)
	handler(w, r)

	result := w.GetBodyString()

	if result != "Good job! you added hello" {
		t.Errorf("Handler did not complete")
	}

	if len(feed.GetAll()) != 1 {
		t.Errorf("Item did not add")
	}

	if feed.GetAll()[0].Title != "hello" {
		t.Errorf("Item bad")
	}
}
