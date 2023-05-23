package file

import (
	"os"
	"path/filepath"
)

func DeleteIfFolderSizeExceeds(folderPath string, maxSizeMB int64) error {
	var size int64

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})

	if err != nil {
		return err
	}

	sizeMB := size / 1024 / 1024
	if sizeMB > maxSizeMB {
		err := os.RemoveAll(folderPath)
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateFolderIfNotExist(folderPath string) error {
	_, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(folderPath, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
