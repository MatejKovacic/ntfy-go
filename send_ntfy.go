package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"os"
)

func sendNtfyMessage(server, topic, message string) error {
	// Construct the URL for the ntfy topic
	url := fmt.Sprintf("%s/%s", server, topic)

	// Create the request body
	body := bytes.NewBufferString(message)

	// Send the HTTP POST request
	resp, err := http.Post(url, "text/plain", body)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	fmt.Println("Message sent successfully!")
	return nil
}

func main() {
	// Define command-line flags
	server := flag.String("s", "", "The ntfy server URL (required)")
	topic := flag.String("t", "", "The ntfy topic to send the message to (required)")
	message := flag.String("m", "", "The message to send (required)")
	flag.Parse()

	// Validate required flags
	if *server == "" || *topic == "" || *message == "" {
		fmt.Println("Error: All flags are required.")
		fmt.Println("Usage: send_ntfy -s <server> -t <topic> -m <message>")
		os.Exit(1)
	}

	// Send the message
	err := sendNtfyMessage(*server, *topic, *message)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
