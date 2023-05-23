package file

import (
	"net/http"
	"strings"
)

func (f File) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Action type ej: create,READ,DELETE,UPDATE
	action := strings.Join(r.Header["Action-Type"][:], " ")

	// fmt.Println("ACTION FILE: ", action)

	switch action {
	case "create", "update", "delete":

		if r.Method != http.MethodPost {
			f.error(w, "Método No permitido", http.StatusMethodNotAllowed)
			return
		}

		switch action {
		case "create":
			f.createFile(w, r)
		case "update":
			f.updateFile(w, r)
		case "delete":
			f.deleteFile(w, r)
		}

	default:

		if r.Method == http.MethodGet && action == "" {

			// Obtener el parámetro de URL "id_file"
			id_file := r.URL.Query().Get("id_file")
			if id_file != "" {

				f.readFile(id_file, w, r)

			} else {
				// Obtener el parámetro de URL "folder_id"
				folder_id := r.URL.Query().Get("folder_id")
				if folder_id != "" {
					f.readFiles(folder_id, w, r)
				} else {
					f.error(w, "ID de archivo o carpeta no especificado", http.StatusNotFound)
				}

			}

		} else {

			f.error(w, "Error Action-Type en headers. admitidos: create, delete, update. GET no requerido.", http.StatusBadRequest)
		}

	}
}
