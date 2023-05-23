package file

import (
	"github.com/cdvelop/objectdb"
)

type File struct {
	*objectdb.Connection
	extensions  string // ej: ".jpg, .png, .jpeg"
	max_files   int64  // ej 10 archivos
	max_kb_size int64  // ej 100 kb
}
