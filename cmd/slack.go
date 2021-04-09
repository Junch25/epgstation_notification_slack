package cmd

import (
	"fmt"
	"github.com/slack-go/slack"
)

func Slack(Icon string, Col string) error {
	env, err := loadEnv()
	if err != nil {
		return err
	}

	api := slack.New(
		env.SlackToken,
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
		env.SlackChannelId,
		slack.MsgOptionAsUser(false),
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
	return nil
}