package main

import (
	"bufio"
	"fmt"
	"os"
	"flag"

	"github.com/nlopes/slack"
)

func main() {
	slackToken := flag.String("token", "", "slack token")
	slackChannel := flag.String("channel", "#random", "slack post message channelID or channelName")
	flag.Parse()

	var message string
	message += "```"
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message += scanner.Text()
		message += "\n"
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(2)
	}

	message += "```"

	slackClient := slack.New(*slackToken)
	_, _, err := slackClient.PostMessage(*slackChannel, message, slack.NewPostMessageParameters())
	if err != nil {
		fmt.Fprintln(os.Stderr, "slack post message error:", err)
		os.Exit(2)
	}
}

