package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//Get the default temporary directory path
	tempPath := os.TempDir()
	fmt.Printf("Default temp path is %v\n", tempPath) //Default temp path is /var/folders/75/7xt06bbs24d2vkmz88vn4xmc0000gp/T/

	//Create a temporary file use TempFile
	tempContent := []byte("Some text")
	tmpFile, err := ioutil.TempFile(tempPath, "tempfile_*.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	if _, err := tmpFile.Write(tempContent); err != nil {
		log.Fatal(err.Error())
	}

	//Read and Print the tempFile content
	data, _ := ioutil.ReadFile(tmpFile.Name())
	fmt.Printf("%s\n", data) //Some text

	//Remove the tempfile when finished
	fmt.Printf("tmpFile name: %v\n", tmpFile.Name()) //tmpFile name: /var/folders/75/7xt06bbs24d2vkmz88vn4xmc0000gp/T/tempfile_2932882046.txt
	defer os.Remove(tmpFile.Name())

	//Creae a temporary directory with ioutil's version of TempDir
	tmpDir, _ := ioutil.TempDir("", "tempdir*")
	fmt.Printf("tmpDir name: %v\n", tmpDir) //tmpDir name: /var/folders/75/7xt06bbs24d2vkmz88vn4xmc0000gp/T/tempdir298616879
	defer os.Remove(tmpDir)
}
