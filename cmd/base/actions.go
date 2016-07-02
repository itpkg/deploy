package base

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"

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
		fn := path.Join(cmd.STAGES, fmt.Sprintf("%s%s", sn, cmd.EXT))
		fmt.Printf("generate file %s\n", fn)
		if err := cmd.Write(
			fn,
			&cmd.Stage{
				Name:   sn,
				To:     fmt.Sprintf("/var/www/%s", sn),
				Scm:    "git",
				Repo:   "http://github.com/change-me.git",
				Branch: "master",
				Files:  []string{"config.toml"},
				Dirs:   []string{"logs", "public", "tmp"},
				Env: map[string]string{
					"created": time.Now().String(),
				},
				Keep: 5,
				Tmp:  os.TempDir(),
				Roles: map[string][]string{
					"web": []string{"www.host1.com", "www.host2.com"},
					"app": []string{"app.host1.com", "app.host2.com"},
				},
				Debug: true,
			}); err != nil {
			return err
		}
	}

	if tn != "" {
		fn := path.Join(cmd.TASKS, fmt.Sprintf("%s%s", tn, cmd.EXT))
		fmt.Printf("generate file %s\n", fn)
		if err := cmd.Write(
			fn,
			&cmd.Task{
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
