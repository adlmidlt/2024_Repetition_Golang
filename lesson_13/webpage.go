package lesson_13

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Page struct {
	URL  string
	Size int
}

func responseSize(url string, channel chan Page) {
	//fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	channel <- Page{URL: url, Size: len(body)}
}

func HttpNet() {
	pages := make(chan Page)
	urls := []string{"https://example.com", "https://golang.org", "https://golang.org/doc"}
	for _, url := range urls {
		go responseSize(url, pages)
	}

	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Printf("%s: %d\n", page.URL, page.Size)
	}

	//time.Sleep(5 * time.Second)
	//myChan := make(chan string)
	//go greeting(myChan)
	//str := <-myChan
	//fmt.Println(str)

	//channel1 := make(chan string)
	//channel2 := make(chan string)
	//
	//go abc(channel1)
	//go def(channel2)
	//
	//fmt.Print(<-channel1)
	//fmt.Print(<-channel2)
	//
	//fmt.Print(<-channel1)
	//fmt.Print(<-channel2)
	//
	//fmt.Print(<-channel1)
	//fmt.Print(<-channel2)
	//
	//fmt.Println()

	//myChan := make(chan string)
	//go send(myChan)
	//reportNap("receiving goroutine", 5)
	//fmt.Println(<-myChan)
	//fmt.Println(<-myChan)
}

func reportNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "sleeping")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, "wake up!")
}

func send(myChannel chan string) {
	reportNap("sending goroutine", 2)
	fmt.Println("***sending value***")
	myChannel <- "a"
	fmt.Println("***sending value***")
	myChannel <- "b"
}

func myChannel() {
	myChan := make(chan float64)
	myChan <- 3.14
}

func greeting(myChan chan string) {
	myChan <- "hi"
}

func abc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"

}

func def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}
