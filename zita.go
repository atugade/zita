package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/nlopes/slack"
	"github.com/davecgh/go-spew/spew"
)

func event_loop(rtm *slack.RTM) {

EventLoop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			//fmt.Println("Event Received: ", msg)
                        spew.Dump(msg)
			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				fmt.Println("Connection counter:", ev.ConnectionCount)

			case *slack.MessageEvent:
				fmt.Printf("Message: %v\n", ev)
				info := rtm.GetInfo()
				prefix := fmt.Sprintf("<@%s> ", info.User.ID)
				//user, _ := rtm.GetUserInfo(ev.User)

				if ev.User != info.User.ID && strings.HasPrefix(ev.Text, prefix) {
					//reply := fmt.Sprintf("What's up %s!?!?", user.Name)
					//rtm.SendMessage(rtm.NewOutgoingMessage(reply, ev.Channel))

					// this is how you send a user a private message
					params := slack.PostMessageParameters{
						Text:     "Testing",
						Username: ev.User,
						AsUser:   true}
					rtm.PostMessage(ev.User, "testing", params)
				}

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break EventLoop

			default:
				//Take no action
			}
		}
	}

}

func slack_init() *slack.RTM {

	token := os.Getenv("SLACK_TOKEN")
	api := slack.New(token)
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	return rtm

}

func main() {

	rtm := slack_init()
	event_loop(rtm)
}
