package cmd

import (
	"fmt"
	"github.com/slack-go/slack"
)

func Slack(Icon string, Col string) error {
	env, err1 := loadEnv()
	if err1 != nil {
		return err1
	}
	config, err2 := loadCfg()
	if err2 != nil {
		return err2
	}
	api := slack.New(
		config.SlackCfg.SlackKey,
		slack.OptionDebug(true),
	)
	attachment := slack.Attachment{
		Fallback: Icon + env.Name,
		Color:    Col,
		Title:    Icon + env.Name,
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
		config.SlackCfg.Channel,
		slack.MsgOptionAsUser(false),
		slack.MsgOptionAttachments(attachment),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
	return nil
}
