package lesson_16

import (
	"headfirstgo/lesson_16/myproject"
	"net/http"
)

func StartHandler() {
	http.HandleFunc("/guestbook", myproject.ViewHandler)
	http.HandleFunc("/guestbook/new", myproject.NewHandler)
	http.HandleFunc("/guestbook/create", myproject.CreateHandler)
	err := http.ListenAndServe(":8080", nil)
	myproject.Check(err)
}
