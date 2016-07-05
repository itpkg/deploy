package deploy

import (
	"os"
	"path/filepath"

	"github.com/itpkg/deploy/store"
)

func walk(dir string, fn func(string, store.Store) error) error {
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			ste, err := store.Get(filepath.Ext(info.Name()))
			if err != nil {
				return err
			}
			return fn(path, ste)
		}
		return nil
	})
}
