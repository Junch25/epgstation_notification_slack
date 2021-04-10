package cmd

import (
	"fmt"
	"github.com/slack-go/slack"
	"golang.org/x/tools/go/ssa/interp/testdata/src/os"
)

func Slack(Icon string, Col string) error {
	env, err := loadEnv()
	if err != nil {
		return err
	}

	SlackToken := os.Getenv("SLACK_API_TOKEN")
	SlackChannel := os.Getenv("SLACK_CHANNEL_ID")
	api := slack.New(
		SlackToken,
		slack.OptionDebug(true),
		)
	attachment := slack.Attachment{
		Fallback: Icon + env.Name,
		Color:   Col,
		Title: Icon + env.Name,
		Fields: []slack.AttachmentField{
			{
				Title: "Channel",
				Value: env.ChannelName + "/" + env.ChannelType,
				Short: false,
			},
			{
				Title: "Description",
				Value: env.Description,
				Short: false,
			},
			{
				Title: "Extended",
				Value: env.Extended,
				Short: false,
			},
			{
				Title: "RecPath",
				Value: env.RecPath,
				Short: false,
			},
			{
				Title: "LogPath",
				Value: env.LogPath,
				Short: false,
			},
		},
	}
	channelID, timestamp, err := api.PostMessage(
		SlackChannel,
		slack.MsgOptionAsUser(false),
		slack.MsgOptionAttachments(attachment),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
	return nil
}
