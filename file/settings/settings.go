package settings

import (
	"sparta/file"

	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Config holds all settings for the user.
type Config struct {
	Theme string `xml:"theme"`
}

func (c Config) Write() {
	//Marchal the xml content in to a file variable.
	data, err := xml.Marshal(c)
	if err != nil {
		fmt.Print("Could not marchal the data.", err)
	}

	// Write to the file.
	err = ioutil.WriteFile(filepath.Join(file.Config(), "sparta", "settings.xml"), data, 0644)
	if err != nil {
		fmt.Print(err)
	}
}

func readData() (config Config) {
	// Open up the file and it's content.
	data, err := os.Open(filepath.Join(file.Config(), "sparta", "settings.xml"))
	if err != nil {
		fmt.Print(err)
	}

	// Read the data to extract the content in a readable fashion.
	content, err := ioutil.ReadAll(data)
	if err != nil {
		fmt.Print(err)
	}

	// Close the file opening.
	go data.Close()

	// Unmarshal the xml data in to our Settings struct.
	err = xml.Unmarshal(content, &config)
	if err != nil {
		fmt.Print(err)
	}

	return config
}

// Check makes relevant checks around finding the stetings file.
func Check() (config Config) {
	// Check if the user has a data file directory.
	if _, err := os.Stat(filepath.Join(file.Config(), "sparta", "settings.xml")); err == nil { // The file does exist.
		config = readData()
	} else if os.IsNotExist(err) { // The file doesn't exist, we should create it.

		// Check if the directory exists. If not, we create it.
		if _, err := os.Stat(filepath.Join(file.Config(), "sparta")); os.IsNotExist(err) {
			err := os.Mkdir(filepath.Join(file.Config(), "sparta"), os.ModePerm)
			if err != nil {
				fmt.Print(err)
			}
		}

		// We then create the file.
		_, err := os.Create(filepath.Join(file.Config(), "sparta", "settings.xml"))
		if err != nil {
			fmt.Print("Could not create the file.", err)
		}

		// Create new data for the settings with sane defaults.
		config = Config{Theme: "Dark"}

		// Write the changes and do so concurrently.
		go config.Write()
	}

	return config
}
