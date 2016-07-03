package scm

import (
	"fmt"
	"reflect"
	"strings"
)

//Scm the Source Control Management used
type Scm interface {
	Clone() []string
}

var plugins = make(map[string]Scm)

//Register register scm
func Register(scms ...Scm) {
	for _, s := range scms {
		n := strings.ToLower(reflect.Indirect(reflect.ValueOf(s)).Type().Name())
		plugins[n] = s
	}
}

//Get get scm
func Get(n string) (Scm, error) {
	s, o := plugins[n]
	if o {
		return s, nil
	}
	return nil, fmt.Errorf("bad name %s", n)
}
