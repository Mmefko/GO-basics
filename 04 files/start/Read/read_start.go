package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func getFileLeght(filepath string) int64 {
	f, err := os.Stat(filepath)
	handleErr(err)
	return f.Size()
}

func main() {
	//Use readFile to read the entire file
	content, err := ioutil.ReadFile("sample.txt")
	handleErr(err)

	//Readfile reads the data as butes, we have to convert to a string
	fmt.Println(string(content))

	//read function can perform a generic read
	const BufSize = 20
	f, _ := os.Open("sample.txt")
	defer f.Close()

	b1 := make([]byte, BufSize)
	for {
		n, err := f.Read(b1)
		if err != nil {
			if err != io.EOF {
				handleErr(err)
			}
			break
		}
		fmt.Println("Bytes read:", n)
		fmt.Println("Content:", string(b1[:n]))
	}

	//Get the length of a file
	l := getFileLeght("sample.txt")
	fmt.Println("File length:", l)

	//Use ReadAt to read from specific index
	b2 := make([]byte, 8)
	_, err = f.ReadAt(b2, l-int64(len(b2)))
	handleErr(err)
	fmt.Println(string(b2))
}
