package base

import (
	"github.com/itpkg/deploy"
	"github.com/itpkg/deploy/cmd"
	"github.com/urfave/cli"
)

func init() {
	deploy.Register(
		cli.Command{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "generate stage/task config files",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  deploy.FLAG_STAGE.Name,
					Usage: deploy.FLAG_STAGE.Usage,
				},
				cli.StringFlag{
					Name:  deploy.FLAG_TASK.Name,
					Usage: deploy.FLAG_TASK.Usage,
				},
			},
			Action: generate,
		},
		cli.Command{
			Name:    "stages",
			Aliases: []string{"st"},
			Usage:   "show all stages",
			Action:  list(cmd.STAGES),
		},
		cli.Command{
			Name:    "tasks",
			Aliases: []string{"tk"},
			Usage:   "show all tasks",
			Action:  list(cmd.TASKS),
		},
	)
}
