// https://golangcode.com/download-a-file-from-a-url/

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	url := "https://images-na.ssl-images-amazon.com/images/I/8166xCVDGnL._SL1500_.jpg"
	err := downloader("rubber_duck.jpg", url)
	if err != nil {
		log.Fatal(err)
	}
}

func downloader(destination, url string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	out, err := os.Create(destination)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// We use iocopy so the entire file is not loaded into memory
	// "x" contains the number of bytes written to file
	x, err := io.Copy(out, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v kb file downloaded --> %s \n", float32(x/1024), destination)

	return err
}
