package base

import (
	"fmt"

	"github.com/itpkg/deploy"
	"github.com/urfave/cli"
)

func init() {
	deploy.Register(
		cli.Command{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "generate stage/task config files",
			Flags: []cli.Flag{
				deploy.FLAG_STAGE,
				deploy.FLAG_TASK,
				deploy.FLAG_FORMAT,
			},
			Action: generate,
		},
		cli.Command{
			Name:    "stages",
			Aliases: []string{"S"},
			Usage:   "show all stages",
			Action: func(*cli.Context) error {
				fmt.Println("NAME\tDESCRIPTION")
				for _, s := range deploy.STAGES {
					fmt.Println(s)
				}
				return nil
			},
		},
		cli.Command{
			Name:    "tasks",
			Aliases: []string{"T"},
			Usage:   "show all tasks",
			Action: func(*cli.Context) error {
				fmt.Println("NAME\tDESCRIPTION")
				for _, s := range deploy.STAGES {
					fmt.Println(s)
				}
				return nil
			},
		},
	)
}
