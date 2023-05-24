package file

import (
	"net/http"
)

func (f File) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		f.Create(w, r)

	case http.MethodGet:

		switch {
		case r.URL.Query().Get("id") != "":
			id_file := r.URL.Query().Get("id")
			f.ReadFile(id_file, w, r)

		case r.URL.Query().Get("read_one") != "":
			id_file := r.URL.Query().Get("read_one")
			f.ReadOne(id_file, w, r)

		case r.URL.Query().Get("read_all") != "":
			folder_id := r.URL.Query().Get("read_all")
			f.ReadAll(folder_id, w, r)

		default:
			f.error(w, "error parámetro de lectura no especificado", http.StatusNotFound)
		}

	case http.MethodPatch:
		f.Update(w, r)

	case http.MethodDelete:
		f.Delete(w, r)

	default:
		f.error(w, "Método No permitido", http.StatusMethodNotAllowed)

	}
}
