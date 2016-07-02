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
					Name:  cmd.FLAG_STAGE.Name,
					Usage: cmd.FLAG_STAGE.Usage,
				},
				cli.StringFlag{
					Name:  cmd.FLAG_TASK.Name,
					Usage: cmd.FLAG_TASK.Usage,
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
