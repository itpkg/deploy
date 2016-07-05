package store

import "fmt"

//Store store
type Store interface {
	Ext() string
	Write(n string, v interface{}) error
	Read(n string, v interface{}) error
}

var plugins []Store

//Register register storages
func Register(ss ...Store) {
	// for _, s := range ss {
	// 	// n := strings.ToLower(reflect.Indirect(reflect.ValueOf(s)).Type().Name())
	// 	// plugins[n] = s
	// }
	plugins = append(plugins, ss...)
}

//Get get store
func Get(ext string) (Store, error) {
	for _, s := range plugins {
		if s.Ext() == ext {
			return s, nil
		}
	}

	return nil, fmt.Errorf("bad format name %s", ext)
}
