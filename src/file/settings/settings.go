package settings

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sparta/src/file"
)

// TODO: Clean up duplicate code between file and settings packages. Remove duplicate directory checks.

// SettingsFile is a global variable for storing the path of the settings file.
var SettingsFile = filepath.Join(file.DataPath, "settings.xml")

// Settings holds all settings for the user.
type Settings struct {
	Theme string `xml:"theme"`
}

// NewSettings initializes a new settings file with default values.
func NewSettings() Settings {
	return Settings{Theme: "Dark"}
}

func (s Settings) Write() {
	//Marchal the xml content in to a file variable.
	file, err := xml.Marshal(s)
	if err != nil {
		fmt.Print("Could not marchal the data.", err)
	}

	// Write to the file.
	ioutil.WriteFile(SettingsFile, file, 0644)
}

func readData() (settings Settings) {
	// Open up the xml file that already exists.
	file, err := os.Open(SettingsFile)
	if err != nil {
		fmt.Print(err)
	}

	// Read the content of the file.
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Print(err)
	}

	// Unmarshal the xml data in to our Settings struct.
	xml.Unmarshal(content, &settings)

	return settings
}

// Check makes relevant checks around finding the stetings file.
func Check() (settings Settings) {
	// Check if the user has a data file directory.
	if _, err := os.Stat(SettingsFile); err == nil { // The file does exist.
		settings = readData()
	} else if os.IsNotExist(err) { // The file doesn't exist, we should create it.

		// Check if the directory exists. If not, we create it.
		if _, err := os.Stat(SettingsFile); os.IsNotExist(err) {
			os.Mkdir(filepath.Join(file.Config(), "sparta"), os.ModePerm)
		}

		// We then create the file.
		_, err := os.Create(SettingsFile)
		if err != nil {
			fmt.Print("Could not create the file.", err)
		}

		// Create new data for the settings.
		settings = NewSettings()

		// Write the changes and do so concurrently.
		go settings.Write()
	}

	return settings
}
