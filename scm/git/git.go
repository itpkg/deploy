package git

import (
	"fmt"

	"github.com/itpkg/rma/scm"
)

//Git git
type Git struct {
}

//Clone clone
func (p *Git) Clone(repo, dir string) string {
	return fmt.Sprintf("git clone %s %s", repo, dir)
}

//Fetch fetch
func (p *Git) Fetch() string {
	return "git fetch"
}

//Checkout checkout
func (p *Git) Checkout(branch string) string {
	return fmt.Sprintf("git checkout -b %s", branch)
}

func init() {
	scm.Register(&Git{})
}
