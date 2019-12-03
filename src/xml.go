package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Data has the xml data for the initial data tag and then incorporates Exercise.
type Data struct {
	XMLName  xml.Name   `xml:"data"`
	Exercise []Exercise `xml:"exercise"`
}

// Exercise keeps track of the data for each exercise that the user has done.
type Exercise struct {
	Date     string  `xml:"date"`
	Clock    string  `xml:"clock"`
	Distance float64 `xml:"distance"`
	Length   float64 `xml:"length"`
	Activity string  `xml:"activity"`
	Reps     int     `xml:"reps"`
	Sets     int     `xml:"sets"`
	Comment  string  `xml:"comment"`
}

// XMLData contains the data from the xml file.
var XMLData *Data

// CheckDataFile does relevant checks around our data file.
func CheckDataFile() {

	configDir, _ := os.UserConfigDir()
	file := filepath.Join(configDir, "sparta", "exercises.xml")

	// Check if the user has a data file.
	if _, err := os.Stat(file); err == nil { // The file does exist.
		ReadDataFromXML(file)
	} else if os.IsNotExist(err) { // The file doesn't exist, we should create it.
		CreateFile(file)

		// Now read the data
	}
}

// ReadDataFromXML reads data from an xml file, couldn't be simpler.
func ReadDataFromXML(filename string) {

	// Open up the xml file that already exists.
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Could not find the file.", err)
	}

	// Make sure to close it also.
	defer file.Close()

	// Read the data from the opened file.
	byteValue, _ := ioutil.ReadAll(file)

	// Unmarshal the xml data in to our Data struct.
	xml.Unmarshal(byteValue, &XMLData)
}

// CreateFile uses os.Create to make a file.
func CreateFile(filename string) *os.File {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	return file
}
