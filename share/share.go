package share

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sparta/file"
	"sparta/file/encrypt"

	"context"
	"net/http"
	"path/filepath"
	"time"
)

// StartServer starts up the server on the local network and returns it so we can call shutdown.
func StartServer(shutdown chan bool) {
	// Start up the handling of the encrypted exercises on the network.
	http.HandleFunc("/shared-data/encrypted-exercises", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(file.Config(), "sparta", "exercises.json"))
	})

	// Set up the server with an adress port.
	srv := &http.Server{Addr: ": 6230"}

	// Set up a separate goroutine to handle out server shutdown.
	go func() {
		<-shutdown

		// Set up a context to make sure that the server has time to shut down gracefully.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		// Tell the server to not keep connections alive.
		srv.SetKeepAlivesEnabled(false)

		// Call the shutdown function with the given context.
		if err := srv.Shutdown(ctx); err != nil {
			fmt.Printf("Could not gracefully shutdown the server: %v\n", err)
		}
	}()

	// Start listening and serving the file using ther server.
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("Listening returned: %s\n", err)
		return
	}
}

// Retrieve starts the retrieving process for fetching a shared file.
func Retrieve(storedData *file.Data, newAddedExercise chan string, key *[32]byte) {
	// Start by trying to download the data over http.
	resp, err := http.Get("http://localhost:6230/shared-data/encrypted-exercises")
	if err != nil {
		fmt.Printf("Fetching of shared data returned: %s\n", err)
		return
	}

	// Defer the closing of the response body.
	defer resp.Body.Close()

	// Read the data from the http response.
	encrypted, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Could not read from file: %s\n", err)
		return
	}

	// fetchedData will store all fetched data.
	fetchedData := file.Data{}

	// Decrypt the content usign the decrypt function.
	content, err := encrypt.Decrypt(key, encrypted)
	if err != nil {
		fmt.Printf("Could not decrypt content: %s\n", err)
		return
	}

	// unamrchal the content to get the json data from it.
	err = json.Unmarshal(content, &fetchedData)
	if err != nil {
		fmt.Printf("Could not unmarshal json: %s\n", err)
		return
	}

	// Variables for keeping track of compare value sin for loops.
	exists := false

	// Compare the two sets of data and add any non existing data.
	for _, fetched := range fetchedData.Exercise {

		// Make an asumption that it does not exist.
		exists = false

		// For each fetched item, we loop through and see if it can be found inside the stuff we already have.
		for _, stored := range storedData.Exercise {
			if fetched == stored {
				exists = true
				break
			}
		}

		// If the fetched item does not exist, we make sure to add it.
		if !exists {
			storedData.Exercise = append(storedData.Exercise, fetched)
			newAddedExercise <- storedData.Format(len(storedData.Exercise) - 1)
		}
	}

	// Write the updated data to our data file.
	go storedData.Write(key)
}
