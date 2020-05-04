package sync

import (
	"context"
	"sort"

	"github.com/Jacalz/sparta/internal/file"

	"fyne.io/fyne"
	"github.com/psanford/wormhole-william/wormhole"
)

// StartSync starts up the server on the local network and returns it so we can call shutdown.
func StartSync(synccode chan string, username string, key *[]byte) error {
	// Create the wormhole client.
	var c wormhole.Client

	f, err := file.OpenUserFile(username)
	if err != nil {
		return err
	}

	defer f.Close() // #nosec - We are not writing to the file.

	// Decrypt the data to the content variable.
	content, err := file.ReadEncrypted(f, key)
	if err != nil {
		return err
	}

	// Send the file in the background.
	code, status, err := c.SendText(context.Background(), string(content))
	if err != nil {
		fyne.LogError("Error on sending the file to share", err)
		return err
	}

	// Send the code down the drain so it can be shown inside the ui.
	synccode <- code

	// Handle the status of the sharing.
	if s := <-status; s.Error != nil {
		fyne.LogError("Error regarding status of share", err)
		return err
	} else if s.OK {
		return nil
	}

	return nil
}

// Retrieve starts the retrieving process for fetching a shared file.
func Retrieve(stored *file.Data, ReorderExercises chan bool, FirstExercise chan string, key *[]byte, code, username string) error {
	// Create the wormhole client.
	var c wormhole.Client

	// Receive the data from wormhole sharing.
	data, err := c.Receive(context.Background(), code)
	if err != nil {
		fyne.LogError("Error on receiving", err)
		return err
	}

	// received will store all fetched data.
	received, err := file.ReadEncryptedJSON(data, key)
	if err != nil {
		fyne.LogError("Error on reading the encrypted JSON data", err)
		return err
	}

	// Variable for keeping track of compare values in for loops.
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
		return nil
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
		go stored.Write(key, username)
	}

	return nil
}
