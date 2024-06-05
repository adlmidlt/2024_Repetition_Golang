package myproject

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Guestbook struct {
	SignatureCount int
	Signatures     []string
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	Check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	Check(scanner.Err())

	return lines
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	signatures := getStrings("lesson_16/myproject/signature.txt")
	html, err := template.ParseFiles("lesson_16/myproject/view.html")
	Check(err)

	guestbook := Guestbook{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}
	err = html.Execute(w, guestbook)
	Check(err)
}

func NewHandler(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("lesson_16/myproject/new.html")
	Check(err)
	err = html.Execute(w, nil)
	Check(err)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	signature := r.FormValue("signature")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("lesson_16/myproject/signature.txt", options, os.FileMode(0600))
	Check(err)
	_, err = fmt.Fprintln(file, signature)
	defer file.Close()
	Check(err)

	http.Redirect(w, r, "/guestbook", http.StatusFound)
}
