/*
   Ken Bailey
   12/16/18

   Useful net-based Go code snippets and concepts.

*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

const url = "https://old.reddit.com"

func ipToHost(ip string) (hostList []string, err error) {
	hostList, err = net.LookupHost(ip)
	if err != nil {
		return nil, err
	}

	return
}

func hostToIP(host string) (ipList []string, err error) {
	res, err := net.LookupIP(host)
	if err != nil {
		return nil, err
	}

	for _, ip := range res {
		ipList = append(ipList, ip.String())
	}

	return
}

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

	// set request header
	req.Header.Set("password", "passw0rd")
	req.Header.Set("username", "kenkenken")

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

	////////////////////////////////////////////////////////////////

	// Convert Hostname to IP Address
	// Used to find the IP Address of a particular URL
	ips, err := hostToIP("https://www.reddit.com")
	if err != nil {
		log.Fatal(err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
