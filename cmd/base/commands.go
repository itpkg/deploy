package base

import (
	"fmt"

	"github.com/itpkg/deploy"
	"github.com/itpkg/deploy/cmd"
	"github.com/itpkg/deploy/store"
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
				cmd.FLAG_FORMAT,
			},
			Action: generate,
		},
		cli.Command{
			Name:    "stages",
			Aliases: []string{"S"},
			Flags:   []cli.Flag{cmd.FLAG_FORMAT},
			Usage:   "show all stages",
			Action: list(cmd.STAGES, func(c *cli.Context, p, n string) error {
				var st cmd.Stage
				so, err := store.Get(c.String("format"))
				if err != nil {
					return err
				}
				if err = so.Read(p, &st); err != nil {
					return err
				}
				fmt.Printf("%s\t%s\n", n, st.Description)
				return nil
			}),
		},
		cli.Command{
			Name:    "tasks",
			Aliases: []string{"T"},
			Flags:   []cli.Flag{cmd.FLAG_FORMAT},
			Usage:   "show all tasks",
			Action: list(cmd.TASKS, func(c *cli.Context, p, n string) error {
				var tk cmd.Task
				so, err := store.Get(c.String("format"))
				if err != nil {
					return err
				}
				if err = so.Read(p, &tk); err != nil {
					return err
				}
				fmt.Printf("%s\t%s\n", n, tk.Description)
				return nil
			}),
		},
	)
}
