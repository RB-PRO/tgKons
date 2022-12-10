package main

import (
	"log"
	"strings"

	"github.com/NicoNex/echotron/v3"
)

const adminID string = "RB_PRO"

type bot struct {
	chatID int64
	echotron.API
}

func newBot(chatID int64) echotron.Bot {
	return &bot{
		chatID,
		echotron.NewAPI(tokenFile()),
	}
}

// Страуктура консультации
type object struct {
	timeSet string
	people  string
}

var konsts []object

// This method is needed to implement the echotron.Bot interface.
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
				b.SendMessage(val, b.chatID, nil)
				konsts = append(konsts, object{timeSet: val})
			}
		} else {
			b.SendMessage("Вам не доступен этот функционал", b.chatID, nil)
		}

	}
	if update.Message.Text == "/kons" {
		b.SendMessage("Выберите время для записи на консультацию.", b.chatID, nil)
	}
}

func main() {
	// This is the entry point of echotron library.
	dsp := echotron.NewDispatcher(tokenFile(), newBot)
	log.Println(dsp.Poll())
}
