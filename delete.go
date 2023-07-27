package file

import "os"

func deleteFileFromHdd(file_path string) error {

	// Borrar archivos
	err := os.Remove(file_path)
	if err != nil {
		return err
	}

	return nil
}
