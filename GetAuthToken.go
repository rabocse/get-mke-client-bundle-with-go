package main

import (
	"flag"
	"fmt"
)

func main() {

	// New code from here

	// Requesting flags to user via CLI.
	// NOTE: flag.String returns a pointer.
	clus := flag.String("clus", " ", "Server cluster Name")
	user := flag.String("user", " ", "Username from cluster")
	pass := flag.String("pass", " ", "Password from cluster")

	// Execute the command-line parsing
	flag.Parse()

	// Convert the string pointer to a string
	cluster := *clus
	username := *user
	password := *pass

	// Define the components for the HTTP Request.
	const method string = "POST"
	protocol := "http://"
	resource := "/auth/login"

	// Concatenate to build the URL
	url := fmt.Sprintf("%s%s%s", protocol, cluster, resource)
	fmt.Println(url)

	// Concatenate to build the payload
	fmt.Println(url)
	fmt.Println(username)
	fmt.Println(password)

	// New block of code until here above

	/*  commenting out the rest of the code for now

	// Define the components for the HTTP Request.
	url := "https://xxxxxxxxxxxxxxx/auth/login"
	const method string = "POST"
	// method := "POST"
	payload := strings.NewReader(`{
		"username":"xxxxxxx",
	"password":"xxxxxxxx"
	}
	`)

	// Make the Go client to ignore the TLS verification
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: transCfg}

	// Build the request (req) with the previous components
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Header to specify that our request sends plain text format.
	req.Header.Add("Content-Type", "text/plain")

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

	fmt.Println(string(body))

	*/
}
