package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Oops", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "Hello World %s\n", d)
	})

	http.HandleFunc("/video", func(http.ResponseWriter, *http.Request) {
		log.Println("The video is supposed to be read and shown here")
	})

	http.ListenAndServe(":9090", nil)
}
