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

// craftRequest prepares a valid HTTP request with the specified HTTP method, URL and payload("nil" is none)
func craftRequest(m string, u string, h string, p io.Reader) *http.Request {

	// Build the request (req) with the previous components
	req, err := http.NewRequest(m, u, p)

	if err != nil {
		fmt.Println(err)
	}

	// Adding Authorization Header
	req.Header.Add("Authorization", h)        //
	req.Header.Add("Accept-Encoding", "gzip") // I think this can be removed.

	return req

}

// sendRequest executes the so far crafted Request.
func sendRequest(r *http.Request) []byte {

	// Make the Go client to ignore the TLS verification
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: transCfg}

	res, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	return b

}

// saveFile saves the downloaded body to the local filesystem
func saveFile(b []byte) {

	// Declaring constants to use
	const zipFileName string = "bundle.zip"
	const downloadedFileName string = "clientbundle"

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
	if _, err := f.Write(b); err != nil {
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

func main() {

	// Values are passed via CLI
	cluster, token := flagsHandler()

	// Define the components for the HTTP Request.
	const method string = "GET"

	// Cluster URL is built.
	url := buildURL(cluster)

	// Crafting a valid HTTPS request with TLS ignore.
	req := craftRequest(method, url, token, nil)

	// Sending the request and getting the file
	body := sendRequest(req)

	// Writing the file to the local filesystem
	saveFile(body)

}
