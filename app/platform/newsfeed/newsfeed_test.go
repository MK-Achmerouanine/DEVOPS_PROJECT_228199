package newsfeed

import (
	"testing"
)

func TestAdd(t *testing.T) {
	feed := New()
	type args struct {
		title string
		post string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Add newsfeed 1",
			args: args{title: "Title 1", post: "Post 1"},
			want: 1,
		},
		{
			name: "Add newsfeed 2",
			args: args{title: "Title 2", post: "Post 2"},
			want: 2,
		},
		{
			name: "Add newsfeed 3",
			args: args{title: "Title 3", post: "Post 3"},
			want: 3,
		},
	}
	for _, tt := range tests {
		item := Item{
			Title: tt.args.title,
			Post:  tt.args.post,
		}
		feed.Add(item)

		t.Run(tt.name, func(t *testing.T) {
			if got := len(feed.Items); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
	// feed.Add(Item{})
	// if len(feed.Items) != 1 {
	// 	t.Errorf("Item was not added")
	// }
}

func TestGetAll(t *testing.T) {
	feed := New()
	feed.Add(Item{})
	results := feed.GetAll()
	if len(results) != 1 {
		t.Errorf("Item was not added")
	}
}
