package cmd

import (
	"github.com/kelseyhightower/envconfig"
)

type cmdEnv struct {
	ChannelType          string `envconfig:"CHANNELTYPE"`
	ChannelID            string `envconfig:"CHANNELID"`
	ChannelName          string `envconfig:"CHANNELNAME"`
	Name                 string `envconfig:"NAME"`
	Description          string `envconfig:"DESCRIPTION"`
	Extended             string `envconfig:"EXTENDED"`
	RecPath              string `envconfig:"RECPATH"`
	LogPath              string `envconfig:"LOGPATH"`
	SlackToken           string `envconfig:"SLACK_API_TOKEN"`
	SlackChannelId       string `envconfig:"SLACK_CHANNEL_ID"`
}

func loadEnv() (env cmdEnv, err error) {
	if err := envconfig.Process("", &env); err != nil{
		return env, err
	}
	return env, nil
}
