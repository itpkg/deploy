package cmd

import (
	"fmt"
	"path"
)

//ROOT config dir
const ROOT = "config"

//EXT config file's ext
const EXT = ".toml"

//STAGES stages dir
var STAGES = path.Join(ROOT, "stages")

//TASKS tasks dir
var TASKS = path.Join(ROOT, "tasks")

//Stage stage config file name
func Stage(n string) string {
	return path.Join(STAGES, fmt.Sprintf("%s%s", n, EXT))
}

//Task task filename
func Task(n string) string {
	return path.Join(TASKS, fmt.Sprintf("%s%s", n, EXT))
}
