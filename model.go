package file

import (
	"github.com/cdvelop/model"
)

type File struct {
	//table
	Name             string // file
	FieldIdFile      string //0 id_file
	FieldModuleName  string //1 module_name
	FieldName        string //2 field_name
	FieldFolderId    string //3 folder_id
	FieldDescription string //4 description
	FieldFilePath    string //5 file_path
	FieldFiles       string //6 files

	//config

	object *model.Object
	db     model.DataBaseAdapter

	filetype          string //imagen, video, document
	root_folder       string //ej: "./app_files"
	extensions        string // ej: ".jpg, .png, .jpeg"
	max_files         int64  // ej 10 archivos
	max_kb_size       int64  // ej 100 kb
	maximum_file_size int64
}
