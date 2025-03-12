package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func sendNtfyMessage(server, topic, message, username, password string) error {
	// Ensure the server URL does not have a trailing slash
	server = strings.TrimSuffix(server, "/")

	// Construct the URL for the ntfy topic
	url := fmt.Sprintf("%s/%s", server, topic)

	// Create the request body
	body := bytes.NewBufferString(message)

	// Create a new request
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "text/plain")

	// Add basic authentication if credentials are provided
	if username != "" && password != "" {
		req.SetBasicAuth(username, password)
	}

	// Send the request
	resp, err := http.DefaultClient.Do(req)
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
	username := flag.String("u", "", "Username for authentication (optional)")
	password := flag.String("p", "", "Password for authentication (optional)")
	flag.Parse()

	// Validate required flags
	if *server == "" || *topic == "" || *message == "" {
		fmt.Println("Error: All flags are required.")
		fmt.Println("Usage: send_ntfy -s <server> -t <topic> -m <message> [-u <username> -p <password>]")
		os.Exit(1)
	}

	// Send the message
	err := sendNtfyMessage(*server, *topic, *message, *username, *password)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
