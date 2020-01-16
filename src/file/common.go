package file

import (
	"fmt"
	"io/ioutil"
	"os"
)

// OpenFile is a common code for opening and reading the data from files (used by settings and file packages).
func OpenFile(filename string) []byte {
	// Open up the file and it's content.
	file, err := os.Open(filename)
	if err != nil {
		fmt.Print(err)
	}

	// Read the data to extract the content in a readable fashion.
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Print(err)
	}

	// Close the file opening.
	go file.Close()

	return content
}
