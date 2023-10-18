package file

import (
	"fmt"
	"net/http"

	"github.com/cdvelop/model"
)

// CreateFile upload files http handler
func (f File) CreateFile(u *model.User, r *http.Request, params map[string]string) ([]map[string]string, error) {

	fmt.Println("*** FILES=", f.Files)
	upload_folder := f.buildUploadFolder(params)

	// fmt.Println("FILE FORM NAME: ", params[f.FieldFiles], " upload_folder", upload_folder)

	files := r.MultipartForm.File[f.Files]
	if len(files) == 0 {
		return nil, fmt.Errorf("CreateFile error no hay archivos detectados")
	}

	if len(files) > int(f.max_files) {
		return nil, fmt.Errorf("error se pretende subir %v archivos, pero el máximo permitido es: %v", len(files), f.max_files)
	}

	return f.processFiles(files, upload_folder, params)
}

func (f File) buildUploadFolder(new_data map[string]string) string {
	return f.root_folder + "/" + new_data[f.Module_name] + "/" + new_data[f.Field_name] + "/" + new_data[f.Folder_id]
}
