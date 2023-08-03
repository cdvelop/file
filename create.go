package file

import (
	"fmt"
	"net/http"
)

// CreateFile upload files http handler
func (f File) CreateFile(r *http.Request, params map[string]string) ([]map[string]string, error) {

	upload_folder := f.buildUploadFolder(params)

	// fmt.Println("FILE FORM NAME: ", params[f.FieldFiles], " upload_folder", upload_folder)

	files := r.MultipartForm.File[f.FieldFiles]
	if len(files) == 0 {
		return nil, fmt.Errorf("CreateFile error no hay archivos detectados")
	}

	if len(files) > int(f.max_files) {
		return nil, fmt.Errorf("error se pretende subir %v archivos, pero el máximo permitido es: %v", len(files), f.max_files)
	}

	return f.processFiles(files, upload_folder, params)
}

func (f File) buildUploadFolder(new_data map[string]string) string {
	return f.root_folder + "/" + new_data[f.FieldModuleName] + "/" + new_data[f.FieldName] + "/" + new_data[f.FieldFolderId]
}
