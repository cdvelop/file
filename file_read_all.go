package file

import (
	"fmt"
	"net/http"

	"github.com/cdvelop/dbtools"
)

func (f File) readFiles(folder_id string, w http.ResponseWriter, r *http.Request) {
	data_out, msg, ok := f.ReadAllFilesID(folder_id)
	if !ok {
		f.error(w, msg, http.StatusNotFound)
		return
	}

	var message string
	total := len(data_out)
	if total != 0 {
		message = fmt.Sprintf("total archivos encontrados: %v", total)
	} else {
		message = fmt.Sprintf("carpeta id: %v no contiene archivos", folder_id)
	}
	// fmt.Println("DATA SOLICITADA: ", out)

	f.response(w, 200, "read", message, "file", data_out...)
}

func (f File) ReadAllFilesID(folder_id string) ([]map[string]string, string, bool) {

	query := `SELECT id_file FROM file WHERE folder_id = ?`
	args := []interface{}{folder_id}

	rows, err := f.Query(query, args...)
	if err != nil {
		return nil, err.Error(), false
	}

	result, err := dbtools.FetchAll(rows)
	if err != nil {

		return nil, err.Error(), false
	}

	return result, "", true

}
