package file

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/Jacalz/sparta/internal/crypto"

	"fyne.io/fyne/v2"
)

// Data has the xml data for the initial data tag and then incorporates the Exercise struct.
type Data struct {
	Exercise []Exercise `json:"exercise,omitempty"`
}

// Exercise keeps track of the data for each exercise that the user has done.
type Exercise struct {
	Time     time.Time `json:"time,omitempty"`
	Date     string    `json:"date,omitempty"`
	Clock    string    `json:"clock,omitempty"`
	Activity string    `json:"activity,omitempty"`
	Distance float64   `json:"distance,omitempty"`
	Duration float64   `json:"duration,omitempty"`
	Reps     uint32    `json:"reps,omitempty"`
	Sets     uint32    `json:"sets,omitempty"`
	Comment  string    `json:"comment,omitempty"`
}

// zeroData is a variable containing an empty Data struct.
var zeroData = &Data{}

// ReadEncrypted reads the data from a reader, decrypts it and outputs the content.
func ReadEncrypted(r io.Reader, key *[]byte) (content []byte, err error) {
	// Read the data to extract the encrypted content.
	encrypted, err := ioutil.ReadAll(r)
	if err != nil {
		fyne.LogError("Error on reading data from file", err)
		return nil, err
	} else if string(encrypted) == "" {
		return nil, nil
	}

	// Decrypt the data to the content variable.
	content, err = crypto.Decrypt(key, encrypted)
	if err != nil {
		return nil, err
	}

	return content, nil
}

// ReadEncryptedJSON reads encrypted data and then outputs the JSON as a struct.
func ReadEncryptedJSON(r io.Reader, key *[]byte) (exercises Data, err error) {
	content, err := ReadEncrypted(r, key)
	if err != nil {
		return exercises, err
	} else if content == nil {
		return exercises, nil
	}

	// Unmarshal the xml data in to our Data struct.
	err = json.Unmarshal(content, &exercises)
	if err != nil {
		fyne.LogError("Error on JSON unmarshal", err)
		return exercises, err
	}

	return exercises, nil
}

// ReadJSON is an adaptation of ReadEncryptedJSON that handles already unencrypted JSON text.
func ReadJSON(r io.Reader) (exercises Data, err error) {
	text, err := ioutil.ReadAll(r)
	if err != nil {
		fyne.LogError("Error on reading data from the reader", err)
		return exercises, err
	} else if string(text) == "" {
		return exercises, nil
	}

	// Unmarshal the xml data in to our Data struct.
	err = json.Unmarshal(text, &exercises)
	if err != nil {
		fyne.LogError("Error on JSON unmarshal", err)
		return exercises, err
	}

	return exercises, nil
}

// OpenUserFile is used to open up the file for the specified user.
func OpenUserFile(username string) (f *os.File, err error) {
	// Open up the file and it's content.
	f, err = os.Open(filepath.Join(ConfigDir(), username+"-exercises.json")) // #nosec - The username is checked and can not be used outside of the folder.
	if err != nil {
		fyne.LogError("Error on opening the file for the user", err)
		return nil, err
	}

	return f, nil
}

// ReadData reads data from an xml file, couldn't be simpler. Unexported.
func ReadData(key *[]byte, username string) (exercises Data, err error) {
	// Open up the file and it's content.
	f, err := OpenUserFile(username)
	if err != nil {
		return exercises, err
	}

	defer f.Close() // #nosec - We are not writing to the file.

	// Read the JSON data from the encrypted file.
	exercises, err = ReadEncryptedJSON(f, key)
	if err != nil {
		return exercises, err
	}

	return exercises, nil
}

// Write writes new exercises to the data file.
func (d *Data) Write(key *[]byte, username string) {
	//Marchal the xml content in to a file variable.
	file, err := json.Marshal(d)
	if err != nil {
		fyne.LogError("Error on marshalling of json", err)
		return
	}

	// Write to the file.
	err = ioutil.WriteFile(filepath.Join(ConfigDir(), username+"-exercises.json"), crypto.Encrypt(key, file), 0600)
	if err != nil {
		fyne.LogError("Error on writing to file", err)
		return
	}
}

// Format formats the latest updated data in the Data struct to display information.
func (d *Data) Format(i int) (output string) {
	if d.Exercise[i].Reps == 0 && d.Exercise[i].Sets == 0 {
		output = fmt.Sprintf("\nAt %s on %s, you trained %s. The distance was %.2f kilometers and the exercise lasted for %v minutes.\nThis resulted in an average speed of %.3f km/h.\n",
			d.Exercise[i].Clock, d.Exercise[i].Date, d.Exercise[i].Activity, d.Exercise[i].Distance, d.Exercise[i].Duration, (d.Exercise[i].Distance/d.Exercise[i].Duration)*60)
	} else if d.Exercise[i].Distance == 0 {
		output = fmt.Sprintf("\nAt %s on %s, you trained %s for %v minutes. You did %v sets with %v reps each.\n",
			d.Exercise[i].Clock, d.Exercise[i].Date, d.Exercise[i].Activity, d.Exercise[i].Duration, d.Exercise[i].Sets, d.Exercise[i].Reps)
	} else {
		output = fmt.Sprintf("\nAt %s on %s, you trained %s for %v minutes. The distance was %.2f kilometers and you did %v sets with %v reps each.\n",
			d.Exercise[i].Clock, d.Exercise[i].Date, d.Exercise[i].Activity, d.Exercise[i].Duration, d.Exercise[i].Distance, d.Exercise[i].Sets, d.Exercise[i].Reps)
	}

	if d.Exercise[i].Comment != "" {
		output += fmt.Sprintf("Comment: %s\n", d.Exercise[i].Comment)
	}

	return output
}

// Delete removes all content in the case of a user wanting to start fresh.
func (d *Data) Delete(username string) {
	// Clear the data by directing the pointer to point at the zeroData pointer.
	*d = *zeroData

	// Remove the file to clear it.
	err := os.Remove(filepath.Join(ConfigDir(), username+"-exercises.json"))
	if err != nil {
		fyne.LogError("Error on removing the json file", err)
	}
}
