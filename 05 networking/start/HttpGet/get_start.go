package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func getRequest() {
	const httpbin = "https://httpbin.org/get"

	//Perform a GET operation
	resp, err := http.Get(httpbin)
	if err != nil {
		return
	}

	//The caller is responsible for closing the response
	defer resp.Body.Close()

	// Can access parts of the response to get information:
	fmt.Printf("Status %v\n", resp.Status)                //Status 200 OK
	fmt.Printf("Status Code %v\n", resp.StatusCode)       //Status Code 200
	fmt.Printf("Protocol %v\n", resp.Proto)               //Protocol HTTP/2.0
	fmt.Printf("Content Length %v\n", resp.ContentLength) //Content Length 273

	//Use a String Builder to build the content from bytes
	var sb strings.Builder

	//Read the content and write it to the builder
	content, _ := ioutil.ReadAll(resp.Body)
	bytecount, _ := sb.Write(content)

	//Format the output
	fmt.Println(bytecount, sb.String()) //273 ... and a JSON
}

func main() {
	//Execute a GET request
	getRequest()
}
