package file

import (
	"fmt"
	"net/http"
	"strings"
)

const error_message_big_file = "error archivo muy grande. tamaño máximo admitido: %v kb"

// Create upload files http handler
func (f File) Create(w http.ResponseWriter, r *http.Request) {

	max_size := int64(float64(f.max_files*f.max_kb_size*1024) * 1.05)

	r.Body = http.MaxBytesReader(w, r.Body, max_size) // 220 KB

	err := r.ParseMultipartForm(max_size)
	if err != nil {
		if strings.Contains(err.Error(), "multipart") {
			f.error(w, err.Error(), http.StatusNotAcceptable)
		} else {
			f.error(w, fmt.Sprintf(error_message_big_file, f.max_kb_size), http.StatusNotAcceptable)
		}
		return
	}

	new_data := f.parseFormData(r)

	upload_folder := f.buildUploadFolder(new_data)

	files := r.MultipartForm.File[new_data["field_name"]]
	if len(files) == 0 {
		f.error(w, "no hay archivos detectados", http.StatusNotAcceptable)
		return
	}

	if len(files) > int(f.max_files) {
		f.error(w, fmt.Sprintf("error se pretende subir %v archivos, pero el máximo permitido es: %v", len(files), f.max_files), http.StatusNotAcceptable)
		return
	}

	data_out := f.processFiles(files, max_size, upload_folder, new_data, w)

	f.response(w, http.StatusOK, "create", "Carga exitosa", new_data["module_name"], data_out...)
}
