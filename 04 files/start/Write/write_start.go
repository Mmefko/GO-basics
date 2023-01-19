package main

import (
	"fmt"
	"log"
	"os"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func checkFileExists(filepath string) bool {
	if _, err := os.Stat(filepath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func writeFileData() {
	//Create a new file
	f, err := os.Create("test.txt")
	handleErr(err)

	//Defer the close after you open the file
	defer f.Close()

	//Get the Name of the file
	fmt.Println("The file name is:", f.Name()) //The file name is: test.txt

	//Write some string data to the file
	f.WriteString("Some text again\n")

	//Write some individual bytes to the file
	data2 := []byte{'a', 'b', 'c', 'd', '\n'}
	f.Write(data2)

	//the Sync function forces the data to be written out
	f.Sync()
}

func appendFileData(fname string, data string) {
	f, err := os.OpenFile(fname, os.O_APPEND|os.O_WRONLY, 0644)
	handleErr(err)
	defer f.Close()

	_, err = f.WriteString(data)
}

func main() {
	//Dump some data to file
	//data1 := []byte("Just a simple string")
	//ioutil.WriteFile("datafile.txt", data1, 0644)

	//Granular writing of data
	if !checkFileExists("test.txt") {
		writeFileData()
	} else {
		os.Truncate("test.txt", 10) //Remove all chars instead of first 10
		fmt.Println("Trimmed the file data")
	}

	//Append data to a file
	appendFileData("datafile.txt", "\nAppended data")
}
