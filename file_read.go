package file

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cdvelop/dbtools"
)

func (f File) ReadFile(id_file string, w http.ResponseWriter, r *http.Request) {

	file_path, ok := f.GetPathFileByID(id_file)
	if !ok {
		f.error(w, file_path, http.StatusNotFound)
		return
	}
	// fmt.Println("DATA SOLICITADA: ", out)
	// Servir el archivo encontrado
	http.ServeFile(w, r, file_path)

}
func (f File) ReadOne(id_file string, w http.ResponseWriter, r *http.Request) {

	var message string
	data := f.ReadObject(f.Name(), map[string]string{"id_file": id_file})
	delete(data, "file_path")
	// fmt.Println("DATA SOLICITADA: ", data)
	// Servir json
	f.response(w, 200, "read", message, "file", data)

}

func (f File) GetPathFileByID(id_file string) (string, bool) {

	out, err := f.QueryOne(fmt.Sprintf("SELECT file_path FROM %v WHERE id_%v ='%v';", f.Object().Name, f.Object().Name, id_file))
	if err != nil {
		return fmt.Sprintf("Error al Obtener Data Tabla [%v]\n%v", f.Object().Name, err.Error()), false
	}

	return out["file_path"], true
}

func (f File) GetAllDataFromDB(data_with_id_file ...map[string]string) ([]map[string]string, string, bool) {
	f.Open()
	defer f.Close()
	// crear argumentos ids de slice de interfaces
	args_ids := make([]interface{}, len(data_with_id_file))
	for i, data := range data_with_id_file {
		args_ids[i] = data["id_file"]
	}

	query := "SELECT * FROM " + f.Object().Name + " WHERE id_file IN (?" + strings.Repeat(",?", len(data_with_id_file)-1) + ")"

	rows, err := f.Query(query, args_ids...)
	if err != nil {
		return nil, err.Error(), false
	}

	result, err := dbtools.FetchAll(rows)
	if err != nil {

		return nil, err.Error(), false
	}

	return result, "", true
}
