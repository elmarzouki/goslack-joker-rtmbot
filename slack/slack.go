package slack

import (
	"../joker"
	"encoding/json"
	"fmt"
	"github.com/nlopes/slack"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type slackConfig struct {
	AccessToken string `json:"access_token"`
	ChannelID   string `json:"channel_id"`
}

type config struct {
	Slack slackConfig `json:"slack"`
}

func getConfig() config {
	// Open our jsonFile
	baseDir, _ := os.Getwd()
	jsonFile, err := os.Open(baseDir + "/slack/config.override.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var config config

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	err = json.Unmarshal(byteValue, &config)

	return config
}

func ConnectSlack() {
	// get configurations
	config := getConfig()
	slackClient := slack.New(config.Slack.AccessToken)
	logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
	slackClient.SetDebug(true)

	rtm := slackClient.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		//fmt.Print("Event Received: ")
		switch ev := msg.Data.(type) {

		case *slack.MessageEvent:
			if strings.Contains(ev.Msg.Text, "joke") {
				// get a joke and send it
				message := joker.GetJoke().Joke
				rtm.SendMessage(rtm.NewOutgoingMessage(message, config.Slack.ChannelID))
			} else {
				message := "why so serious, just ask me to tell you a nerdy joke!"
				rtm.SendMessage(rtm.NewOutgoingMessage(message, config.Slack.ChannelID))
			}

		case *slack.RTMError:
			fmt.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials")
			return

		}
	}
}
