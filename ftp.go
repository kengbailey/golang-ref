package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/jlaffaye/ftp"
)

// newFTP creates a new ftp connection
func newFTP(username, password, ip string) (conn *ftp.ServerConn, err error) {
	conn, err = ftp.Dial(ip+":21", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		// enable ftp and retry
		err = enableFTP(username, password, ip)
		if err != nil {
			log.Fatal(err)
		}
		conn, err = ftp.Dial(ip+":21", ftp.DialWithTimeout(5*time.Second))
		if err != nil {
			log.Fatal(err)
		}
	}
	err = conn.Login(username, password)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// putFile logs into cam via ftp and puts file @ storeLocation
func putFile(username, password, ip, file, storeLocation string) (err error) {
	// login
	conn, err := newFTP(username, password, ip)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Quit()

	// open local file
	inFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer inFile.Close()

	// store remote file
	err = conn.Stor(storeLocation+file, inFile)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// getFile logs into cam via ftp and retrieves a file to current location
func getFile(username, password, ip, fileLocation string) (err error) {
	// login
	conn, err := newFTP(username, password, ip)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Quit()

	// retrieve file
	resp, err := conn.Retr(fileLocation)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Close()

	// store file
	outFile, err := os.Create(filepath.Base(fileLocation))
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	buf := make([]byte, 1024)
	for {
		n, err := resp.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		if _, err := outFile.Write(buf[:n]); err != nil {
			log.Fatal(err)
		}
	}
	return
}
