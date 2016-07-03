package run

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"github.com/itpkg/deploy/cmd"
	"golang.org/x/crypto/ssh"
)

//Exec run scripts on host
func Exec(stage *cmd.Stage, host string, deploy bool, scripts ...string) error {

	job := stage.Scm.Clone()
	job = append(job,
		"mkdir -pv {{.To}}/shared",
		"cd {{.To}}",
	)
	for _, s := range stage.Dirs {
		s = stage.Shared(s)
		job = append(job, fmt.Sprintf("test -d %s || mkdir -pv %s", s, s))
	}
	for _, f := range stage.Files {
		f = stage.Shared(f)
		job = append(job, fmt.Sprintf("test -f %s || touch %s", f, f))
	}
	job = append(job, "ln -s {{.To}}/shared/* {{.To}}/{{.Version}}")
	job = append(job, scripts...)

	if deploy {
		job = append(
			job,
			"cd {{.To}} && ln -sfn {{.Version}} current",
			fmt.Sprintf(
				//`cd {{.To}} && $(find . -maxdepth 1 -name '20*' | wc -l) -lt %d || find . -maxdepth 1 -name '20*' | sort -r | tail -n +%d | tr '\n' '\0' |  xargs -0 rm -r --`,
				`cd {{.To}} && test $(find . -maxdepth 1 -name '20*' | wc -l) -lt %d || find . -maxdepth 1 -name '20*' | sort -r | tail -n +%d | tr '\n' '\0' |  xargs -0 rm -r --`,
				stage.Keep,
				stage.Keep+1,
			),
		)
	}

	//run
	ss := strings.Split(host, "@")
	if len(ss) != 2 {
		return fmt.Errorf("bad host: %s", host)
	}
	if strings.Index(ss[1], ":") == -1 {
		ss[1] += ":22"
	}
	con, err := ssh.Dial("tcp", ss[1], &ssh.ClientConfig{
		User: ss[0],
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(stage.Signers...),
			//ssh.Password("yourpassword"),
		},
	})
	if err != nil {
		return err
	}

	//parse template
	for _, s := range job {
		//tpl, err := template.New("").Parse(strings.Join(job, "; "))
		tpl, err := template.New("").Parse(s)
		if err != nil {
			return err
		}
		var in bytes.Buffer
		if err = tpl.Execute(&in, stage); err != nil {
			return err
		}

		stage.Logger.Debugf("%s: %s", host, in.String())
		ses, err := con.NewSession()
		if err != nil {
			return err
		}
		defer ses.Close()
		var out bytes.Buffer
		ses.Stderr = &out
		ses.Stdout = &out
		if err = ses.Run(in.String()); err != nil {
			return err
		}
		stage.Logger.Debug(out.String())
	}
	return nil
}
