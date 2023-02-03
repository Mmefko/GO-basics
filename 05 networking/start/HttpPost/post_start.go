package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func postRequest() {
	const httpbin = "https://httpbin.org/post"

	//POST operation using post
	reqBody := strings.NewReader(`
	{
		"field1" : "This is field 1",
		"field2" : 250
	}
		`)

	resp, err := http.Post(httpbin, "application/json", reqBody)
	if err != nil {
		return
	}
	content, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Printf("%s\n", content)

	//POST operation using PostForm
	data := url.Values{}
	data.Add("field1", "This is field 1")
	data.Add("field2", "350")
	resp, err = http.PostForm(httpbin, data)
	if err != nil {
		return
	}
	content, _ = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Printf("%s\n", content)
}

func main() {
	//Execute a POST
	postRequest()
}
