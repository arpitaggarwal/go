package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!") // send data to client side
}

func main() {
	http.HandleFunc("/", helloWorld)         // set router
	err := http.ListenAndServe(":8080", nil) // set listen port
	if err != nil {
		log.Fatal("Error while starting GO http server on port - 8080 : ", err) //log error and exit in case of error at server boot up
	}
}
