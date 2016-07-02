package init

import (
	"github.com/itpkg/deploy"
	"github.com/urfave/cli"
)

func init() {
	deploy.Register(cli.Command{
		Name:    "init",
		Aliases: []string{"i"},
		Usage:   "init default config files",
		Action: deploy.Action(func(*cli.Context, *deploy.Stage) error {
			return nil
		}),
	})
}
