package file

import (
	"strconv"
	"strings"

	"github.com/cdvelop/gotools"
	"github.com/cdvelop/input"
	"github.com/cdvelop/model"
	"github.com/cdvelop/object"
	. "github.com/cdvelop/output"
	"github.com/cdvelop/unixid"
)

// conf:
// field_name:voucher,user_photo,boleta... default file
// filetype:video, pdf, document. default imagen
// root_folder:static_files default "app_files"
// max_files:1, 4, 6.. default 6
// max_kb_size:100, 400 default 50
func New(m *model.Module, db model.DataBaseAdapter, id model.IdHandler, conf ...string) (*File, error) {

	err := m.AddInputs([]*model.Input{unixid.InputPK(), input.TextNumCode(), input.TextNum(), input.FilePath(), input.Text()}, "file pkg")
	if err != nil {
		return nil, err
	}

	f := File{}

	// crear objeto
	err = object.New(&f, m)
	if err != nil {
		return nil, err
	}

	f.db = db
	f.idh = id
	f.filetype = "imagen"
	f.root_folder = "app_files"
	f.extensions = ".jpg, .png, .jpeg, .webp"
	f.max_files = 6
	f.max_kb_size = 50

	var field_name string

	for _, option := range conf {

		switch {

		case strings.Contains(option, "field_name:"):
			gotools.ExtractTwoPointArgument(option, &field_name)

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
				ShowErrorAndExit("Error al convertir max_files la cadena: " + max_files + " " + err.Error())
			}
			f.max_files = num

		case strings.Contains(option, "max_kb_size:"):
			var max_kb_size string
			gotools.ExtractTwoPointArgument(option, &max_kb_size)

			num, err := strconv.ParseInt(max_kb_size, 10, 64)
			if err != nil {
				ShowErrorAndExit("Error al convertir max_kb_size la cadena: " + max_kb_size + " " + err.Error())
			}
			f.max_kb_size = num

		}
	}

	if field_name == "" {
		return nil, model.Error(`error field_name:"nombre_campo" no ingresado`)
	}

	f.Object.Name += "." + field_name

	f.maximum_file_size = int64(float64(f.max_files*f.max_kb_size*1024) * 1.05)

	// err = db.CreateTablesInDB([]*model.Object{&o}, nil)
	// if err != nil {
	// 	return nil, err
	// }

	//nota: al no declarar punteros se pierden posteriormente

	return &f, nil
}

func (File) HtmlName() string {
	return "file"
}

func (f File) MaximumFileSize() int64 {
	return f.maximum_file_size
}
