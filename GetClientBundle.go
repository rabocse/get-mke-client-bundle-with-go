package main

import (
	"archive/zip"
	"bytes"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// flagsHander parses the flags passed by the user via CLI
func flagsHandler() (c, t string) {

	// Requesting flags to user via CLI.
	// NOTE: flag.String returns a pointer.
	clus := flag.String("clus", " ", "Server cluster Name")
	tokn := flag.String("tokn", " ", "Authentication Token from cluster")

	// Execute the command-line parsing
	flag.Parse()

	// Convert (dereference) the string pointer to get a string
	c = *clus
	token := *tokn
	t = fmt.Sprintf("Bearer %s", token)

	return c, t

}

//  buildURL returns a valid string URL
func buildURL(clusterName string) string {

	// Define the components for the HTTP Request.
	const protocol string = "https://"
	const resource string = "/api/clientbundle"

	// Concatenate to build the URL
	url := fmt.Sprintf("%s%s%s", protocol, clusterName, resource)

	return url
}

// craftRequest prepares a valid HTTP request with a POST method and the specified URL and payload.
func craftRequest(m string, u string, p io.Reader) *http.Request {

	// Build the request (req) with the previous components
	req, err := http.NewRequest(m, u, p)

	if err != nil {
		fmt.Println(err)
	}

	// Header to specify that our request sends plain text format.
	req.Header.Add("Content-Type", "text/plain")

	return req

}

func main() {

	// It will be used later...
	zipFileName := "bundle.zip"
	downloadedFileName := "clientbundle"

	// Values are passed via CLI
	cluster, token := flagsHandler()

	// Define the components for the HTTP Request.
	const method string = "GET"

	// Cluster URL is built.
	url := buildURL(cluster)

	fmt.Println("########### INPUT: Server ##########################")
	fmt.Println("Cluster: ", url)
	fmt.Println(" ")
	fmt.Println("########### INPUT: Token #####################")
	fmt.Println("Authentication Token: ", token) // Just for testing purposes.
	fmt.Println(" ")

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

	// Concatenate "Bearer" string + token varible to then pass in a Header

	authToken := fmt.Sprintf("Bearer %s", token)

	// Adding Authorization Header
	req.Header.Add("Authorization", authToken) //
	req.Header.Add("Accept-Encoding", "gzip")  // I think this can be removed.

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
