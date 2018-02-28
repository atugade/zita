package main

import (
	"log"
	"os"

	"github.com/nlopes/slack"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	Log *log.Logger

	debug = kingpin.Flag("debug", "Enable debug mode.").Bool()
)

func get_slack_token() string {
	return os.Getenv("SLACK_TOKEN")
}

func slack_client_init(token string) *slack.Client {
	return slack.New(token)
}

func slack_init(client *slack.Client) *slack.RTM {

	rtm := client.NewRTM()
	go rtm.ManageConnection()

	return rtm

}

func log_init(logpath string) {
	println("LogFile: " + logpath)
	file, err := os.Create(logpath)

	if err != nil {
		panic(err)
	}

	Log = log.New(file, "zita: ", log.Lshortfile|log.LstdFlags)
}

func main() {

	kingpin.Version("0.0.1")
	kingpin.Parse()

	config, _ := load_config("config.toml")
	log_init("zita.log")
	slack.SetLogger(Log)

	token := get_slack_token()
	client := slack_client_init(token)

	if *debug {
		client.SetDebug(true)
	}

	rtm := slack_init(client)

	event_loop(rtm, config)
}
