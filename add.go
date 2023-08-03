package file

import (
	"strconv"
	"strings"

	"github.com/cdvelop/gotools"
	"github.com/cdvelop/input"
	"github.com/cdvelop/model"
)

// conf:
// api_name:voucher,user_photo,boleta... default file
// filetype:video, pdf, document. default imagen
// root_folder:static_files default "app_files"
// max_files:1, 4, 6.. default 6
// max_kb_size:100, 400 default 50
func New(m *model.Module, db model.DataBaseAdapter, conf ...string) *File {

	f := File{
		Name:             "file",
		FieldIdFile:      "id_file",
		FieldModuleName:  "module_name",
		FieldName:        "field_name",
		FieldFolderId:    "folder_id",
		FieldDescription: "description",
		FieldFilePath:    "file_path",
		FieldFiles:       "files",

		object: nil,
		db:     db,

		filetype:    "imagen",
		root_folder: "app_files",
		extensions:  ".jpg, .png, .jpeg, .webp",
		max_files:   6,
		max_kb_size: 50,
	}

	var api_name string

	for _, option := range conf {

		switch {

		case strings.Contains(option, "api_name:"):
			gotools.ExtractTwoPointArgument(option, &api_name)

		case strings.Contains(option, "root_folder:"):
			gotools.ExtractTwoPointArgument(option, &f.root_folder)

		case strings.Contains(option, "filetype:"):
			gotools.ExtractTwoPointArgument(option, &f.filetype)

			switch f.filetype {
			case "video":
				f.extensions = ".avi, .mkv, .mp4"
			case "document":
				f.extensions = ".doc, .xlsx, .txt"
			case "pdf":
				f.extensions = ".pdf"
			}

		case strings.Contains(option, "max_files:"):
			var max_files string
			gotools.ExtractTwoPointArgument(option, &max_files)

			num, err := strconv.ParseInt(max_files, 10, 64)
			if err != nil {
				gotools.ShowErrorAndExit("Error al convertir max_files la cadena: " + max_files + " " + err.Error())
			}
			f.max_files = num

		case strings.Contains(option, "max_kb_size:"):
			var max_kb_size string
			gotools.ExtractTwoPointArgument(option, &max_kb_size)

			num, err := strconv.ParseInt(max_kb_size, 10, 64)
			if err != nil {
				gotools.ShowErrorAndExit("Error al convertir max_kb_size la cadena: " + max_kb_size + " " + err.Error())
			}
			f.max_kb_size = num

		}
	}

	f.maximum_file_size = int64(float64(f.max_files*f.max_kb_size*1024) * 1.05)

	o := model.Object{
		Name:           f.Name,
		TextFieldNames: []string{f.FieldName, f.FieldDescription},
		Fields: []model.Field{
			{Name: f.FieldIdFile, Legend: "Id", Input: input.Pk()},
			{Name: f.FieldModuleName, Legend: "Modulo", Input: input.TextNumCode()},
			{Name: f.FieldName, Legend: "Carpeta Campo", Input: input.TextOnly()},
			{Name: f.FieldFolderId, Legend: "Carpeta Registro", Input: input.Pk()},
			{Name: f.FieldDescription, Legend: "Descripción", Input: input.Text(`title="Min. 3 Max. 50 caracteres"`, `pattern="^[A-Za-zÑñáéíóú ]{3,50}$"`), SkipCompletionAllowed: true},
			{Name: f.FieldFilePath, Legend: "Ubicación", Input: input.FilePath(), SkipCompletionAllowed: true},
			{Name: f.FieldFiles, NotRequiredInDB: true, Legend: "Archivos", Input: input.Text()},
		},
		BackendRequest: model.BackendRequest{
			CreateApi: nil,
			ReadApi:   f,
			UpdateApi: f,
			DeleteApi: f,
			FileApi:   f,
		},
		FrontendResponse: model.FrontendResponse{},
	}

	f.object = &o
	o.AddModule(m, api_name)

	err := db.CreateTablesInDB(&o)
	if err != nil {
		gotools.ShowErrorAndExit(err.Error())
	}

	//nota: al no declarar punteros se pierden posteriormente

	// fmt.Println("Api objeto: ", o.Api())

	return &f
}

func (f File) Object() *model.Object {
	return f.object
}

func (File) HtmlName() string {
	return "file"
}

func (f File) MaximumFileSize() int64 {
	return f.maximum_file_size
}
