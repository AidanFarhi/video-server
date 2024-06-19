package main

import (
	"fmt"
	"net/http"
)

func init() {
	fmt.Println("init")
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "views/index.html")
}

func StaticHandler() http.Handler {
	return http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", IndexHandler)

	mux.Handle("/static/", StaticHandler())

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
