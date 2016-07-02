package deploy

import "github.com/op/go-logging"

//Stage load from config/stages/<name>.toml
type Stage struct {
	Name string `toml:"-"`
	//The path on the remote server where the application should be deployed.
	//default "/var/www/{{.Name}}"
	To string `toml:"deploy_to"`
	//The Source Control Management used.
	//default: :git
	//Currently :git are supported.
	Scm string `toml:"scm"` //default git
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
	Tmp string `toml:"tmp"`
	//Log level
	//default debug
	Level string `toml:"log_level"`

	Logger *logging.Logger `toml:"-"`
}
