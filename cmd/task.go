package cmd

//Task load from config/tasks/name>.toml
type Task struct {
	Name        string   `toml:"-"`
	Description string   `toml:"description"`
	Hosts       []string `toml:"hosts"`
	Roles       []string `toml:"roles"`
	Script      []string `toml:"script"`
	Deploy      bool     `toml:"deploy"`
}
