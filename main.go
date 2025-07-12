package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get environment variables
	botToken := os.Getenv("SLACK_BOT_TOKEN")
	channelID := os.Getenv("CHANNEL_ID")

	// Validate required environment variables
	if botToken == "" {
		log.Fatal("SLACK_BOT_TOKEN is required")
	}
	if channelID == "" {
		log.Fatal("CHANNEL_ID is required")
	}

	// Initialize Slack API client
	api := slack.New(botToken)

	channelArr := []string{channelID}

	fileArr := []string{`C:\Users\........YOUR_FILE_PATH.pdf`} // Replace with actual file paths like path/to/your/file.txt

	// Log successful initialization
	// log.Printf("Slack bot initialized successfully for channel: %s", channelID)

	for i := 0; i < len(fileArr); i++ {
		filePath := fileArr[i]
		if filePath == "" {
			log.Printf("Skipping empty file path at index %d", i)
			continue
		}

		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     filePath,
		}

		file, err := api.UploadFile(params)
		if err != nil {
			log.Printf("Error uploading file %s: %v", filePath, err)
			continue
		}

		log.Printf("File uploaded successfully: %s (ID: %s)", file.Name, file.ID)
	}
}
