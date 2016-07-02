package store

import (
	"os"
	"path"

	"github.com/BurntSushi/toml"
)

//TomlStore store by toml
type TomlStore struct {
}

//Ext config file's ext
func (p *TomlStore) Ext() string {
	return ".toml"
}

//Write write object to toml file
func (p *TomlStore) Write(n string, v interface{}) error {

	if err := os.MkdirAll(path.Dir(n), 0700); err != nil {
		return err
	}
	fd, err := os.OpenFile(n, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		return err
	}
	end := toml.NewEncoder(fd)
	return end.Encode(v)
}

//Read read object from file
func (p *TomlStore) Read(n string, v interface{}) error {
	_, err := toml.DecodeFile(n, v)
	return err

}
