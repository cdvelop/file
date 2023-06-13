package file

import "github.com/cdvelop/sqlite"

// root_folder ej: "./app_files", "./static_files"
// extensions: ej: ".jpg, .png, .jpeg"
// max_files ej: 1,4 6..
// max_kb_size ej: 100, 400
func New(rootFolder, extensions string, max_files, max_kb_size int64) *File {

	f := File{
		root_folder: rootFolder,
		extensions:  extensions,
		max_files:   max_files,
		max_kb_size: max_kb_size,
	}

	f.Connection = sqlite.NewConnection(f.root_folder, "stored_files_index.db", false, f.Object())

	return &f
}

func (File) Name() string {
	return "file"
}

func (File) HtmlName() string {
	return "file"
}
