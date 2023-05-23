package file

var root_folder string //ej: "./app_files/"

// root_folder ej: "./app_files/", "./static_files/"
// MODULE_NAME
// FIELD_NAME ej: voucher, endoscopia, foto, article_photo
// LEGEND ej: Comprobante, Endoscopia, Fotos, Imágenes Articulo
// extensions: ej: ".jpg, .png, .jpeg"
// max_files ej: 1,4 6..
// max_kb_size ej: 100, 400
// https://developer.mozilla.org/es/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Common_types
func New(rootFolder, extensions string, max_files, max_kb_size int64) *File {
	// ./root_folder/MODULE_NAME/FIELD_NAME/
	// URL EXAMPLE : ".app_files/medicalhistory/endoscopia/"
	if root_folder == "" {
		root_folder = rootFolder
	}

	f := File{
		extensions:  extensions,
		max_files:   max_files,
		max_kb_size: max_kb_size,
	}

	f.checkDataBase()

	return &f
}

func (File) Name() string {
	return "file"
}

func (File) HtmlName() string {
	return "file"
}
