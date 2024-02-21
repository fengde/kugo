package command

import "fmt"

func Exec(cmd string, args ...string) error {
	v, ok := commands[cmd]
	if !ok {
		return fmt.Errorf("command '%s' is unsupported!", cmd)
	}

	return v.Exec(args...)
}
