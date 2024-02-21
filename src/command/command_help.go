package command

import (
	"fmt"
	"sort"
)

type HelpCommand struct{}

func (self HelpCommand) Cmd() string {
	return "help"
}

func (self HelpCommand) Exec(args ...string) error {
	var keys []string
	for key := range commands {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(commands[key].Help())
	}

	return nil
}

func (self HelpCommand) Help() string {
	return "kugo help: show all command supported."
}
