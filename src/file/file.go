package file

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"sparta/src/file/encrypt"
	"strconv"
	"time"
)

// Data has the xml data for the initial data tag and then incorporates the Exercise struct.
type Data struct {
	XMLName     xml.Name   `xml:"data"`
	LastUpdated time.Time  `xml:"updated"`
	Exercise    []Exercise `xml:"exercise"`
}

// Exercise keeps track of the data for each exercise that the user has done.
type Exercise struct {
	Date     string  `xml:"date"`
	Clock    string  `xml:"clock"`
	Activity string  `xml:"activity"`
	Distance float64 `xml:"distance"`
	Time     float64 `xml:"time"`
	Reps     int     `xml:"reps"`
	Sets     int     `xml:"sets"`
	Comment  string  `xml:"comment"`
}

var fileStatusEmpty bool

// Empty returns if we have a config file or not.
func Empty() bool {
	return fileStatusEmpty
}

// Config returns the config directory and handles the error accordingly.
func config() string {
	var dir string

	// Workaround golang 1.12 in the cross compiling tool.
	switch runtime.GOOS {
	case "windows":
		dir = os.Getenv("AppData")

	case "darwin":
		dir = os.Getenv("HOME")
		dir += "/Library/Preferences"

	default: // Unix
		dir = os.Getenv("XDG_CONFIG_HOME")
		if dir == "" {
			dir = os.Getenv("HOME")
			dir += "/.config"
		}
	}

	return dir

	/*
		directory, err := os.UserConfigDir()
		if err != nil {
			fmt.Print(err)
		}


		return directory
	*/
}

// DataFile specifies the loacation of our data file.
var DataFile string = filepath.Join(config(), "sparta", "exercises.xml")

// Check does relevant checks around our data file.
func Check(key *[32]byte) (exercises Data) {

	// Check if the user has a data file directory.
	if _, err := os.Stat(DataFile); err == nil { // The folder does exist.
		exercises = readData(key)
	} else if os.IsNotExist(err) { // The file doesn't exist, we should create it.

		// Check if the directory exists. If not, we create it.
		if _, err := os.Stat(DataFile); os.IsNotExist(err) {
			os.Mkdir(filepath.Join(config(), "sparta"), os.ModePerm)
		}

		// We then create the file.
		_, err := os.Create(DataFile)
		if err != nil {
			fmt.Print("Could not create the file.", err)
		}

		// Specify that the file is empty if not proven otherwise.
		fileStatusEmpty = true
	}

	return exercises
}

// ReadData reads data from an xml file, couldn't be simpler. Unexported.
func readData(key *[32]byte) (XMLData Data) {

	// Open up the xml file that already exists.
	file, err := os.Open(DataFile)
	if err != nil {
		fmt.Print(err)
	}

	// Read the data from the opened file and then check if it is empty.
	encrypted, err := ioutil.ReadAll(file)
	if string(encrypted) == "" {
		fileStatusEmpty = true
		return XMLData
	} else if err != nil {
		fmt.Print(err)
	}

	// Close the loading of the file now that it has served is perpous.
	go file.Close()

	// Unencrypt the data to the content variable.
	content := encrypt.Decrypt(key, encrypted)

	// Unmarshal the xml data in to our Data struct.
	xml.Unmarshal(content, &XMLData)

	fileStatusEmpty = false
	return XMLData
}

// Write writes new exercieses to the data file.
func (d *Data) Write(key *[32]byte) {
	// Update the section containing the time that our file was last updated.
	d.LastUpdated = time.Now()

	//Marchal the xml content in to a file variable.
	file, err := xml.Marshal(d)
	if err != nil {
		fmt.Print("Could not marchal the data.", err)
	}

	// Write to the file.
	ioutil.WriteFile(DataFile, encrypt.Encrypt(key, file), 0644)
}

// ParseFloat is a wrapper around strconv.ParseFloat that handles the error to make the function usable inline.
func ParseFloat(input string) float64 {
	if input == "" {
		return 0
	}

	output, err := strconv.ParseFloat(input, 32)
	if err != nil {
		fmt.Print(err)
	}

	return output
}

// ParseInt is just a wrapper around strconv.Atoi().
func ParseInt(input string) int {
	if input == "" {
		return 0
	}

	output, err := strconv.Atoi(input)
	if err != nil {
		fmt.Print(err)
	}

	return output
}

// Format formats the latest updated data in the Data struct to display information.
func (d *Data) Format(i int) (output string) {
	if d.Exercise[i].Reps == 0 && d.Exercise[i].Sets == 0 {
		output = fmt.Sprintf("\nAt %s on %s, you trained %s. The distance was %v kilometers and the exercise lasted for %v minutes.\nThis resulted in an average speed of %.3f km/min.\n",
			d.Exercise[i].Clock, d.Exercise[i].Date, d.Exercise[i].Activity, d.Exercise[i].Distance, d.Exercise[i].Time, d.Exercise[i].Distance/d.Exercise[i].Time)
	} else if d.Exercise[i].Distance == 0 {
		output = fmt.Sprintf("\nAt %s on %s, you trained %s for %v minutes. You did %v sets with %v reps each.\n",
			d.Exercise[i].Clock, d.Exercise[i].Date, d.Exercise[i].Activity, d.Exercise[i].Time, d.Exercise[i].Sets, d.Exercise[i].Reps)
	} else {
		output = fmt.Sprintf("\nAt %s on %s, you trained %s for %v minutes. The distance was %v kilometers and you did %v sets with %v reps each.\n",
			d.Exercise[i].Clock, d.Exercise[i].Date, d.Exercise[i].Activity, d.Exercise[i].Time, d.Exercise[i].Distance, d.Exercise[i].Sets, d.Exercise[i].Reps)
	}

	if d.Exercise[i].Comment != "" {
		output += fmt.Sprintf("Comment: %s\n", d.Exercise[i].Comment)
	}

	return output
}
