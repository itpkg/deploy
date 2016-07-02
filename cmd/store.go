package cmd

import (
	"os"
	"path"

	"github.com/BurntSushi/toml"
)

//Read read object from file
func Read(n string, v interface{}) error {
	_, err := toml.DecodeFile(n, v)
	return err
}

//Write write object to toml file
func Write(n string, v interface{}) error {
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