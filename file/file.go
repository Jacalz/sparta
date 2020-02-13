package file

import (
	"sparta/file/encrypt"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

// Data has the xml data for the initial data tag and then incorporates the Exercise struct.
type Data struct {
	LastUpdated time.Time  `json:"updated"`
	Exercise    []Exercise `json:"exercise"`
}

// Exercise keeps track of the data for each exercise that the user has done.
type Exercise struct {
	Date     string  `json:"date"`
	Clock    string  `json:"clock"`
	Activity string  `json:"activity"`
	Distance float64 `json:"distance"`
	Time     float64 `json:"time"`
	Reps     int     `json:"reps"`
	Sets     int     `json:"sets"`
	Comment  string  `json:"comment"`
}

// fileStatusEmpty defines if the file is empty or not.
var fileStatusEmpty bool

// zeroData is a variable containing an empty Data struct.
var zeroData = &Data{}

// Empty returns if we have a config file or not.
func Empty() bool {
	return fileStatusEmpty
}

// SetNonEmpty tells us that the file is not empty anymore.
func SetNonEmpty() {
	fileStatusEmpty = false
}

// Config returns the config directory and handles the error accordingly.
func Config() (dir string) {
	// Workaround having golang 1.12.x in the cross compiling tool.
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
}

// Check does relevant checks around our data file.
func Check(key *[32]byte) (exercises Data, err error) {

	// Check if the user has a data file.
	if _, err := os.Stat(filepath.Join(Config(), "sparta", "exercises.json")); err == nil {
		// The file exists and we read the data. Return error if decryption failed (wrong password).
		exercises, err = readData(key)
		if err != nil {
			return exercises, err
		}

	} else if os.IsNotExist(err) {
		// Since the file didn't exist, we create it.
		_, err := os.Create(filepath.Join(Config(), "sparta", "exercises.json"))
		if err != nil {
			fmt.Print("Could not create the file.", err)
		}

		// Specify that the file is empty if not proven otherwise.
		fileStatusEmpty = true
	}

	return exercises, nil
}

// ReadData reads data from an xml file, couldn't be simpler. Unexported.
func readData(key *[32]byte) (exercises Data, err error) {
	// Open up the file and it's content.
	data, err := os.Open(filepath.Join(Config(), "sparta", "exercises.json"))
	if err != nil {
		fmt.Print(err)
	}

	// Read the data to extract the encrypted content.
	encrypted, err := ioutil.ReadAll(data)
	if err != nil {
		fmt.Print(err)
	} else if string(encrypted) == "" {
		fileStatusEmpty = true
		return exercises, nil
	}

	// Close the file opening.
	go data.Close()

	// Decrypt the data to the content variable.
	content, err := encrypt.Decrypt(key, encrypted)
	if err != nil {
		return exercises, err
	}

	// Unmarshal the xml data in to our Data struct.
	err = json.Unmarshal(content, &exercises)
	if err != nil {
		fmt.Print(err)
	}

	fileStatusEmpty = false
	return exercises, nil
}

// Write writes new exercieses to the data file.
func (d *Data) Write(key *[32]byte) {
	// Update the section containing the time that our file was last updated.
	d.LastUpdated = time.Now()

	//Marchal the xml content in to a file variable.
	file, err := json.Marshal(d)
	if err != nil {
		fmt.Print(err)
	}

	// Write to the file.
	err = ioutil.WriteFile(filepath.Join(Config(), "sparta", "exercises.json"), encrypt.Encrypt(key, file), 0644)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("%s\n", file)
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
	// Clear the data by directing the pointer to point at the zeroData pointer.
	*d = *zeroData

	// Set the file status to be empty.
	fileStatusEmpty = true

	// Remove the file to clear it.
	err := os.Remove(filepath.Join(Config(), "sparta", "exercises.json"))
	if err != nil {
		fmt.Print(err)
	}
}
