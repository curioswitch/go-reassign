package a

import (
	"config"
	"io"
	"net/http"
)

var DefaultClient = &http.Client{}

func reassignPattern() {
	io.EOF = nil

	config.DefaultClient = nil
	DefaultClient = nil

	http.DefaultClient = nil    // want "reassigning variable"
	http.DefaultTransport = nil // want "reassigning variable"
}
