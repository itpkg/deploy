package cmd

import (
	"fmt"
	"path"
	"sort"

	"github.com/itpkg/deploy/scm"
	"github.com/itpkg/deploy/store"
	"github.com/op/go-logging"
	"golang.org/x/crypto/ssh"
)

//Stage load from config/stages/<name>.toml
type Stage struct {
	Name        string `toml:"-"`
	Description string `toml:"description"`
	//The path on the remote server where the application should be deployed.
	//default "/var/www/{{.Name}}"
	To string `toml:"deploy_to"`
	//The Source Control Management used.
	//default: :git
	//Currently :git are supported.
	ScmF string `toml:"scm"` //default git
	//URL to the repository.
	//Must be a valid URL for the used SCM.
	Repo string `toml:"repo_url"`
	//default master
	Branch string `toml:"branch"`
	//Listed files will be symlinked into each release directory during deployment.
	//default []
	Files []string `toml:"linked_files"`
	//Listed directories will be symlinked into the release directory during deployment.
	//default []
	Dirs []string `toml:"linked_dirs"`
	//Default shell environment used during command execution.
	//default {}
	Env map[string]string `toml:"default_env"`
	//The last n releases are kept for possible rollbacks.
	//default 5
	Keep uint `toml:"keep_releases"`
	//Temporary directory used during deployments to store data.
	//default /tmp
	Tmp   string `toml:"tmp"`
	Debug bool   `toml:"debug"`

	//split hosts by roles
	Roles map[string][]string `toml:"roles"`
	Keys  []string            `toml:"keys"`

	Logger  *logging.Logger `toml:"-"`
	Store   store.Store     `toml:"-"`
	Scm     scm.Scm         `toml:"-"`
	Signers []ssh.Signer    `toml:"-"`

	Version string `toml:"-"`
}

//Shared shared path
func (p *Stage) Shared(n string) string {
	return path.Join(p.To, "shared", n)
}

//Hosts get ordered hosts
func (p *Stage) Hosts(task *Task, roles, hosts []string) ([]string, error) {
	if len(roles) == 0 {
		roles = task.Roles
	}
	if len(hosts) == 0 {
		hosts = task.Hosts
	}

	all := false
	for _, r := range roles {
		if r == "all" {
			all = true
			break
		}
	}

	if all {
		for _, hs := range p.Roles {
			hosts = append(hosts, hs...)
		}
	} else {
		for _, r := range roles {
			if hs, ok := p.Roles[r]; ok {
				hosts = append(hosts, hs...)
			} else {
				return nil, fmt.Errorf("could not find role %s", r)
			}
		}
	}
	target := make(map[string]bool)
	for _, h := range hosts {
		target[h] = true
	}

	ret := make([]string, 0)
	for h, _ := range target {
		ret = append(ret, h)
	}
	sort.Strings(ret)
	return ret, nil
}
