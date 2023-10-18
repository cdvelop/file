package file

import (
	"github.com/cdvelop/model"
)

type File struct {
	//table
	Id_file     string `Legend:"File Id" Input:"InputPK"`                                                   //0 id_file
	Module_name string `Legend:"Modulo" Input:"TextNumCode"`                                                //1 module_name
	Field_name  string `Legend:"Carpeta Campo" Input:"TextNum" PrincipalField:"ok"`                         //2 field_name
	Folder_id   string `Legend:"Carpeta Registro" Input:"InputPK"`                                          //3 folder_id
	Description string `Legend:"Descripción" Input:"Text" PrincipalField:"ok" SkipCompletionAllowed:"true"` //4 description
	File_path   string `Legend:"Ubicación" Input:"FilePath" SkipCompletionAllowed:"true"`                   //5 file_path

	//config
	Files string `Input:"Text" NotRequiredInDB:"true"` //nombre por defecto input html = files

	Object *model.Object
	db     model.DataBaseAdapter

	idh model.IdHandler

	filetype          string //imagen, video, document
	root_folder       string //ej: "./app_files"
	extensions        string // ej: ".jpg, .png, .jpeg"
	max_files         int64  // ej 10 archivos
	max_kb_size       int64  // ej 100 kb
	maximum_file_size int64
}
