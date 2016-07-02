package deploy

import "github.com/urfave/cli"

var commands []cli.Command

//Register register action func
func Register(cmd ...cli.Command) {
	commands = append(commands, cmd...)
}
