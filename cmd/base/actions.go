package base

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/itpkg/deploy"
	"github.com/itpkg/deploy/cmd"
	"github.com/urfave/cli"
)

func list(p string) cli.ActionFunc {
	return func(*cli.Context) error {
		return filepath.Walk(p, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			name := info.Name()
			if !info.IsDir() && filepath.Ext(name) == cmd.EXT {
				fmt.Println(name[:len(name)-len(cmd.EXT)])
			}
			return nil
		})
	}
}

func generate(c *cli.Context) error {
	sn := c.String("stage")
	tn := c.String("task")

	if sn != "" {
		fn := cmd.Stage(sn)
		fmt.Printf("generate file %s\n", fn)
		if err := Write(
			fn,
			&deploy.Stage{
				Name: sn,
			}); err != nil {
			return err
		}
	}

	if tn != "" {
		fn := cmd.Task(tn)
		fmt.Printf("generate file %s\n", fn)
		if err := Write(
			fn,
			&deploy.Task{
				Name:   tn,
				Roles:  []string{"web", "app"},
				Hosts:  []string{"deploy@host1.com", "deploy@host2.com"},
				Script: []string{"uname -a", "whoami"},
			},
		); err != nil {
			return err
		}
	}

	if sn == "" && tn == "" {
		cli.ShowCommandHelp(c, "generate")
	}
	return nil
}
