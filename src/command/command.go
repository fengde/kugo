package command

func init() {
	register(&HelpCommand{})
	register(&NewCommand{})
	register(&CheckCommand{})
	register(&BuildCommand{})
}

var commands = map[string]ICommand{}

type ICommand interface {
	Cmd() string
	Exec(args ...string) error
	Help() string
}

func register(cmd ICommand) {
	commands[cmd.Cmd()] = cmd
}
