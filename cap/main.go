package main

import (
	"log"

	"github.com/itpkg/deploy"
	_ "github.com/itpkg/deploy/cmd/base"
	_ "github.com/itpkg/deploy/cmd/run"
	_ "github.com/itpkg/deploy/scm/git"
	_ "github.com/itpkg/deploy/store/toml"
	_ "github.com/itpkg/deploy/store/yaml"
)

func main() {
	err := deploy.Load()
	if err == nil {
		err = deploy.Run("cap")
	}
	if err != nil {
		log.Fatal(err)
	}
}
