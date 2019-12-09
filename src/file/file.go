package file

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Data has the xml data for the initial data tag and then incorporates the Exercise struct.
type Data struct {
	XMLName xml.Name `xml:"data"`
	// Add a filed for when the file was last updated.
	Exercise []Exercise `xml:"exercise"`
}

// Exercise keeps track of the data for each exercise that the user has done.
type Exercise struct {
	Date     string  `xml:"date"`
	Clock    string  `xml:"clock"`
	Activity string  `xml:"activity"`
	Distance float64 `xml:"distance"`
	Time     float64 `xml:"time"`
	//Reps     int     `xml:"reps"`
	//Sets     int     `xml:"sets"`
	//Comment string `xml:"comment"`
}

var configDir, _ = os.UserConfigDir()

// DataFile specifies the loacation of our data file.
var DataFile string = filepath.Join(configDir, "sparta", "exercises.xml")

// Check does relevant checks around our data file.
func Check() (exercises Data, empty bool) {

	// Check if the user has a data file directory.
	if _, err := os.Stat(DataFile); err == nil { // The folder does exist.
		exercises, empty = readData()
	} else if os.IsNotExist(err) { // The file doesn't exist, we should create it.

		// Check if the directory exists. If now, we create it.
		if _, err := os.Stat(DataFile); os.IsNotExist(err) {
			os.Mkdir(filepath.Join(configDir, "sparta"), os.ModePerm)
		}

		// We then create the file.
		_, err := os.Create(DataFile)
		if err != nil {
			fmt.Print("Could not create the file.", err)
		}

		// Specify that teh file is empty.
		empty = true
	}

	return exercises, empty
}

// ReadData reads data from an xml file, couldn't be simpler. Unexported.
func readData() (XMLData Data, empty bool) {

	// Open up the xml file that already exists.
	file, err := os.Open(DataFile)
	if err != nil {
		fmt.Print("Could not find the file.", err)
	}

	if data, _ := ioutil.ReadFile(DataFile); string(data) == "" {
		return XMLData, true
	}

	// Make sure to close it also.
	defer file.Close()

	// Read the data from the opened file.
	byteValue, _ := ioutil.ReadAll(file)

	// Unmarshal the xml data in to our Data struct.
	xml.Unmarshal(byteValue, &XMLData)

	return XMLData, false
}

// Write writes new exercieses to the data file.
func Write(exercises Data) {
	//Marchal the xml content in to a file variable.
	file, err := xml.Marshal(exercises)
	if err != nil {
		fmt.Print("Could not marchal the data.", err)
	}

	// Write to the file.
	_ = ioutil.WriteFile(DataFile, file, 0644)

	// Just some casual debuging thrown in to the mix:
	file2, err := xml.MarshalIndent(exercises, "  ", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(file2))
}
