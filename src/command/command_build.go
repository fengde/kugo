package command

type BuildCommand struct{}

func (self BuildCommand) Cmd() string {
	return "build"
}

func (self BuildCommand) Exec(args ...string) error {

	return nil
}

func (self BuildCommand) Help() string {
	return "kugo build: build project. it will build executor to ./target/"
}
