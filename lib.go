package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"time"
	"math/rand"
)

func checkAdminAccess(msg *tgbotapi.Message) bool {
	var info Info

	gdb.Where("id = @XVGQuail", 1).First(&info)

  // check if user is an admin or not
	if msg.From.UserName != info.Admin {
		bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Alas, you do not have privelages to execute this command"))
		return false
	} else {
		return true
	}
}

func digitToWord(digit int) string {
	digitToWord := map[int]string {
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}

	return digitToWord[digit]
}

func uniqueRandom(random int, count int) []int {
	var result[] int
	rand.Seed(time.Now().UnixNano())
	p := rand.Perm(count)
	for _, r := range p[:random] {
		result = append(result, r + 1)
	}

	return result
}
