/*
   Ken Bailey
   3/7/18

   Useful net-based Go code snippets and concepts.

*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const url = "https://old.reddit.com"

func main() {
	// Making a network request.
	// Two pieces in play are the http Client object and a http request.
	// Http.Client can be used across many requests.
	httpClient := http.Client{Timeout: time.Second * 2}

	// Rqeust object
	// Takes a request type method, url, and optional io.Reader body
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Make Request
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// read respnse into byte slice
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// print page body
	fmt.Println(string(body))
}
