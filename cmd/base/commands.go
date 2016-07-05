package base

import (
	"fmt"

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
				cmd.FLAG_STAGE,
				cmd.FLAG_TASK,
				cmd.FLAG_FORMAT,
			},
			Action: generate,
		},
		cli.Command{
			Name:    "stages",
			Aliases: []string{"S"},
			Flags:   []cli.Flag{cmd.FLAG_FORMAT},
			Usage:   "show all stages",
			Action: func(*cli.Context) error {
				fmt.Println("NAME\tDESCRIPTION")
				for _, s := range cmd.STAGES {
					fmt.Println(s)
				}
				return nil
			},
		},
		cli.Command{
			Name:    "tasks",
			Aliases: []string{"T"},
			Flags:   []cli.Flag{cmd.FLAG_FORMAT},
			Usage:   "show all tasks",
			Action: func(*cli.Context) error {
				fmt.Println("NAME\tDESCRIPTION")
				for _, s := range cmd.STAGES {
					fmt.Println(s)
				}
				return nil
			},
		},
	)
}
