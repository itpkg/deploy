package deploy

import (
	"os"

	"github.com/urfave/cli"
)

//Run main entry
func Run(name string) error {
	app := cli.NewApp()
	app.Name = name
	app.Usage = "Remote multi-server automation tool"
	app.Version = "v20160702"
	app.Flags = []cli.Flag{}

	app.Commands = commands

	return app.Run(os.Args)
}
