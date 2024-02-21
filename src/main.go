package main

import (
	"fmt"
	"kugo/command"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("you can enter 'kugo help' to get how to use this tool.")
		return
	}

	cmd := os.Args[1]
	var args []string
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}

	if err := command.Exec(cmd, args...); err != nil {
		fmt.Println(err.Error())
		return
	}
}
