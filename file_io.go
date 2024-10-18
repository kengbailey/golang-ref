package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := "./rubber_duck.jpg"

	fileData, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v bytes in %s\n", len(fileData), fileName)

	newFileName := "new_rubber_duck.jpg"
	err = os.WriteFile(newFileName, fileData, 0777)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s written to disk\n", newFileName)
	}
}
