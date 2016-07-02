package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"

	"golang.org/x/crypto/ssh"

	"github.com/itpkg/deploy/store"
	"github.com/op/go-logging"
	"github.com/urfave/cli"
)

//Action command action
func Action(fn func(*cli.Context, *Stage) error) cli.ActionFunc {
	return func(c *cli.Context) error {
		var st Stage
		st.Store = store.New(c.String("format"))
		if err := st.Store.Read(
			path.Join(STAGES, fmt.Sprintf("%s%s", c.String("stage"), st.Store.Ext())),
			&st); err != nil {
			return err
		}
		st.Name = c.String("stage")

		lfd, err := os.OpenFile(
			time.Now().Format("2006-01-02.log"),
			os.O_WRONLY|os.O_APPEND|os.O_CREATE,
			0600,
		)
		if err != nil {
			return err
		}
		defer lfd.Close()

		bkd1 := logging.AddModuleLevel(
			logging.NewBackendFormatter(
				logging.NewLogBackend(os.Stderr, "", 0),
				logging.MustStringFormatter(`%{color}%{time:2006-01-02 15:04:05.000} ▶ %{level:.4s} %{color:reset} %{message}`)),
		)

		bkd2 := logging.AddModuleLevel(
			logging.NewBackendFormatter(
				logging.NewLogBackend(lfd, "", 0),
				logging.MustStringFormatter(`%{time:15:04:05.000} %{level:.4s} %{message}`)),
		)

		logging.SetBackend(
			bkd1,
			bkd2,
		)

		if !st.Debug {
			bkd1.SetLevel(logging.INFO, "")
			bkd2.SetLevel(logging.INFO, "")
		}

		l := logging.MustGetLogger(c.App.Name)
		l.Infof("=== BEGIN ===")
		st.Logger = l

		//load ssh keys
		for _, key := range st.Keys {
			buf, er := ioutil.ReadFile(key)
			if er != nil {
				return er
			}
			sig, er := ssh.ParsePrivateKey(buf)
			if er != nil {
				return er
			}
			st.Signers = append(st.Signers, sig)
		}

		err = fn(c, &st)
		l.Infof("=== END ===")
		return err
	}
}
