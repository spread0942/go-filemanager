package filemanager

import (
	"os"
	"path/filepath"
)

func CreateFile(filename string, path string) (bool, error) {
	fullpath := filepath.Join(path, filename)
	file, err := os.Create(fullpath)

	if (err != nil) {
		return false, err
	}

	defer file.Close()
	return true, nil
}

func CreateDirectory(directory string, path string) (bool, error) {
	fullpath := filepath.Join(path, directory)
	err := os.Mkdir(fullpath, 0755)

	if (err != nil) {
		return false, err
	}
	return true, nil
}