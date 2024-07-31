package main

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

const (
	apiEndpoint = "https://api.openai.com/v1/chat/completions"
)

func main() {
	// Replace YOUR_API_KEY with your actual OpenAI API key
	apiKey := "YOUR_API_KEY"
	client := resty.New()

	response, err := client.R().
		SetAuthToken(apiKey).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"model":     "gpt-3.5-turbo",
			"messages":  []interface{}{map[string]interface{}{"role": "system", "content": "You are a helpful assistant."}},
			"max_tokens": 50,
		}).
		Post(apiEndpoint)

	if err != nil {
		log.Fatalf("Failed to send the request: %v", err)
	}

	fmt.Println(response.String())
}
