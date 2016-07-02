package run

import (
	"github.com/itpkg/deploy"
	"github.com/itpkg/deploy/cmd"
	"github.com/urfave/cli"
)

func init() {
	deploy.Register(cli.Command{
		Name:    "run",
		Aliases: []string{"r"},
		Usage:   "run tasks",
		Flags:   []cli.Flag{cmd.FLAG_HOSTS, cmd.FLAG_ROLES, cmd.FLAG_STAGE, cmd.FLAG_TASK},
		Action:  cmd.Action(run),
	})
}
