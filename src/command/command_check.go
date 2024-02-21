package command

type CheckCommand struct{}

func (self CheckCommand) Cmd() string {
	return "check"
}

func (self CheckCommand) Exec(args ...string) error {

	return nil
}

func (self CheckCommand) Help() string {
	return `kugo check: use golangcli-lint to check all go files of the dir and sub dir. (except vendor)`
}
