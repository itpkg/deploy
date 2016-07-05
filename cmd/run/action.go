package run

import (
	"fmt"

	"github.com/itpkg/deploy"
	"github.com/urfave/cli"
)

func run(c *cli.Context, s *deploy.Stage) error {
	tn := c.String("task")
	if len(tn) == 0 {
		cli.ShowCommandHelp(c, "run")
		return nil
	}
	tk, ok := deploy.TASKS[tn]
	if !ok {
		return fmt.Errorf("cann't find task %s", tk)
	}

	s.Logger.Infof("task: %s@%s", tk.Name, s.Name)
	// s.Logger.Infof("roles: %q", c.StringSlice("roles"))
	// s.Logger.Infof("hosts: %q", c.StringSlice("hosts"))
	hosts, err := s.Hosts(
		tk,
		c.StringSlice("roles"),
		c.StringSlice("hosts"),
	)
	if err != nil {
		return err
	}
	s.Logger.Debugf("hosts: %q", hosts)

	for _, h := range hosts {
		if err := Exec(s, h, tk.Refresh, tk.Script...); err != nil {
			return err
		}
	}

	return nil
}
