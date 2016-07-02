package base

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/itpkg/deploy/cmd"
	"github.com/itpkg/deploy/store"
	"github.com/urfave/cli"
)

func list(p string) cli.ActionFunc {
	return func(c *cli.Context) error {
		return filepath.Walk(p, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			name := info.Name()
			st := store.New(c.String("format"))
			if !info.IsDir() && filepath.Ext(name) == st.Ext() {
				fmt.Println(name[:len(name)-len(st.Ext())])
			}
			return nil
		})
	}
}

func generate(c *cli.Context) error {
	sn := c.String("stage")
	tn := c.String("task")
	st := store.New(c.String("format"))

	if sn != "" {
		fn := path.Join(cmd.STAGES, fmt.Sprintf("%s%s", sn, st.Ext()))
		fmt.Printf("generate file %s\n", fn)
		if err := st.Write(
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
					"web": []string{"deploy@www.host1.com", "deploy@www.host2.com"},
					"app": []string{"deploy@app.host1.com", "deploy@app.host2.com"},
					"db":  []string{"deploy@db.host1.com", "deploy@db.host2.com"},
				},
				Debug: true,
			}); err != nil {
			return err
		}
	}

	if tn != "" {
		fn := path.Join(cmd.TASKS, fmt.Sprintf("%s%s", tn, st.Ext()))
		fmt.Printf("generate file %s\n", fn)
		if err := st.Write(
			fn,
			&cmd.Task{
				Name:   tn,
				Roles:  []string{"web", "app"},
				Hosts:  []string{"deploy@host1.com", "deploy@host2.com"},
				Script: []string{"uname -a", "whoami", "echo {{.Name}}"},
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
