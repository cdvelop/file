package file

import (
	"github.com/cdvelop/input"
	"github.com/cdvelop/model"
	"github.com/cdvelop/objectdb"
)

type File struct {
	root_folder string //ej: "./app_files"
	*objectdb.Connection
	extensions  string // ej: ".jpg, .png, .jpeg"
	max_files   int64  // ej 10 archivos
	max_kb_size int64  // ej 100 kb
}

func (f File) Object() model.Object {
	return model.Object{
		Name:           "file",
		TextFieldNames: []string{"module_name", "field_name"},
		Fields: []model.Field{
			{Name: "id_file", Legend: "Id", Input: input.Pk()},
			{Name: "module_name", Legend: "Modulo", Input: input.TextNumCode()},
			{Name: "field_name", Legend: "Carpeta Campo", Input: input.TextOnly()},
			{Name: "folder_id", Legend: "Carpeta Registro", Input: input.Pk()},
			{Name: "description", Legend: "Descripción", Input: input.Text(`title="Min. 3 Max. 50 caracteres"`, `pattern="^[A-Za-zÑñáéíóú ]{3,50}$"`), SkipCompletionAllowed: true},
			{Name: "file_path", Legend: "Ubicación", Input: input.FilePath(), SkipCompletionAllowed: true},
		},
	}
}
