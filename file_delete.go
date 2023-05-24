package file

import (
	"net/http"
	"os"

	json "github.com/fxamacker/cbor/v2"

	"github.com/cdvelop/model"
)

func (f File) Delete(w http.ResponseWriter, r *http.Request) {

	// Decodificar los datos JSON recibidos
	var requestData model.Response
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		f.error(w, "Error Json Decode", http.StatusBadRequest)
		return
	}

	// Hacer algo con los datos recibidos
	// fmt.Println("DATA RECIBIDA PARA ELIMINAR: ", requestData)

	message, ok := f.DeleteFilesFromHDDandDB(requestData.Data...)
	if !ok {
		f.error(w, message, http.StatusBadRequest)
		return
	}

	f.response(w, http.StatusOK, "DELETE", message, "moduleTest")

}

func (f File) DeleteFilesFromHDDandDB(data_with_id_file ...map[string]string) (string, bool) {

	// fmt.Println("Eliminando recurso del disco:", data_with_id_file)

	// recuperar toda la info de la db de los archivos antes de eliminarlos
	recovered_data, message, ok := f.GetAllDataFromDB(data_with_id_file...)
	if !ok {
		return message, false
	}

	message, ok = f.DeleteObjects(f.Object().Name, data_with_id_file...)
	if !ok {
		return message, false
	}

	for _, data := range recovered_data {
		// Borrar archivos
		err := os.Remove(data["file_path"])
		if err != nil {
			return err.Error(), false
		}
	}

	// reemplazamos la data recuperada para su uso por ej en el frontend
	copy(data_with_id_file, recovered_data)

	return "Archivo(s) borrado(s) con éxito.", true

}
