package deploy

import "fmt"

//Task load from config/tasks/name>.toml
type Task struct {
	File        string   `toml:"-" yaml:"-"`
	Name        string   `toml:"name" yaml:"name"`
	Description string   `toml:"description" yaml:"description"`
	Refresh     bool     `toml:"refresh" yaml:"refresh"`
	Hosts       []string `toml:"hosts" yaml:"hosts"`
	Roles       []string `toml:"roles" yaml:"roles"`
	Script      []string `toml:"script" yaml:"script"`
}

func (p *Task) String() string {
	return fmt.Sprintf("%s\t%s", p.Name, p.Description)
}

//-----------------------------------------------------------------------------
var TASKS = make(map[string]*Task)
