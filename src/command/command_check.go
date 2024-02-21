package command

import (
	"os"
	"os/exec"
)

type CheckCommand struct{}

func (self CheckCommand) Cmd() string {
	return "check"
}

func (self CheckCommand) Exec(args ...string) error {
	cmd := exec.Command("golangci-lint", append([]string{"run"}, args...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
	return nil
}

func (self CheckCommand) Help() string {
	return `kugo check: use golangcli-lint to check go code`
}
