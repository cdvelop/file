package file_test

import (
	"github.com/cdvelop/file"
	"github.com/cdvelop/model"
	"github.com/cdvelop/testools"
)

type dataTest struct {
	file *file.File

	field_name string   //ej: endoscopia, voucher, foto_mascota, foto_usuario
	files      []string //ej: "gatito.jpg, perro.png"
	file_type  string   //ej: imagen,video,document,pdf
	max_files  string
	max_size   string

	expected string

	*model.Module

	*testools.Request
}
