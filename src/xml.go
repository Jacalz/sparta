package main

import "encoding/xml"

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
