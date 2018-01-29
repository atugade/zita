package main

import (
	"fmt"
	//"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/nlopes/slack"
)

type Command interface {
	Command([]string)
}

func event_loop(rtm *slack.RTM) {

EventLoop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Println("Event Received: ", msg)
			Log.Println("Event Received: ", msg)

			if *debug {
				spew.Dump(msg)
			}

			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				fmt.Println("Connection counter:", ev.ConnectionCount)

			case *slack.MessageEvent:
				process_message_event(rtm, ev)

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break EventLoop

			case *slack.ConnectionErrorEvent:
				fmt.Printf("No credentials set")
				break EventLoop

			default:
				//Take no action
			}
		}
	}

}

func process_message_event(rtm *slack.RTM, ev *slack.MessageEvent) {
	fmt.Printf("Message: %v\n", ev)

	a := string_to_list(ev.Text)
	a = pop_list(a)
	spew.Dump(a)

	plug := load_plugin("plugins/jenkins.so")

	spew.Dump(plug)

	symCommand := get_symbol(plug)
	//symCommand.(func(string))(ev.Text)
	spew.Dump(symCommand)

	command, ok := symCommand.(Command)

	if !ok {
		fmt.Println("unexpected type from module symbol")
	}

	spew.Dump(command)

	command.Command(a)

	//	info := rtm.GetInfo()
	//	prefix := fmt.Sprintf("<@%s> ", info.User.ID)
	//	//user, _ := rtm.GetUserInfo(ev.User)

	//	if ev.User != info.User.ID && strings.HasPrefix(ev.Text, prefix) {
	//		//reply := fmt.Sprintf("What's up %s!?!?", user.Name)
	//		//rtm.SendMessage(rtm.NewOutgoingMessage(reply, ev.Channel))

	//		// this is how you send a user a private message
	//		params := slack.PostMessageParameters{
	//			Text:     "Testing",
	//			Username: ev.User,
	//			AsUser:   true}
	//		rtm.PostMessage(ev.User, "testing", params)
	//	}
}
