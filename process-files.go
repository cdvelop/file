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

func (f File) processFiles(files []*multipart.FileHeader, upload_folder string, new_data map[string]string) ([]map[string]string, error) {
	data_out := []map[string]string{}
	for _, fileHeader := range files {
		if fileHeader.Size > f.maximum_file_size {
			return nil, fmt.Errorf(fmt.Sprintf("error archivo(s) excede(n) el tamaño admitido de: %v kb", f.max_kb_size), http.StatusNotAcceptable)
		}

		file, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}
		defer file.Close()

		extension := f.getExtension(fileHeader)

		if !strings.Contains(f.extensions, extension) {
			return nil, fmt.Errorf("formato de archivo no valido solo se admiten: %v", f.extensions)
		}
		extension = filepath.Ext(fileHeader.Filename)

		new_file_name := f.idh.GetNewID()

		new_data[f.Id_file] = new_file_name
		new_data[f.File_path] = upload_folder + "/" + new_file_name + extension

		if len(fileHeader.Filename) > 5 {
			new_data[f.Description] = fileHeader.Filename[:len(fileHeader.Filename)-len(extension)]
		}

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
		delete(new_data, f.Files)

		err = f.db.CreateObjectsInDB(f.Object.Table, new_data)
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
			f.Id_file:     new_file_name,
			f.Description: new_data[f.Description],
		}

		data_out = append(data_out, out)
	}

	return data_out, nil
}
