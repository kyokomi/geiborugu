package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/nlopes/slack"
)

var messageFormat = "```\n%s\n```"

type Slack struct {
	*slack.Slack
	channelName string
	dryRun      bool
}

func main() {
	slackToken := flag.String("token", "", "slack token")
	if *slackToken == "" {
		*slackToken = os.Getenv("SLACK_TOKEN")
	}
	slackChannel := flag.String("channel", "", "slack post message channelID or channelName")
	if *slackChannel == "" {
		*slackChannel = os.Getenv("SLACK_CHANNEL")
	}
	dryRun := flag.Bool("dry-run", false, "dry-run")
	flag.Parse()

	slackClient := Slack{
		Slack:       slack.New(*slackToken),
		channelName: *slackChannel,
		dryRun:      *dryRun,
	}

	var message string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message += scanner.Text()
		message += "\n"

		if len(message) >= 5000 {
			slackClient.postMessage(fmt.Sprintf(messageFormat, message))
			message = ""
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(2)
	}
}

func (s Slack) postMessage(message string) error {
	if s.dryRun {
		fmt.Println(message)
		return nil
	}
	_, _, err := s.Slack.PostMessage(s.channelName, message, slack.NewPostMessageParameters())
	return err
}
