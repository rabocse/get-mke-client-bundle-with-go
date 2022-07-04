package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const method string = "POST"

func flagsHandler() (c, u, p string) {

	// Requesting flags to user via CLI.
	// NOTE: flag.String returns a pointer.
	clus := flag.String("clus", " ", "Server cluster Name")
	user := flag.String("user", " ", "Username from cluster")
	pass := flag.String("pass", " ", "Password from cluster")

	// Execute the command-line parsing
	flag.Parse()

	// Convert the string pointer to a string
	c = *clus
	u = *user
	p = *pass

	return c, u, p

}

func buildURL(clusterName string) string {

	// Define the components for the HTTP Request.

	const protocol string = "https://"
	const resource string = "/auth/login"

	// Concatenate to build the URL
	url := fmt.Sprintf("%s%s%s", protocol, clusterName, resource)

	return url
}

func craftPayload(userValue, passwordValue string) io.Reader {

	// Marshall the credentials: From Go struct to JSON.

	type credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	cred := &credentials{Username: userValue, Password: passwordValue}
	credJSON, err := json.Marshal(cred) // credJSON is type []byte
	if err != nil {
		fmt.Println(err)

	}

	p := bytes.NewReader(credJSON) // so credJSON needs to be converted io.Reader to be accepted by next function (http.NewRequest)

	return p
}

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

	cluster, username, password := flagsHandler()

	url := buildURL(cluster)

	payload := craftPayload(username, password)

	req := craftRequest(method, url, payload)

	// Make the Go client to ignore the TLS verification
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: transCfg}

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
	fmt.Println(" ")
	fmt.Println("########### OUTPUT: AUTH TOKEN #####################")
	fmt.Println(string(body))

}
