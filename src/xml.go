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

// CheckData does relevant checks around our data file.
func CheckData() (exercises *Data, empty bool) {

	configDir, _ := os.UserConfigDir()
	file := filepath.Join(configDir, "sparta", "exercises.xml")

	// Check if the user has a data file directory.
	if _, err := os.Stat(file); err == nil { // The folder does exist.
		exercises, empty = ReadData(file)
	} else if os.IsNotExist(err) { // The file doesn't exist, we should create it.

		// Check if the directory exists. If now, we create it.
		if _, err := os.Stat(file); os.IsNotExist(err) {
			os.Mkdir(filepath.Join(configDir, "sparta"), os.ModePerm)
		}

		// We then create the file.
		_, err := os.Create(file)
		if err != nil {
			panic(err)
		}

		// Specify that teh file is empty.
		empty = true
	}

	return exercises, empty
}

// ReadData reads data from an xml file, couldn't be simpler.
func ReadData(filename string) (XMLData *Data, empty bool) {

	// Open up the xml file that already exists.
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Could not find the file.", err)
	}

	if data, _ := ioutil.ReadFile(filename); string(data) == "" {
		return nil, true
	}

	// Make sure to close it also.
	defer file.Close()

	// Read the data from the opened file.
	byteValue, _ := ioutil.ReadAll(file)

	// Unmarshal the xml data in to our Data struct.
	xml.Unmarshal(byteValue, &XMLData)

	return XMLData, false
}

// WriteData writes new exercieses to the data file.
