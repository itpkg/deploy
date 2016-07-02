package scm

import (
	"reflect"
	"strings"
)

//Scm the Source Control Management used
type Scm interface {
	Clone(repo, dir string) string
	Fetch() string
	Checkout(branch string) string
}

var plugins = make(map[string]Scm)

//Register register scm
func Register(s Scm) {
	plugins[strings.ToLower(reflect.TypeOf(s).Name())] = s
}

//Get get scm
func Get(n string) Scm {
	return plugins[n]
}
