package deploy

import (
	"fmt"
	"path"

	"github.com/itpkg/deploy/store"
)

//Task load from config/tasks/name>.toml
type Task struct {
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

func loadTasks() {

	if err := walk(path.Join("config", "tasks"), func(n string, s store.Store) error {
		var tk Task
		if err := s.Read(n, &tk); err != nil {
			return err
		}
		TASKS[tk.Name] = &tk
		return nil
	}); err != nil {
		panic(err)
	}
}
