package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//Create a directory
	os.Mkdir("newdir", os.ModePerm) //new directory created ./newdir

	//Create a deep directory
	os.MkdirAll("path/to", os.ModePerm) //new directory created ./path/to

	//Remove will remove an item
	defer os.Remove("newdir")

	//RemoveAll will remove an item and all children
	defer os.RemoveAll("path")

	//Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Current dir is %v\n", dir) //Current dir is /Users/....

	//Get the dirrectory of the curent process
	procdir, err := os.Executable()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Current dir for the process is %v\n", procdir) //Current dir for the process is /var/folders/...

	//Read the contents of a directory
	contents, _ := ioutil.ReadDir(".")
	for _, fi := range contents {
		fmt.Printf("Name is %v : is dir = %v\n", fi.Name(), fi.IsDir())
	}
}
