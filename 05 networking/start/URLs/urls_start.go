package main

import (
	"fmt"
	"net/url"
)

func main() {
	//Define URL
	s := "https://www.example.com:8080/user?username=John"

	//TODO: Use the Parse function to parse the URL content
	result, _ := url.Parse(s)
	fmt.Println(result.Scheme)   //https
	fmt.Println(result.Host)     //www.example.com:8080
	fmt.Println(result.Path)     // /user
	fmt.Println(result.Port())   //8080
	fmt.Println(result.RawQuery) //username=John

	//Extract the query components into a Values struct
	vals := result.Query()
	fmt.Println(vals["username"]) //[John]

	//Create a URL from components
	newurl := &url.URL{
		Scheme:   "https",
		Host:     "www.example.com:9137",
		Path:     "/args",
		RawQuery: "x=1&y=2",
	}
	s = newurl.String()
	fmt.Println(s) //https://www.example.com:9137/args?x=1&y=2
	newurl.Host = "testtest.com"
	s = newurl.String()
	fmt.Println(s) //https://testtest.com/args?x=1&y=2

	//Create a new Values struct and encode it as a query string
	newvals := url.Values{}
	newvals.Add("x", "100")
	newvals.Add("z", "somestr")
	newurl.RawQuery = newvals.Encode()
	s = newurl.String()
	fmt.Println(s) // https://testtest.com/args?x=100&z=somestr

}
