package file

import (
	"github.com/cdvelop/model"
)

type config struct {
	model.Object
	model.DataBaseAdapter

	name              string // nombre principal archivo
	filetype          string //imagen, video, document
	root_folder       string //ej: "./app_files"
	extensions        string // ej: ".jpg, .png, .jpeg"
	max_files         int64  // ej 10 archivos
	max_kb_size       int64  // ej 100 kb
	maximum_file_size int64
}

// type table struct{
// 	name string
// 	id_file string
// 	module_name string
// 	field_name string
// 	folder_id string
// 	description string
// 	file_path string
// }
