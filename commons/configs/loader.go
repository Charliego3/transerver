package configs

import (
	"errors"
	"github.com/gookit/goutil/strutil"
	"os"
)

type Loader interface {
	Load() ([]byte, error)
}

// FileLoader load content from os file
type FileLoader string

func (fl FileLoader) Load() ([]byte, error) {
	if strutil.IsBlank(string(fl)) {
		return nil, errors.New("can't read from empty path")
	}

	return os.ReadFile(string(fl))
}

func NewFileLoader(path string) Loader {
	return FileLoader(path)
}
