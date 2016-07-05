package base

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"time"

	"github.com/itpkg/deploy"
	"github.com/itpkg/deploy/store"
	"github.com/urfave/cli"
)

func generate(c *cli.Context) error {
	sn := c.String("stage")
	tn := c.String("task")
	st, err := store.Get("." + c.String("format"))
	if err != nil {
		return err
	}
	user, err := user.Current()
	if err != nil {
		return err
	}

	if sn != "" {
		fn := path.Join("config", "stages", fmt.Sprintf("%s%s", sn, st.Ext()))
		fmt.Printf("generate file %s\n", fn)
		if err := st.Write(
			fn,
			&deploy.Stage{
				Name:        sn,
				Description: "change me",
				To:          fmt.Sprintf("/var/www/%s", sn),
				ScmF:        "git",
				Repo:        "http://github.com/change-me.git",
				Branch:      "master",
				Files:       []string{"config.toml"},
				Dirs:        []string{"logs", "public", "tmp"},
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
				Keys: []string{
					path.Join(user.HomeDir, ".ssh", "id_rsa"),
				},
				Debug: true,
			}); err != nil {
			return err
		}
	}

	if tn != "" {
		fn := path.Join("config", "tasks", fmt.Sprintf("%s%s", tn, st.Ext()))
		fmt.Printf("generate file %s\n", fn)
		if err := st.Write(
			fn,
			&deploy.Task{
				Name:        tn,
				Description: "change me",
				Roles:       []string{"web", "app"},
				Hosts:       []string{"deploy@host1.com", "deploy@host2.com"},
				Script:      []string{"uname -a", "whoami", "echo {{.Name}}"},
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
