package file

import (
	"net/http"

	json "github.com/fxamacker/cbor/v2"

	"github.com/cdvelop/model"
)

func (f File) Update(w http.ResponseWriter, r *http.Request) {

	// Decodificar los datos JSON recibidos
	var requestData model.Response
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		f.error(w, "Error Json Decode", http.StatusBadRequest)
		return
	}

	// Hacer algo con los datos recibidos
	message, ok := f.UpdateObjects(f.Object().Name, requestData.Data...)
	if !ok {
		f.error(w, message, http.StatusBadRequest)
		return
	}

	f.response(w, http.StatusOK, "UPDATE", message, "moduleTest")

}
