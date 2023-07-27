package file

import (
	"strconv"
	"strings"

	"github.com/cdvelop/gotools"
	"github.com/cdvelop/input"
	"github.com/cdvelop/model"
)

// conf:
// filetype:video, pdf, document. default imagen
// root_folder:static_files default "app_files"
// max_files:1, 4, 6.. default 6
// max_kb_size:100, 400 default 50
func New(m *model.Module, db model.DataBaseAdapter, conf ...string) *config {

	c := config{
		Object: model.Object{
			Name:           "file",
			TextFieldNames: []string{"field_name", "description"},
			Fields: []model.Field{
				{Name: "id_file", Legend: "Id", Input: input.Pk()},
				{Name: "module_name", Legend: "Modulo", Input: input.TextNumCode()},
				{Name: "field_name", Legend: "Carpeta Campo", Input: input.TextOnly()},
				{Name: "folder_id", Legend: "Carpeta Registro", Input: input.Pk()},
				{Name: "description", Legend: "Descripción", Input: input.Text(`title="Min. 3 Max. 50 caracteres"`, `pattern="^[A-Za-zÑñáéíóú ]{3,50}$"`), SkipCompletionAllowed: true},
				{Name: "file_path", Legend: "Ubicación", Input: input.FilePath(), SkipCompletionAllowed: true},
				{Name: "files", NotRequiredInDB: true, Legend: "Archivos", Input: input.Text()},
			},
			Module:           m,
			BackendRequest:   model.BackendRequest{},
			FrontendResponse: model.FrontendResponse{},
		}, //nota: al no declarar el puntero se pierde posteriormente
		DataBaseAdapter: db,
		name:            "file",
		filetype:        "imagen",
		root_folder:     "app_files",
		extensions:      ".jpg, .png, .jpeg, .webp",
		max_files:       6,
		max_kb_size:     50,
	}

	for _, option := range conf {

		switch {

		case strings.Contains(option, "root_folder:"):
			gotools.ExtractTwoPointArgument(option, &c.root_folder)

		case strings.Contains(option, "filetype:"):
			gotools.ExtractTwoPointArgument(option, &c.filetype)

			switch c.filetype {
			case "video":
				c.extensions = ".avi, .mkv, .mp4"
			case "document":
				c.extensions = ".doc, .xlsx, .txt"
			case "pdf":
				c.extensions = ".pdf"
			}

		case strings.Contains(option, "max_files:"):
			var max_files string
			gotools.ExtractTwoPointArgument(option, &max_files)

			num, err := strconv.ParseInt(max_files, 10, 64)
			if err != nil {
				gotools.ShowErrorAndExit("Error al convertir max_files la cadena: " + max_files + " " + err.Error())
			}
			c.max_files = num

		case strings.Contains(option, "max_kb_size:"):
			var max_kb_size string
			gotools.ExtractTwoPointArgument(option, &max_kb_size)

			num, err := strconv.ParseInt(max_kb_size, 10, 64)
			if err != nil {
				gotools.ShowErrorAndExit("Error al convertir max_kb_size la cadena: " + max_kb_size + " " + err.Error())
			}
			c.max_kb_size = num

		}
	}

	c.maximum_file_size = int64(float64(c.max_files*c.max_kb_size*1024) * 1.05)

	c.Object.BackendRequest = model.BackendRequest{
		CreateApi: c,
		ReadApi:   c,
		UpdateApi: c,
		DeleteApi: c,
		FileApi:   c,
	}

	// for _, field := range c.Fields {
	// 	fmt.Println("CAMPO: ", field.Name)

	c.Object.AddModule(m)

	err := db.CreateTablesInDB(&c.Object)
	if err != nil {
		gotools.ShowErrorAndExit(err.Error())
	}

	return &c
}

func (c config) Name() string {
	return c.name
}

func (config) HtmlName() string {
	return "file"
}

func (c config) MaximumFileSize() int64 {
	return c.maximum_file_size
}
