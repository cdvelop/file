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

func (f File) processFiles(files []*multipart.FileHeader, max_size int64, upload_folder string, new_data map[string]string, w http.ResponseWriter) []map[string]string {
	data_out := []map[string]string{}
	for _, fileHeader := range files {
		if fileHeader.Size > max_size {
			f.error(w, fmt.Sprintf(error_message_big_file, f.max_kb_size), http.StatusNotAcceptable)
			return data_out
		}

		file, err := fileHeader.Open()
		if err != nil {
			f.error(w, err.Error(), http.StatusNotAcceptable)
			return data_out
		}
		defer file.Close()

		extension := f.getExtension(fileHeader)

		if !strings.Contains(f.extensions, extension) {
			f.error(w, "formato de archivo no valido solo se admiten: "+f.extensions, http.StatusBadRequest)
			return data_out
		}
		extension = filepath.Ext(fileHeader.Filename)

		new_file_name := getNewID()
		new_data["id_file"] = new_file_name
		new_data["file_path"] = upload_folder + "/" + new_file_name + extension

		if len(fileHeader.Filename) > 5 {
			new_data["description"] = fileHeader.Filename[:len(fileHeader.Filename)-len(extension)]
		}

		if mg, ok := f.Object().ValidateData(true, new_data); !ok {
			f.error(w, mg, http.StatusNotAcceptable)
			return data_out
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			f.error(w, err.Error(), http.StatusInternalServerError)
			return data_out
		}

		err = os.MkdirAll(upload_folder, os.ModePerm)
		if err != nil {
			f.error(w, err.Error(), http.StatusInternalServerError)
			return data_out
		}

		dst, err := os.Create(fmt.Sprintf("%v/%v%s", upload_folder, new_file_name, extension))
		if err != nil {
			f.error(w, err.Error(), http.StatusBadRequest)
			return data_out
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			f.error(w, err.Error(), http.StatusInternalServerError)
			return data_out
		}

		mg, ok := f.CreateObjects(f.Object().Name, new_data)
		if !ok {
			f.error(w, mg, http.StatusInternalServerError)
			return data_out
		}

		out := map[string]string{
			"id_file":     new_file_name,
			"description": new_data["description"],
		}

		data_out = append(data_out, out)
	}
	return data_out
}
