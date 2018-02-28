package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/davecgh/go-spew/spew"
)

type tomlConfig struct {
	Globals globalsInfo
	Access  accessInfo
}

type globalsInfo struct {
	Logpath string
}

type accessInfo struct {
	Users  []string
	Admins []string
}

func string_to_list(message string) []string {
	return strings.Split(message, " ")
}

func pop_list(message []string) ([]string, string) {
	element := message[0]
	return append(message[:0], message[0+1:]...), element
}

func load_config(confpath string) (*tomlConfig, error) {
	var config tomlConfig

	if _, err := os.Stat(confpath); os.IsNotExist(err) {
		return nil, errors.New("Config file does not exist.")
	}

	if _, err := toml.DecodeFile(confpath, &config); err != nil {
		fmt.Println(err)
		return nil, err
	}

	spew.Dump(config)
	return &config, nil
}
