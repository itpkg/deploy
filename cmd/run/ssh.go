package run

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/itpkg/deploy/cmd"
	"golang.org/x/crypto/ssh"
)

//Exec run scripts on host
func Exec(stage *cmd.Stage, host string, scripts ...string) error {
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

	job = append(
		job,
		"cd {{.To}}",
		"ln -sfn {{.Version}} current",
		`find . -maxdepth 1 -name '20*' | sort -r | tail -n +{{.Keep}} | tr '\n' '\0' |  xargs -0 rm -r --`,
	)

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
	tpl, err := template.New("").Parse(strings.Join(job, "; "))
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if err = tpl.Execute(&buf, stage); err != nil {
		return err
	}

	ses, err := con.NewSession()
	if err != nil {
		return err
	}
	defer ses.Close()
	var out bytes.Buffer
	ses.Stderr = os.Stderr
	ses.Stdout = &out
	if err = ses.Run(buf.String()); err != nil {
		return err
	}
	stage.Logger.Debugf("%s: %s\n%s", host, buf.String(), out.String())

	return nil
}
