package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Concurrency works/dealing a lot of things at once while parallism is about doing a lots of things at once

// this simulates blocking functions
func slowFunc() {
	time.Sleep(time.Second * 2) // pause execution for 2 seconds
	fmt.Println("sleeper() finished")
}

func responseTime(url string) {
	start := time.Now()

	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close() // clean up

	elapsed := time.Since(start).Seconds()

	fmt.Printf("%s took %v seconds \n", url, elapsed)
}

func main() {
	// go slowFunc()
	// fmt.Println("I'm not shown until slowFunc() completes")

	// show concurrent execution
	// time.Sleep(time.Second * 3)

	urls := make([]string, 3)
	urls[0] = "https://www.usa.gov/"
	urls[1] = "https://www.gov.uk/"
	urls[2] = "https://www.gouvernement.fr/"

	for _, u := range urls {
		go responseTime(u)
	}

	time.Sleep(time.Second * 5)
}
