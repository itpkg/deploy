package deploy

import "github.com/urfave/cli"

var FLAG_HOSTS = cli.StringSliceFlag{
	Name:   "hosts",
	Usage:  "hosts filter, like: 'deploy@host1.com,deploy@host2.com'",
	EnvVar: "HOSTS",
}
var FLAG_ROLES = cli.StringSliceFlag{
	Name:   "roles",
	Usage:  "roles filter, like: 'app,web'",
	EnvVar: "ROLES",
}

var FLAG_FORMAT = cli.StringFlag{
	Name:   "format,f",
	Value:  "toml",
	Usage:  "file format, like: toml, yaml",
	EnvVar: "FORMAT",
}

var FLAG_STAGE = cli.StringFlag{
	Name:   "stage,s",
	Value:  "",
	Usage:  "stage name like: production, development, test",
	EnvVar: "STAGE",
}

var FLAG_TASK = cli.StringFlag{
	Name:   "task,t",
	Value:  "",
	Usage:  "task's name",
	EnvVar: "TASK",
}
