package file

import (
	"sparta/src/file/encrypt"

	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
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
func Config() string {
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

// DataPath is the path for the sparta config directory.
var DataPath string = filepath.Join(Config(), "sparta")

// DataFile specifies the location of our data file.
var DataFile string = filepath.Join(DataPath, "exercises.xml")

// Check does relevant checks around our data file.
func Check(key *[32]byte) (exercises Data, err error) {

	// Check if the user has a data file directory.
	if _, err := os.Stat(DataFile); err == nil {
		// The file exists and we read the data. Return error if decryption failed (wrong password).
		exercises, err = readData(key)
		if err != nil {
			return exercises, err
		}

	} else if os.IsNotExist(err) {
		// Since the file didn't exist, we create it.
		_, err := os.Create(DataFile)
		if err != nil {
			fmt.Print("Could not create the file.", err)
		}

		// Specify that the file is empty if not proven otherwise.
		fileStatusEmpty = true
	}

	return exercises, nil
}

// ReadData reads data from an xml file, couldn't be simpler. Unexported.
func readData(key *[32]byte) (XMLData Data, err error) {

	// Open up the xml file that already exists.
	file, err := os.Open(DataFile)
	if err != nil {
		fmt.Print(err)
	}

	// Read the data from the opened file and then check if it is empty.
	encrypted, err := ioutil.ReadAll(file)
	if string(encrypted) == "" {
		fileStatusEmpty = true
		return XMLData, nil
	} else if err != nil {
		fmt.Print(err)
	}

	// Close the loading of the file now that it has served is perpous.
	go file.Close()

	// Unencrypt the data to the content variable.
	content, err := encrypt.Decrypt(key, encrypted)
	if err != nil {
		return XMLData, err
	}

	// Unmarshal the xml data in to our Data struct.
	xml.Unmarshal(content, &XMLData)

	fileStatusEmpty = false
	return XMLData, nil
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
	err = ioutil.WriteFile(DataFile, encrypt.Encrypt(key, file), 0644)
	if err != nil {
		fmt.Print(err)
	}
}

// Format formats the latest updated data in the Data struct to display information.
func (d *Data) Format(i int) (output string) {
	if d.Exercise[i].Reps == 0 && d.Exercise[i].Sets == 0 {
		output = fmt.Sprintf("\nAt %s on %s, you trained %s. The distance was %.2f kilometers and the exercise lasted for %v minutes.\nThis resulted in an average speed of %.3f km/h.\n",
			d.Exercise[i].Clock, d.Exercise[i].Date, d.Exercise[i].Activity, d.Exercise[i].Distance, d.Exercise[i].Time, (d.Exercise[i].Distance/d.Exercise[i].Time)*60)
	} else if d.Exercise[i].Distance == 0 {
		output = fmt.Sprintf("\nAt %s on %s, you trained %s for %v minutes. You did %v sets with %v reps each.\n",
			d.Exercise[i].Clock, d.Exercise[i].Date, d.Exercise[i].Activity, d.Exercise[i].Time, d.Exercise[i].Sets, d.Exercise[i].Reps)
	} else {
		output = fmt.Sprintf("\nAt %s on %s, you trained %s for %v minutes. The distance was %.2f kilometers and you did %v sets with %v reps each.\n",
			d.Exercise[i].Clock, d.Exercise[i].Date, d.Exercise[i].Activity, d.Exercise[i].Time, d.Exercise[i].Distance, d.Exercise[i].Sets, d.Exercise[i].Reps)
	}

	if d.Exercise[i].Comment != "" {
		output += fmt.Sprintf("Comment: %s\n", d.Exercise[i].Comment)
	}

	return output
}

// Delete removes all content in the case of a user wanting to start fresh.
func (d *Data) Delete() {
	// Set the XMLData variable to be empty.
	d = &Data{}

	// Remove the file concurrently to not be slowed down by disk io.
	go func() {
		err := os.Remove(DataFile)
		if err != nil {
			fmt.Print(err)
		}
	}()
}
