package share

import (
	"sort"
	"sparta/file"

	"context"
	"fmt"
	"os"
	"path"

	"github.com/psanford/wormhole-william/wormhole"
)

// StartSharing starts up the server on the local network and returns it so we can call shutdown.
func StartSharing(sharecode chan string, errors chan error, finished chan struct{}) {
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
		errors <- err
		fmt.Printf("Could not share file: %s\n", err)
		return
	}

	// Send the code down the drain so it can be shown inside the ui.
	sharecode <- code

	// Handle the status of the sharing.
	if s := <-status; s.Error != nil {
		errors <- s.Error
		fmt.Printf("Sharing returned an error: %s\n", s.Error)
		return
	} else if s.OK {
		close(finished)
	}
}

// Retrieve starts the retrieving process for fetching a shared file.
func Retrieve(stored *file.Data, ReorderExercises chan bool, FirstExercise chan string, errors chan error, key *[32]byte, code string) {
	// Create the wormhole client.
	var c wormhole.Client

	// Receive the data from wormhole sharing.
	data, err := c.Receive(context.Background(), code)
	if err != nil {
		errors <- err
		fmt.Printf("Receiving content returned: %s\n", err)
		return
	}

	// received will store all fetched data.
	received, err := file.ReadEncryptedJSON(data, key)
	if err != nil {
		errors <- err
		fmt.Printf("Parsing JSON from file returned: %s", err)
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
		}
	}

	// Check the length of our combined structs.
	length := len(stored.Exercise)

	// Handle different lengths accordingly to avoid out of bounds index checks.
	switch length {
	case 0:
		return
	case 1:
		FirstExercise <- stored.Format(length - 1)
	default:
		// Sort all old and new data to make sure that new exercises come first.
		sort.Slice(stored.Exercise, func(i, j int) bool {
			return stored.Exercise[i].Time.Before(stored.Exercise[j].Time)
		})

		// Indicate that the whole slice needs to be redisplayed.
		ReorderExercises <- true

		// Write the updated data to our data file.
		go stored.Write(key)
	}
}
