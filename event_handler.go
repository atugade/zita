package main

import (
	"fmt"
	"os"
	//"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/nlopes/slack"
)

type Command interface {
	Command([]string)
}

func event_loop(rtm *slack.RTM, config *tomlConfig) {

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
				go process_message_event(rtm, ev, config)

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

func process_message_event(rtm *slack.RTM, ev *slack.MessageEvent, config *tomlConfig) int{
	if !is_authorized(ev.User, config) {
		return 1
	}

	fmt.Printf("Message: %v\n", ev)
	spew.Dump(ev)

	a := string_to_list(ev.Text)

	// pops the userid the message was addressed to
	a, _ = pop_list(a)

	// pops the subcommand name
	a, subcommand := pop_list(a)

	plugpath := get_plugin_path(subcommand)

	if _, err := os.Stat(plugpath); err == nil {
		plug, _ := load_plugin(plugpath)
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
	}

	return 0

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
