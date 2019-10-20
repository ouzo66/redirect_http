package main

import (
	"log"
	"net/http"

	"google.golang.org/appengine"
)

func redirect(w http.ResponseWriter, req *http.Request) {
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	log.Printf("redirect to: %s", target)
	http.Redirect(w, req, target,
		http.StatusTemporaryRedirect)
}

func main() {
	http.HandleFunc("/", redirect)
	appengine.Main()
}
