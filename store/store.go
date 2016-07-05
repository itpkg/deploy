package store

import (
	"fmt"
	"reflect"
	"strings"
)

//Store store
type Store interface {
	Ext() string
	Write(n string, v interface{}) error
	Read(n string, v interface{}) error
}

var plugins = make(map[string]Store)

//Register register storages
func Register(ss ...Store) {
	for _, s := range ss {
		n := strings.ToLower(reflect.Indirect(reflect.ValueOf(s)).Type().Name())
		plugins[n] = s
	}

}

//Get get store
func Get(n string) (Store, error) {
	s, o := plugins[n]
	if o {
		return s, nil
	}
	return nil, fmt.Errorf("bad format name %s", n)
}
