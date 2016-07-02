package run

import (
	"github.com/itpkg/deploy"
	"github.com/urfave/cli"
)

func init() {
	deploy.Register(cli.Command{
		Name:    "run",
		Aliases: []string{"r"},
		Usage:   "run tasks",
		Flags:   []cli.Flag{deploy.FLAG_HOSTS, deploy.FLAG_ROLES},
		Action: deploy.Action(func(*cli.Context, *deploy.Stage) error {
			return nil
		}),
	})
}
