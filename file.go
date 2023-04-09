package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func createFile(path string) *os.File {
	var file *os.File
	// check if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if isError(err) {
			log.Panic(err)
		}
		defer file.Close()
	}

	return file
}

func writeFile(path string) *os.File {
	// Open file using READ & WRITE permission.
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		log.Panic(err)
	}
	defer file.Close()

	// Write some text line-by-line to file.
	_, err = file.WriteString("Hello \n")
	if isError(err) {
		log.Panic(err)
	}
	_, err = file.WriteString("World \n")
	if isError(err) {
		log.Panic(err)
	}

	// Save file changes.
	err = file.Sync()
	if isError(err) {
		log.Panic(err)
	}

	return file
}

func readFile(path string) string {
	// Open file for reading.
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		log.Panic(err)
	}
	defer file.Close()

	// Read file, line by line
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)

		// Break if finally arrived at end of file
		if err == io.EOF {
			break
		}

		// Break if error occured
		if err != nil && err != io.EOF {
			isError(err)
			break
		}
	}

	return string(text)
}

func deleteFile(path string) {
	// delete file
	var err = os.Remove(path)
	if isError(err) {
		return
	}
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
