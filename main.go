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
		d, _ := ioutil.ReadAll(r.Body)
		log.Printf("Data %s\n", d)

		fmt.Fprintf(rw, "Hello World %s\n", d)
	})

	http.HandleFunc("/video", func(http.ResponseWriter, *http.Request) {
		log.Println("The video is supposed to be read and shown here")
	})

	http.ListenAndServe(":9090", nil)
}
