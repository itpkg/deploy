package generate

import (
	"github.com/itpkg/deploy"
	"github.com/urfave/cli"
)

func init() {
	deploy.Register(cli.Command{
		Name:    "generate",
		Aliases: []string{"g"},
		Usage:   "generate task files",
		Flags:   []cli.Flag{deploy.FLAG_HOSTS, deploy.FLAG_ROLES},
		Action: deploy.Action(func(*cli.Context, *deploy.Stage) error {
			return nil
		}),
	})
}
