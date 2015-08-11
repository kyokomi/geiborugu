package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/nlopes/slack"
)

var messageFormat = "```\n%s```"

func main() {
	slackToken := flag.String("token", "", "slack token")
	slackChannel := flag.String("channel", "", "slack post message channelID or channelName")
	userName := flag.String("name", "", "slack bot name")
	iconURL := flag.String("icon", "", "slack bot icon url")
	dryRun := flag.Bool("dry-run", false, "dry-run")
	flag.Parse()

	if *slackToken == "" {
		*slackToken = os.Getenv("SLACK_TOKEN")
	}
	if *slackChannel == "" {
		*slackChannel = os.Getenv("SLACK_CHANNEL")
	}

	slackClient := Slack{
		Slack:       slack.New(*slackToken),
		channelName: *slackChannel,
		userName:    *userName,
		iconURL:     *iconURL,
		dryRun:      *dryRun,
	}

	var message string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message += scanner.Text()
		message += "\n"

		fmt.Println(message)

		if len(message) >= 5000 {
			slackClient.postMessage(fmt.Sprintf(messageFormat, message))
			message = ""
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(2)
	}

	if message != "" {
		slackClient.postMessage(fmt.Sprintf(messageFormat, message))
	}
}

type Slack struct {
	*slack.Slack
	channelName string
	userName    string
	iconURL     string
	dryRun      bool
}

func (s Slack) newPostMessageParams() slack.PostMessageParameters {
	params := slack.NewPostMessageParameters()
	if s.userName != "" {
		params.Username = s.userName
	}
	if s.iconURL != "" {
		params.IconURL = s.iconURL
	}
	return params
}

func (s Slack) postMessage(message string) error {
	if s.dryRun {
		fmt.Println(message)
		return nil
	}
	_, _, err := s.Slack.PostMessage(s.channelName, message, s.newPostMessageParams())
	return err
}
