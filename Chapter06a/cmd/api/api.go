package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	mux.HandleFunc("/timeout", timeoutHandler)
	mux.HandleFunc("/error", errorHandler)

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("error starting api: ", err)
		os.Exit(1)
	}
}

func timeoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got /timeout request")
	<-time.After(time.Second * 2)
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("this took a long time"))
	log.Fatal(err)
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got /error request")
	w.WriteHeader(http.StatusInternalServerError)
	_, err := w.Write([]byte("internal service error"))
	log.Fatal(err)

}
