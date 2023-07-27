package file

import (
	"fmt"
	"net/http"
)

// UploadFile upload files http handler
func (c config) UploadFile(r *http.Request, params map[string]string) ([]map[string]string, error) {

	upload_folder := c.buildUploadFolder(params)

	fmt.Println("FILE FORM NAME: ", params[c.Fields[6].Name], " upload_folder", upload_folder)

	files := r.MultipartForm.File[c.Fields[6].Name]
	if len(files) == 0 {
		return nil, fmt.Errorf("UploadFile error no hay archivos detectados")
	}

	if len(files) > int(c.max_files) {
		return nil, fmt.Errorf("error se pretende subir %v archivos, pero el máximo permitido es: %v", len(files), c.max_files)
	}

	return c.processFiles(files, upload_folder, params)
}

func (f config) buildUploadFolder(new_data map[string]string) string {
	return f.root_folder + "/" + new_data[f.Fields[1].Name] + "/" + new_data[f.Fields[2].Name] + "/" + new_data[f.Fields[3].Name]
}
