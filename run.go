package deploy

import (
	"fmt"
	"os"
	"path"

	"github.com/itpkg/deploy/store"
	"github.com/urfave/cli"
)

func Load() error {
	if err := walk(path.Join("config", "stages"), func(n string, s store.Store) error {
		var st Stage
		if err := s.Read(n, &st); err != nil {
			return err
		}
		if tm, ok := STAGES[st.Name]; ok {
			return fmt.Errorf("file %s and %s have same name", tm.File, n)
		} else {
			st.File = n
			STAGES[st.Name] = &st
		}

		return nil
	}); err != nil {
		return err
	}
	if err := walk(path.Join("config", "tasks"), func(n string, s store.Store) error {
		var tk Task
		if err := s.Read(n, &tk); err != nil {
			return err
		}
		if tm, ok := TASKS[tk.Name]; ok {
			return fmt.Errorf("file %s and %s have same name", tm.File, n)
		} else {
			tk.File = n
			TASKS[tk.Name] = &tk
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

//Run main entry
func Run(name string) error {
	app := cli.NewApp()
	app.Name = name
	app.Usage = "Remote multi-server automation tool"
	app.Version = "v20160705"
	app.Flags = []cli.Flag{}

	app.Commands = commands

	return app.Run(os.Args)
}
