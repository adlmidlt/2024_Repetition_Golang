package lesson_15

import (
	"log"
	"net/http"
)

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Begin page!")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func StartWeb() {
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/hello", viewHandlerHello)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

func viewHandlerHello(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello web!")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}
