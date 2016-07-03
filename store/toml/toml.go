package toml

import (
	"os"
	"path"

	_toml "github.com/BurntSushi/toml"
	"github.com/itpkg/deploy/store"
)

//Toml store by toml
type Toml struct {
}

//Ext config file's ext
func (p *Toml) Ext() string {
	return ".toml"
}

//Write write object to toml file
func (p *Toml) Write(n string, v interface{}) error {

	if err := os.MkdirAll(path.Dir(n), 0700); err != nil {
		return err
	}
	fd, err := os.OpenFile(n, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		return err
	}
	end := _toml.NewEncoder(fd)
	return end.Encode(v)
}

//Read read object from file
func (p *Toml) Read(n string, v interface{}) error {
	_, err := _toml.DecodeFile(n, v)
	return err
}

func init() {
	store.Register(&Toml{})
}
