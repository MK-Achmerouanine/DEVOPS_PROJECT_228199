package main

import "testing"

func TestHello(t *testing.T) {
    got := Hello()
    want := "Hello, DevOps"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}