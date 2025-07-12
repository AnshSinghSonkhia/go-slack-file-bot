# go-slack-file-bot

[![Go](https://img.shields.io/badge/Go-1.21-blue?logo=go)](https://golang.org/)
[![Slack Bot](https://img.shields.io/badge/Bot-Slack-4A154B?logo=slack)](https://api.slack.com/bot-users)
[![MIT License](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
<!-- [![Build Status](https://img.shields.io/github/actions/workflow/status/yourusername/go-slack-file-bot/ci.yml?branch=main)](https://github.com/yourusername/go-slack-file-bot/actions) -->
<!-- [![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/go-slack-file-bot)](https://goreportcard.com/report/github.com/yourusername/go-slack-file-bot) -->
<!-- [![Issues](https://img.shields.io/github/issues/yourusername/go-slack-file-bot)](https://github.com/yourusername/go-slack-file-bot/issues) -->
<!-- [![Last Commit](https://img.shields.io/github/last-commit/yourusername/go-slack-file-bot)](https://github.com/yourusername/go-slack-file-bot/commits/main) -->


## Overview

**go-slack-file-bot** is a simple Slack bot written in Go that allows users to upload, download, and manage files directly from Slack channels. It leverages Slack's API and Go's concurrency features for efficient file handling.

## Features

- Upload files to Slack channels
- Download files from Slack
- List files shared in a channel
- Delete files (admin only)
- Simple command interface

## Getting Started

### Prerequisites

- Go 1.21 or higher
- A Slack workspace
- A Slack bot token with appropriate permissions (`files:read`, `files:write`, `chat:write`, etc.)

### Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/go-slack-file-bot.git
    cd go-slack-file-bot
    ```

2. Build the bot:
    ```bash
    go build -o slack-file-bot
    ```

3. Set your Slack bot token as an environment variable:
    ```bash
    export SLACK_BOT_TOKEN="xoxb-your-slack-bot-token"
    ```

4. Run the bot:
    ```bash
    ./slack-file-bot
    ```

## Usage

Interact with the bot in your Slack workspace using the following commands:

- `upload <filepath>`: Upload a file to the current channel.
- `download <file_id>`: Download a file from Slack.
- `list`: List recent files in the channel.
- `delete <file_id>`: Delete a file (admin only).

Example:
```
@go-slack-file-bot upload report.pdf
```

## Configuration

You can configure the bot using environment variables:

| Variable           | Description                |
|--------------------|---------------------------|
| SLACK_BOT_TOKEN    | Your Slack bot token      |
| SLACK_APP_TOKEN    | (Optional) App-level token|

## Contributing

Contributions are welcome! Please open issues or submit pull requests for new features, bug fixes, or improvements.

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.

## Acknowledgements

- [Slack API](https://api.slack.com/)
- [Go Slack SDK](https://github.com/slack-go/slack)
