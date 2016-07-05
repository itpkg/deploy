package cmd

//Task load from config/tasks/name>.toml
type Task struct {
	Name        string   `toml:"-"`
	Description string   `toml:"description"`
	Refresh     bool     `toml:"refresh"`
	Hosts       []string `toml:"hosts"`
	Roles       []string `toml:"roles"`
	Script      []string `toml:"script"`
}
