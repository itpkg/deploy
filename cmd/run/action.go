package run

import (
	"fmt"
	"path"

	"github.com/itpkg/deploy/cmd"
	"github.com/urfave/cli"
)

//Exec run scripts on host
func Exec(host string, scripts ...string) error {
	return nil
}

func run(c *cli.Context, s *cmd.Stage) error {
	var task cmd.Task
	tn := c.String("task")
	if err := cmd.Read(
		path.Join(cmd.TASKS, fmt.Sprintf("%s%s", tn, cmd.EXT)),
		&task); err != nil {
		return err
	}
	task.Name = tn

	hosts := c.StringSlice("hosts")
	roles := c.StringSlice("roles")
	if len(hosts) == 0 {
		hosts = task.Hosts
	}
	if len(roles) == 0 {
		roles = task.Roles
	}

	s.Logger.Infof("task: %s", task.Name)
	s.Logger.Infof("roles: %q", roles)
	s.Logger.Infof("hosts: %q", hosts)

	return nil
}
