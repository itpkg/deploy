package deploy

//Task load from config/tasks/name>.toml
type Task struct {
	Name   string   `toml:"-"`
	Hosts  []string `toml:"hosts"`
	Roles  []string `toml:"roles"`
	Script []string `toml:"script"`
}
