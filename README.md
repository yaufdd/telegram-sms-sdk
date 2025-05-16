# Telegram SMS SDK

SDK for wotk with Telegram SMS API.

## Install

```bash
go get github.com/yaufdd/telegram-sms-sdk@latest
```

## How to use

```go
package main

import (
    "context"
    "log"
    
    "github.com/yaufdd/telegram-sms-sdk"
)

func main() {
    // Create client telegram-sms-sdk
    client := sdk.NewClient("http://localhost:8080")
    if err != nil {
        log.Fatal(err)
    }

    //Bot registration
    botResp, err := client.RegisterBot(12345, "your_bot_token")
    if err != nil {
        log.Fatalf("Error registration bot: %v", err)
    }
    fmt.Printf("Bot registered: UUID=%s, Hash=%s\n", botResp.UUID, botResp.Hash)


    // Error sending message
    err = client.SendMessage("UUID", "hash", 54321, "text_you_want_to_send")
    if err != nil {
        log.Fatalf("Error sending message: %v", err)
    }
    fmt.Println("Message was sucessfully sent!")
}
```

## Functions

- `NewClient(url string)` - Create new client
- `RegisterBot(telegramID int64, token string)` - Register bot
- `SendMessage(UUID string, hash string, receiverTgID int64, text string)` - Send message

