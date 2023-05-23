package file_test

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/cdvelop/model"
)

func deleteTest(endpoint string, server *httptest.Server, read_response *model.Response) {

	requestDataBytes, err := json.Marshal(read_response)
	if err != nil {
		log.Fatal(err)
	}

	// Crear una solicitud DELETE
	req, err := http.NewRequest("POST", server.URL+endpoint, bytes.NewBuffer(requestDataBytes))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Action-Type", "delete")

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

	var response model.Response
	err = json.Unmarshal(resp, &response)
	if err != nil {
		log.Fatal("Error al decodificar datos JSON:", err)
		return
	}

	// fmt.Println("*** RESPUESTA SOLICITUD DELETE: ")
	// Procesar la respuesta recibida
	// fmt.Println("Código de estado de la respuesta delete:", response)

}
