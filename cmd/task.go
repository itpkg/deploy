package cmd

//Task load from config/tasks/name>.toml
type Task struct {
	Name        string   `toml:"name" yaml:"name"`
	Description string   `toml:"description" yaml:"description"`
	Refresh     bool     `toml:"refresh" yaml:"refresh"`
	Hosts       []string `toml:"hosts" yaml:"hosts"`
	Roles       []string `toml:"roles" yaml:"roles"`
	Script      []string `toml:"script" yaml:"script"`
}
