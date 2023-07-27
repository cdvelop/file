package file

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func (c config) processFiles(files []*multipart.FileHeader, upload_folder string, new_data map[string]string) ([]map[string]string, error) {
	data_out := []map[string]string{}
	for _, fileHeader := range files {
		if fileHeader.Size > c.maximum_file_size {
			return nil, fmt.Errorf(fmt.Sprintf("error archivo(s) excede(n) el tamaño admitido de: %v kb", c.max_kb_size), http.StatusNotAcceptable)
		}

		file, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}
		defer file.Close()

		extension := c.getExtension(fileHeader)

		if !strings.Contains(c.extensions, extension) {
			return nil, fmt.Errorf("formato de archivo no valido solo se admiten: %v", c.extensions)
		}
		extension = filepath.Ext(fileHeader.Filename)

		new_file_name := getNewID()
		new_data["id_file"] = new_file_name
		new_data["file_path"] = upload_folder + "/" + new_file_name + extension

		if len(fileHeader.Filename) > 5 {
			new_data["description"] = fileHeader.Filename[:len(fileHeader.Filename)-len(extension)]
		}

		// err = c.Object.ValidateData(true, false, new_data)
		// if err != nil {
		// 	return nil, err
		// }

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			return nil, err
		}

		err = os.MkdirAll(upload_folder, os.ModePerm)
		if err != nil {
			return nil, err
		}

		dst, err := os.Create(fmt.Sprintf("%v/%v%s", upload_folder, new_file_name, extension))
		if err != nil {
			return nil, err
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			return nil, err
		}

		// borramos el campo files
		delete(new_data, c.Fields[6].Name)

		err = c.CreateObjectsInDB(c.Object.Name, new_data)
		if err != nil {
			//borrar archivo creado en disco
			file_delete := dst.Name()
			dst.Close()

			del_err := os.Remove(file_delete)
			if del_err != nil {
				return nil, fmt.Errorf("error %v y al borrar archivo: %v", err, del_err)
			}

			return nil, err
		}

		out := map[string]string{
			"id_file":     new_file_name,
			"description": new_data["description"],
		}

		data_out = append(data_out, out)
	}

	return data_out, nil
}
