package main

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/op/go-logging"
	"github.com/urfave/cli"
)

//Action command action
func Action(fn func(*cli.Context, *logging.Logger) error) cli.ActionFunc {
	return func(c *cli.Context) error {

		lfd, err := os.OpenFile(
			time.Now().Format("2006-01-02.log"),
			os.O_WRONLY|os.O_APPEND|os.O_CREATE,
			0600,
		)
		if err != nil {
			return err
		}
		defer lfd.Close()

		level := logging.DEBUG

		bkd1 := logging.AddModuleLevel(
			logging.NewBackendFormatter(
				logging.NewLogBackend(os.Stderr, "", 0),
				logging.MustStringFormatter(`%{color}%{time:2006-01-02 15:04:05.000} ▶ %{level:.4s} %{color:reset} %{message}`)),
		)
		bkd1.SetLevel(level, "")

		bkd2 := logging.AddModuleLevel(
			logging.NewBackendFormatter(
				logging.NewLogBackend(lfd, "", 0),
				logging.MustStringFormatter(`%{time:2006-01-02 15:04:05.000} %{level:.4s} %{message}`)),
		)
		bkd2.SetLevel(level, "")
		logging.SetBackend(
			bkd1,
			bkd2,
		)

		l := logging.MustGetLogger(c.App.Name)
		l.Infof("=== BEGIN ===")
		l.Infof("stage: %s", c.Args()[0])
		l.Infof("task: %s", c.Args()[1])
		l.Infof("hosts: %q", c.StringSlice("hosts"))
		l.Infof("roles: %q", c.StringSlice("roles"))
		err = fn(c, l)
		l.Infof("=== END ===")
		return err
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "rma"
	app.Usage = "Remote multi-server automation tool"
	app.Version = "v20160702"
	app.Flags = []cli.Flag{
		cli.StringSliceFlag{
			Name:   "hosts",
			Usage:  "hosts filter, like: 'deploy@host1.com,deploy@host2.com'",
			EnvVar: "HOSTS",
		},
		cli.StringSliceFlag{
			Name:   "roles",
			Usage:  "roles filter, like: 'app,web'",
			EnvVar: "ROLES",
		},
	}
	app.Action = Action(func(c *cli.Context, l *logging.Logger) error {

		if len(c.Args()) != 2 {
			return errors.New("usage: rma stage_name task_name")
		}
		return nil
	})

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
