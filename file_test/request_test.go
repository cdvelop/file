package file_test

import (
	"bytes"
	"io"
	"log"
	"net/http"

	json "github.com/fxamacker/cbor/v2"

	"github.com/cdvelop/model"
)

// method: "POST","GET" "PATCH","DELETE"
func newRequest(method, endpoint string, data_in *model.Response) (response model.Response) {
	var err error
	var bytes_data []byte
	if data_in != nil {
		bytes_data, err = json.Marshal(data_in)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Crear una solicitud UPDATE
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(bytes_data))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	resp, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(resp, &response)
	if err != nil {
		log.Fatal("Error al decodificar datos JSON:", err)
		return
	}

	return
}
