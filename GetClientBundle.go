// This script prints the download file. Which is not that optimal since it needs to be unziped.

package main

import (
	"archive/zip"
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	// It will be used later...
	zipFileName := "bundle.zip"
	downloadedFileName := "clientbundle"
	//

	url := "https://xxxxxxxxxxx/api/clientbundle"
	method := "GET"

	// Make the Go client to ignore the TLS verification
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: transCfg}

	// Build the request (req) with the previous components
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Adding Authorization Header
	req.Header.Add("Authorization", "Bearer ed840c02-9ece-4188-b910-17fc77bf14fb")
	req.Header.Add("Accept-Encoding", "gzip") // I think this can be removed.

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(string(body))

	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)
	fh := &zip.FileHeader{
		Name:     downloadedFileName,
		Modified: time.Now(),
		Method:   0, // This controls whether the files is extracted or inflated. ??
	}
	f, err := w.CreateHeader(fh)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write(body); err != nil {
		log.Fatal(err)
	}
	err = w.Close()
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create(zipFileName)
	if err != nil {
		log.Fatal(err)
	}
	if _, err = io.Copy(file, buf); err != nil {
		log.Fatal(err)
	}
	file.Close()
	fmt.Println("Done.")
}
