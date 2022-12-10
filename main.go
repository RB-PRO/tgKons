package main

import (
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv(tokenFile()))
	if err != nil {
		fmt.Println(err)
	}

	bot.Debug = true

	fmt.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			msg.Text = "I understand /sayhi and /status."
		case "sayhi":
			msg.Text = "Hi :)"
		case "status":
			msg.Text = "I'm ok."
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			fmt.Println(err)
		}
	}
}

/*
func (b *bot) Update(update *echotron.Update) {
	if update.Message.Text == "/start" {
		b.SendMessage("Hello world", b.chatID, nil)
	}
	if update.Message.Text == "/status" {
		if update.Message.Chat.Username == adminID {
			b.SendMessage("Hello, admin", b.chatID, nil)
		} else {
			b.SendMessage("Hello, user", b.chatID, nil)
		}
	}

	if strings.Contains(update.Message.Text, "/setkons") {
		if update.Message.Chat.Username == adminID {
			str := update.Message.Text
			strs := strings.Split(str, "\n")
			strs = strs[1:]
			for _, val := range strs {
				//b.SendMessage(val, b.chatID, nil)
				konsts = append(konsts, object{timeSet: val})

			}
		} else {
			b.SendMessage("Вам не доступен этот функционал", b.chatID, nil)
		}

	}
	if update.Message.Text == "/kons" {
		b.SendMessage("Выберите время для записи на консультацию:", b.chatID, nil)
		var str string
		for ind, val := range konsts {
			str += strconv.Itoa(ind+1) + ". " + val.timeSet + " - " + val.people + "\n"
		}
		b.SendMessage(str, b.chatID, nil)
	}
}
*/
