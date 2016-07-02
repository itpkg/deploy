package run

import (
	"bytes"
	"fmt"
	"html/template"
	"path"
	"sort"
	"strings"

	"github.com/itpkg/deploy/cmd"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh"
)

//Exec run scripts on host
func Exec(stage *cmd.Stage, host string, scripts ...string) error {
	ss := strings.Split(host, "@")
	if len(ss) != 2 {
		return fmt.Errorf("bad host: %s", host)
	}
	if strings.Index(ss[1], ":") == -1 {
		ss[1] += ":22"
	}
	con, err := ssh.Dial("tcp", ss[1], &ssh.ClientConfig{
		User: ss[0],
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(stage.Signers...),
			//ssh.Password("yourpassword"),
		},
	})
	if err != nil {
		return err
	}

	for _, s := range scripts {
		//parse template
		tpl, err := template.New("").Parse(s)
		if err != nil {
			return err
		}
		var buf bytes.Buffer
		if err = tpl.Execute(&buf, stage); err != nil {
			return err
		}

		ses, err := con.NewSession()
		if err != nil {
			return err
		}
		defer ses.Close()
		var out bytes.Buffer
		//ses.Stdout = os.Stdout
		ses.Stdout = &out
		if err = ses.Run(buf.String()); err != nil {
			return err
		}
		stage.Logger.Debugf("%s: %s\n%s", host, buf.String(), out.String())
	}
	return nil
}

func run(c *cli.Context, s *cmd.Stage) error {
	var task cmd.Task
	tn := c.String("task")
	if err := s.Store.Read(
		path.Join(cmd.TASKS, fmt.Sprintf("%s%s", tn, s.Store.Ext())),
		&task); err != nil {
		return err
	}
	task.Name = tn

	//hosts
	hosts := c.StringSlice("hosts")
	roles := c.StringSlice("roles")
	s.Logger.Infof("task: %s@%s", task.Name, s.Name)
	s.Logger.Infof("roles: %q", roles)
	s.Logger.Infof("hosts: %q", hosts)
	if len(hosts) == 0 {
		hosts = task.Hosts
	}
	if len(roles) == 0 {
		roles = task.Roles
	}

	all := false
	for _, r := range roles {
		if r == "all" {
			all = true
			break
		}
	}

	if all {
		for _, hs := range s.Roles {
			hosts = append(hosts, hs...)
		}
	} else {
		for _, r := range roles {
			if hs, ok := s.Roles[r]; ok {
				hosts = append(hosts, hs...)
			} else {
				return fmt.Errorf("could not find role %s", r)
			}
		}
	}
	target := make(map[string]bool)
	for _, h := range hosts {
		target[h] = true
	}
	hosts = make([]string, 0)
	for h, _ := range target {
		hosts = append(hosts, h)
	}
	sort.Strings(hosts)

	s.Logger.Debugf("ordered hosts: %q", hosts)
	for _, h := range hosts {
		if err := Exec(s, h, task.Script...); err != nil {
			return err
		}
	}

	return nil
}
