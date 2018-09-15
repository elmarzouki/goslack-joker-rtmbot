# goslack-joker-rtmbot

Just some nerdy jokes slack bot using go and [Slack's RTM API](https://api.slack.com/rtm).

## Setup the environment

1. Setup [slack app](slack/README.md).
2. 2. Clone this repository: `git clone git@github.com:iSuperMostafa/goslack-joker-rtmbot.git`
3. Navigate to the project directory: `cd goslack-joker-rtmbot`
4. Install requirements: `go get ./...`
5. configure your tokens and settings:
   ```bash
   cp slack/config.json slack/config.override.json
   nano slack/config.override.json   # or open the file and edit the variables manually
   ```

## Run the application

Run the bot: `go run main.go`