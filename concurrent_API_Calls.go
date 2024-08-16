package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// Struct to hold the post data
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// Function to make an API call
func apiCall(id int, input string) (string, error) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var post Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Post %d: %s", id, post.Title), nil
}

func main() {
	start := time.Now()

	// Number of API calls to make
	numCalls := 5

	var wg sync.WaitGroup
	results := make([]string, numCalls)
	resultChan := make(chan string)
	errorChan := make(chan error)

	// Launch goroutine to collect results
	go func() {
		for i := 0; i < numCalls; i++ {
			select {
			case result := <-resultChan:
				results[i] = result
			case err := <-errorChan:
				results[i] = fmt.Sprintf("Error: %v", err)
			}
		}
	}()

	// First API call
	wg.Add(1)
	go func() {
		defer wg.Done()
		result, err := apiCall(1, "initial data")
		if err != nil {
			errorChan <- err
		} else {
			resultChan <- result
		}
	}()

	// Subsequent API calls
	for i := 1; i < numCalls; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// Wait for the previous result
			var prevResult string
			select {
			case prevResult = <-resultChan:
			case err := <-errorChan:
				errorChan <- err
				return
			}
			// Make the API call
			result, err := apiCall(id+1, prevResult)
			if err != nil {
				errorChan <- err
			} else {
				resultChan <- result
			}
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(resultChan)
	close(errorChan)

	// Print results
	for i, result := range results {
		fmt.Printf("API %d result: %s\n", i+1, result)
	}

	fmt.Printf("Total time: %v\n", time.Since(start))
}
