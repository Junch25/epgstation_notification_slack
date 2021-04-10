# EPGStation-Notification 
EPGStationの録画開始・終了・エラーのSlack通知です。

## CI/CD
[![CircleCI](https://circleci.com/gh/Junch25/epgstation_notification.svg?style=svg)](https://circleci.com/gh/Junch25/epgstation_notification)

## 導入手順
### Slackアプリを作成
URL: https://api.slack.com

参考: https://api.slack.com/lang/ja-jp

### スクリプト配置
```shell script
$ git clone https://github.com/Junch25/epgstation_notification.git
```

### Slackの設定
下記の２つをご自身のものへ変更する。
- SLACK_API_TOKEN
- SLACK_CHANNEL_ID
```shell script
# 配置
$ cd epgstation_notification
$ sudo cp config.exp.yml bin/config.yml

# 編集
$ sudo vim bin/config.yml
slack-config:
  token-key: "SLACK_API_TOKEN"
  channel: "SLACK_CHANNEL_ID"
```

### EPGStationへ設定追加
```shell script
# v2
$ vim /path/to/config/config.yml
---
recordingStartCommand: "/path/to/bin/epgstation-notification rec-start"
recordingFinishCommand: "/path/to/bin/epgstation-notification rec-end"
recordingFailedCommand: "/path/to/bin/epgstation-notification rec-error"
---

# v1
$ vim /path/to/config/config.json
---
{
  "recordedStartCommand": "/path/to/bin/epgstation-notification rec-start",
  "recordedEndCommand": "/path/to/bin/epgstation-notification rec-end",
  "recordedFailedCommand": "/path/to/bin/epgstation-notification rec-error"
}
---
```
### EPGStation再起動
```shell script
$ sudo pm2 restart epgstation
```

### Build

`EPGStation-Notification`というバイナリファイルできればOK
```shell script
$ cd epgstation_notification
$ GOOS=linux GOARCH=amd64 go build -o "bin/epgstation-notification" main.go
$ ls bin
  epgstation-notification
```

## License
[Apache License 2.0](https://github.com/Junch25/rec_notice_slack_dev/blob/main/LICENSE)
