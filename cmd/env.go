package cmd

import (
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type cmdEnv struct {
	ChannelType string `envconfig:"CHANNELTYPE"`
	ChannelID   string `envconfig:"CHANNELID"`
	ChannelName string `envconfig:"CHANNELNAME"`
	Name        string `envconfig:"NAME"`
	Description string `envconfig:"DESCRIPTION"`
	Extended    string `envconfig:"EXTENDED"`
	RecPath     string `envconfig:"RECPATH"`
	LogPath     string `envconfig:"LOGPATH"`
}

type cmdCfg struct {
	SlackCfg struct {
		SlackKey string `yaml:"token-key"`
		Channel  string `yaml:"channel"`
	} `yaml:"slack-config"`
}

func loadEnv() (env cmdEnv, err error) {
	if err := envconfig.Process("", &env); err != nil {
		return env, err
	}
	return env, nil
}

func loadCfg() (config cmdCfg, err error) {
	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		return config, err
	}
	err = yaml.UnmarshalStrict([]byte(data), &config)
	if err != nil {
		return config, err
	}
	return config, err
}
