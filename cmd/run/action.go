package run

import (
	"fmt"
	"path"

	"github.com/itpkg/deploy/cmd"
	"github.com/urfave/cli"
)

func run(c *cli.Context, s *cmd.Stage) error {
	var task cmd.Task
	tn := c.String("task")
	if err := s.Store.Read(
		path.Join(cmd.TASKS, fmt.Sprintf("%s%s", tn, s.Store.Ext())),
		&task); err != nil {
		return err
	}
	task.Name = tn
	s.Logger.Infof("task: %s@%s", task.Name, s.Name)
	// s.Logger.Infof("roles: %q", c.StringSlice("roles"))
	// s.Logger.Infof("hosts: %q", c.StringSlice("hosts"))
	hosts, err := s.Hosts(
		&task,
		c.StringSlice("roles"),
		c.StringSlice("hosts"),
	)
	if err != nil {
		return err
	}
	s.Logger.Debugf("hosts: %q", hosts)

	for _, h := range hosts {
		if err := Exec(s, h, task.Refresh, task.Script...); err != nil {
			return err
		}
	}

	return nil
}
