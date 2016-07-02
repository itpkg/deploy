package main

import (
	"log"

	"github.com/itpkg/deploy"
	_ "github.com/itpkg/deploy/cmd/base"
	_ "github.com/itpkg/deploy/cmd/run"
	_ "github.com/itpkg/deploy/scm/git"
)

func main() {
	if err := deploy.Run("rma"); err != nil {
		log.Fatal(err)
	}
}
