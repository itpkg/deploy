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
