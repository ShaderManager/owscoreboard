package main

import (
	"log"
	"strconv"
	"strings"
	"unicode"

	"github.com/gempir/go-twitch-irc/v2"
)

func splitArgs(text string) []string {
	lastQuote := rune(0)
	f := func(c rune) bool {
		switch {
		case c == lastQuote:
			lastQuote = rune(0)
			return true
		case lastQuote != rune(0):
			return false
		case unicode.In(c, unicode.Quotation_Mark):
			lastQuote = c
			return true
		default:
			return unicode.IsSpace(c)
		}
	}

	return strings.FieldsFunc(text, f)
}

func isModeratorOrStreamer(badges map[string]int) bool {
	if _, ok := badges["moderator"]; ok { // Check if user is moderator
		return true
	}
	if _, ok := badges["broadcaster"]; ok { // Check if user is broadcaster
		return true
	}

	return false
}

const twitchWinCmd = "!win"
const twitchLoseCmd = "!lose"
const twitchTieCmd = "!tie"
const twitchResetCmd = "!reset"

func startTwitchPolling() {
	if len(cfg.TwitchChannel) == 0 {
		return
	}

	client := twitch.NewAnonymousClient()
	defer client.Disconnect()

	client.OnPrivateMessage(func(msg twitch.PrivateMessage) {
		if isModeratorOrStreamer(msg.User.Badges) {
			// arg 0: command
			// arg 1 [optional]: role
			// arg 2 [optional]: increment value
			args := splitArgs(msg.Message)

			cmd := ""
			role := ""
			var inc int64
			inc = 1

			if len(args) >= 1 {
				cmd = args[0]
			}

			if len(args) >= 2 {
				role = args[1]
				if len(args) == 3 {
					var err error
					inc, err = strconv.ParseInt(args[2], 10, 32)

					if err != nil {
						inc = 1
					}
				}
			}

			switch {
			case cmd == twitchWinCmd:
				log.Printf("Add %d wins to role %s from %s\n", inc, role, msg.User.DisplayName)
				addWin(role, int(inc))
			case cmd == twitchLoseCmd:
				log.Printf("Add %d loses to role %s from %s\n", inc, role, msg.User.DisplayName)
				addLose(role, int(inc))
			case cmd == twitchTieCmd:
				log.Printf("Add %d ties to role %s from %s\n", inc, role, msg.User.DisplayName)
				addTie(role, int(inc))
			case cmd == twitchResetCmd:
				log.Printf("Reset from %s\n", msg.User.DisplayName)
				resetStats()
			}
		}
	})

	client.Join(cfg.TwitchChannel)

	log.Printf("Connecting to twitch channel %s\n", cfg.TwitchChannel)

	err := client.Connect()

	if err != nil {
		log.Fatal(err)
	}
}
