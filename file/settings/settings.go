package settings

import (
	"sparta/file"

	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// SettingsFile is a global variable for storing the path of the settings file.
var SettingsFile = filepath.Join(file.DataPath, "settings.xml")

// Config holds all settings for the user.
type Config struct {
	Theme string `xml:"theme"`
}

// NewSettings initializes a new settings file with default values.
func NewSettings() Config {
	return Config{Theme: "Dark"}
}

func (c Config) Write() {
	//Marchal the xml content in to a file variable.
	file, err := xml.Marshal(c)
	if err != nil {
		fmt.Print("Could not marchal the data.", err)
	}

	// Write to the file.
	err = ioutil.WriteFile(SettingsFile, file, 0644)
	if err != nil {
		fmt.Print(err)
	}
}

func readData() (config Config) {
	// Unmarshal the xml data in to our Settings struct.
	err := xml.Unmarshal(file.OpenFile(SettingsFile), &config)
	if err != nil {
		fmt.Print(err)
	}

	return config
}

// Check makes relevant checks around finding the stetings file.
func Check() (config Config) {
	// Check if the user has a data file directory.
	if _, err := os.Stat(SettingsFile); err == nil { // The file does exist.
		config = readData()
	} else if os.IsNotExist(err) { // The file doesn't exist, we should create it.

		// Check if the directory exists. If not, we create it.
		if _, err := os.Stat(SettingsFile); os.IsNotExist(err) {
			err := os.Mkdir(filepath.Join(file.Config(), "sparta"), os.ModePerm)
			if err != nil {
				fmt.Print(err)
			}
		}

		// We then create the file.
		_, err := os.Create(SettingsFile)
		if err != nil {
			fmt.Print("Could not create the file.", err)
		}

		// Create new data for the settings.
		config = NewSettings()

		// Write the changes and do so concurrently.
		go config.Write()
	}

	return config
}
