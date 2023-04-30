package main

import (
	"fmt"
	"net/http"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	websites := []string{
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websites {
		go getStatus(web)
		waitGroup.Add(1)
	}
	waitGroup.Wait()
}

func getStatus(endpoint string) {
	defer waitGroup.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		panic(err)
	}

	fmt.Println(res.StatusCode)
}
