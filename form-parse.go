package file

import "net/http"

func (f File) parseFormData(r *http.Request) map[string]string {
	new_data := make(map[string]string)
	for i, field := range f.Object().Fields {
		if i > 0 && i <= 4 { // saltarse id y file_path
			new_data[field.Name] = r.FormValue(field.Name)
		}
	}
	return new_data
}

func (f File) buildUploadFolder(new_data map[string]string) string {
	return f.root_folder + "/" + new_data["module_name"] + "/" + new_data["field_name"] + "/" + new_data["folder_id"]
}
