package git

import "github.com/itpkg/deploy/scm"

//Git git
type Git struct {
}

//Clone clone
func (p *Git) Clone() []string {
	return []string{
		`
		if [ ! -d '{{.To}}/repo' ];
		then
		git clone {{.Repo}} {{.To}}/repo;
		{{ if ne .Branch "master"}}
		cd {{.To}}/repo;
		git checkout -b {{.Branch}};
		git branch --set-upstream-to=origin/{{.Branch}} {{.Branch}};
		{{ end }}

		fi`,
		"cd {{.To}}/repo && git pull",
		"git clone --depth 1 --branch {{.Branch}} file://{{.To}}/repo {{.To}}/{{.Version}}",
	}
}

func init() {
	scm.Register(&Git{})
}
