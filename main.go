package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type kons struct {
	timeSet string
	people  string
}

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("gonum", "https://pkg.go.dev/gonum.org/v1/gonum@v0.12.0/mat#DenseCopyOf"),
		tgbotapi.NewInlineKeyboardButtonData("2", "2"),
		tgbotapi.NewInlineKeyboardButtonData("3", "3"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
	),
)

func main() {
	bot, err := tgbotapi.NewBotAPI(tokenFile())
	if err != nil {
		log.Panic(err)
	}
	var inputkons bool
	var indexs [6]int
	var konsts []kons
	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}

			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}

		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if inputkons == true && update.Message.Chat.UserName == "RB_PRO" {
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "СЧИТАЛ"))
			str := update.Message.Text
			strs := strings.Split(str, "\n")
			for _, val := range strs {
				//b.SendMessage(val, b.chatID, nil)
				konsts = append(konsts, kons{timeSet: val})

			}
			inputkons = false
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "status":
			if update.Message.Chat.UserName == "RB_PRO" {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, admin"))
			} else {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, user"))
			}
		case "start":
			if update.Message.Chat.UserName == "RB_PRO" {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, admin"))
			} else {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, user"))
			}
		case "setkons":
			if update.Message.Chat.UserName == "RB_PRO" {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Введите время для консультаций:"))
				inputkons = true
			} else {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Этот функционал не для Вас"))
			}
		case "tecalkonsall":
			if update.Message.Chat.UserName == "RB_PRO" {
				var str string
				for ind, val := range konsts {
					str += strconv.Itoa(ind+1) + ". " + val.timeSet + " - " + val.people + "\n"
				}
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, str))
			} else {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Этот функционал не для Вас"))
			}
		case "kons":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			var inlinekons [6]kons
			var cout int
			for i := len(konsts) - 1; i >= 0; i-- {
				if konsts[i].people == "" {
					inlinekons[cout] = konsts[i]
					indexs[cout] = i
					cout++
				}
				if cout == 6 {
					break
				}
			}

			fmt.Println(inlinekons)

			numericKeyboard2 := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(inlinekons[0].timeSet, strconv.Itoa(indexs[0])+". Вы записались на консультацию "+inlinekons[0].timeSet),
					tgbotapi.NewInlineKeyboardButtonData(inlinekons[1].timeSet, strconv.Itoa(indexs[1])+". Вы записались на консультацию "+inlinekons[1].timeSet),
					tgbotapi.NewInlineKeyboardButtonData(inlinekons[2].timeSet, strconv.Itoa(indexs[2])+". Вы записались на консультацию "+inlinekons[2].timeSet),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(inlinekons[3].timeSet, strconv.Itoa(indexs[3])+". Вы записались на консультацию "+inlinekons[3].timeSet),
					tgbotapi.NewInlineKeyboardButtonData(inlinekons[4].timeSet, strconv.Itoa(indexs[4])+". Вы записались на консультацию "+inlinekons[4].timeSet),
					tgbotapi.NewInlineKeyboardButtonData(inlinekons[5].timeSet, strconv.Itoa(indexs[5])+". Вы записались на консультацию "+inlinekons[5].timeSet),
				),
			)

			msg.ReplyMarkup = numericKeyboard2

			// Send the message.
			if _, err = bot.Send(msg); err != nil {
				panic(err)
			}
		default:
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "I don't know that command"))
		}

	}
}
