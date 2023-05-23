package file

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cdvelop/model"
)

// request_type ej: create, read, update,delete, error
func (f File) response(w http.ResponseWriter, code int, request_type, message, for_module string, data_out ...map[string]string) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)

	if for_module == "" {
		for_module = f.Name()
	}

	r := model.Response{
		Type:    request_type,
		Data:    data_out,
		Object:  f.Name(),
		Module:  for_module,
		Message: message,
	}

	jsonBytes, err := json.Marshal(r)
	if err != nil {
		fmt.Fprintln(w, `{"Type":"error", "Message":"`+err.Error()+`"}`)
		return
	}

	w.Write(jsonBytes)
}

func (f File) error(w http.ResponseWriter, message string, code int) {
	f.response(w, code, "error", message, "")
}
