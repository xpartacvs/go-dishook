# Dishook

A discord webhook package helper. Please read the [reference](https://birdie0.github.io/discord-webhooks-guide/discord_webhook.html)

## Usage

Donwload the package

```bash
go get github.com/xpartacvs/go-dishook@v1.1.1
```

Package in action:

```go
package main

import "github.com/xpartacvs/go-dishook"

func main() {
    webhookUrl := "https://domain.tld/your/discord/webhook/url"

    payload := dishook.Payload{
        Username:  "Dishook",
        AvatarUrl: "https://pbs.twimg.com/profile_images/554798224154701824/mWd3laxO_400x400.png",
        Content:   "Greeting! This message was sent using github.com/xpartacvs/go-dishook",
    }

    _, _ = dishook.Send(webhookUrl, payload)
}
```
