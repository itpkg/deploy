package cmd

import "path"

//ROOT config dir
const ROOT = "config"

//EXT config file's ext
const EXT = ".toml"

//STAGES stages dir
var STAGES = path.Join(ROOT, "stages")

//TASKS tasks dir
var TASKS = path.Join(ROOT, "tasks")
