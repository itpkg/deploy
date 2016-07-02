package store

//Store store
type Store interface {
	Ext() string
	Write(n string, v interface{}) error
	Read(n string, v interface{}) error
}

//New new store by name
func New(f string) Store {
	switch f {
	default:
		return &TomlStore{}
	}
}
