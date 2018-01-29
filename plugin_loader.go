package main

import (
	"bytes"
	"fmt"
	"os"
	"plugin"
)

//type Subcommand interface {
//	Command(s string)
//}

func get_plugin_path(subcommand string) string {
	var p bytes.Buffer

	p.WriteString("plugins/")
	p.WriteString(subcommand)
	p.WriteString(".so")

	return p.String()
}

func load_plugin(p string) (*plugin.Plugin, error) {
	plug, err := plugin.Open(p)

	return plug, err
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
