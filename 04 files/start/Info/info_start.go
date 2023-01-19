package main

import (
	"fmt"
	"log"
	"os"
)

func checkFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func main() {
	//Use the Stat function to get file stats
	stats, err := os.Stat("sample.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	//Check if file exists
	exists := checkFileExists("sample.txt")
	fmt.Printf("File exists: %v\n", exists) //File exists: true

	//Get the file's modification time
	fmt.Println("Modification time:", stats.ModTime()) //2023-01-19 09:28:10.805988783 +0100 CET
	fmt.Println("File Mode:", stats.Mode())            //File Mode: -rw-r--r--
	fmode := stats.Mode()
	if fmode.IsRegular() {
		fmt.Println("This is a regular file") //This is a regular file
	}

	//Get the file size
	fmt.Println("File size:", stats.Size()) //File size: 1007

	//Is this a directory?
	fmt.Println("Is dir?:", stats.IsDir()) //Is dir?: false
}
