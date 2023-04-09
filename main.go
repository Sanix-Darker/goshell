package main

import (
	"fmt"
)

func main() {
	fmt.Println(WATERMARK)

	// we set up files
	// (clearning them or create them if they don't exist)
	deleteFile(TEMP_FILE)
	createFile(TEMP_FILE)

	deleteFile(HIST_FILE)
	createFile(HIST_FILE)

	// then we write the code and execute it
	ProcessCode()
}
