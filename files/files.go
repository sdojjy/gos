package files

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func AppHome() (appHome string, err error) {
	dir, absErr := filepath.Abs(filepath.Dir(os.Args[0]))
	if nil != absErr {
		err = absErr
	}
	appHome = strings.Replace(dir, "\\", "/", -1)

	return
}

func WriteBytes(filename string, data []byte, perm os.FileMode) error {
	if !CheckExist(filename) {
		parent := filepath.Dir(filename)
		os.MkdirAll(parent, perm)
	}

	return ioutil.WriteFile(filename, data, perm)
}

func CheckExist(filename string) bool {
	var exist = true

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}

	return exist
}

func CheckAllExist(filenames ...string) bool {
	var exist = true

	for _, filename := range filenames {
		if !CheckExist(filename) {
			exist = false
			break
		}
	}

	return exist
}
