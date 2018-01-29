package main

import (
	"fmt"
	"os"
	"plugin"
)

type Subcommand interface {
	Command(s string)
}

func load_plugin(p string) *plugin.Plugin {
	plug, err := plugin.Open(p)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return plug
}

func get_symbol(plug *plugin.Plugin) plugin.Symbol {
	symCommand, err := plug.Lookup("Command")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return symCommand
}

func exec_command(symCommand plugin.Symbol, message string) {
	symCommand.(func())()
}
