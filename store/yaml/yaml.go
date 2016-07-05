package yaml

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/itpkg/deploy/store"
	_yaml "gopkg.in/yaml.v2"
)

//Yaml store by yaml
type Yaml struct {
}

//Ext config file's ext
func (p *Yaml) Ext() string {
	return ".toml"
}

//Write write object to toml file
func (p *Yaml) Write(n string, v interface{}) error {

	if err := os.MkdirAll(path.Dir(n), 0700); err != nil {
		return err
	}
	buf, err := _yaml.Marshal(v)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(n, buf, 0600)
}

//Read read object from file
func (p *Yaml) Read(n string, v interface{}) error {
	buf, err := ioutil.ReadFile(n)
	if err != nil {
		return err
	}
	return _yaml.Unmarshal(buf, v)
}

func init() {
	store.Register(&Yaml{})
}
