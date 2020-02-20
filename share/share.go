package share

import (
	"sparta/file"
	"sparta/file/encrypt"

	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/psanford/wormhole-william/wormhole"
)

// StartSharing starts up the server on the local network and returns it so we can call shutdown.
func StartSharing(sharecode chan string, finished chan struct{}) {
	// Create the wormhole client.
	var c wormhole.Client

	// Open up the file.
	f, err := os.Open(path.Join(file.ConfigDir(), "exercises.json"))
	if err != nil {
		fmt.Printf("Opening file: %s\n", err)
		return
	}

	// Defer the closing of the file.
	defer f.Close()

	// Send the file in the background.
	code, status, err := c.SendFile(context.Background(), path.Join(file.ConfigDir(), "sparta", "exercises.json"), f)
	if err != nil {
		fmt.Printf("Could not share file: %s\n", err)
		return
	}

	// Send the code down the drain so it can be shown inside the ui.
	sharecode <- code

	// Handle the status of the sharing.
	if s := <-status; s.Error != nil {
		fmt.Printf("Sharing returned an error: %s\n", s.Error)
		return
	} else if s.OK {
		close(finished)
	}
}

// Retrieve starts the retrieving process for fetching a shared file.
func Retrieve(stored *file.Data, newAddedExercise chan string, key *[32]byte, code string) {
	// Create the wormhole client.
	var c wormhole.Client

	// Receive the data from wormhole sharing.
	data, err := c.Receive(context.Background(), code)
	if err != nil {
		fmt.Printf("Receiving content returned: %s\n", err)
		return
	}

	// Read the data from the http response.
	encrypted, err := ioutil.ReadAll(data)
	if err != nil {
		fmt.Printf("Could not read from file: %s\n", err)
		return
	}

	// received will store all fetched data.
	received := file.Data{}

	// Decrypt the content usign the decrypt function.
	content, err := encrypt.Decrypt(key, encrypted)
	if err != nil {
		fmt.Printf("Could not decrypt content: %s\n", err)
		return
	}

	// unamrchal the content to get the json data from it.
	err = json.Unmarshal(content, &received)
	if err != nil {
		fmt.Printf("Could not unmarshal json: %s\n", err)
		return
	}

	// Variables for keeping track of compare value sin for loops.
	exists := false

	// Compare the two sets of data and add any non existing data.
	for _, fetched := range received.Exercise {

		// Make an asumption that it does not exist.
		exists = false

		// For each fetched item, we loop through and see if it can be found inside the stuff we already have.
		for _, stored := range stored.Exercise {
			if fetched == stored {
				exists = true
				break
			}
		}

		// If the fetched item does not exist, we make sure to add it.
		if !exists {
			stored.Exercise = append(stored.Exercise, fetched)
			newAddedExercise <- stored.Format(len(stored.Exercise) - 1)
		}
	}

	// Write the updated data to our data file.
	go stored.Write(key)
}
