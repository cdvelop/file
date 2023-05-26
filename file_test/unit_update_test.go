package file_test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/cdvelop/model"
)

func updateTest(endpoint string, server *httptest.Server, rq model.Response) bool {

	// fmt.Println("DATA A ACTUALIZAR: ", rq)
	for i, data := range rq.Data {
		data["description"] = "perro"
		rq.Data[i] = data
	}

	update_resp := newRequest(http.MethodPatch, server.URL+endpoint, &rq)

	// fmt.Println("\n*** RESPUESTA SOLICITUD UPDATE: ", update_resp)
	if update_resp.Message != "Actualización Exitosa" {
		log.Fatalln("Error se esperaba Actualización Exitosa se obtuvo: ", update_resp.Message)
	}

	for _, data := range rq.Data {

		// Construir la URL de la solicitud GET con el parámetro "id_file"
		url := fmt.Sprintf(server.URL+endpoint+"?read_one=%s", data["id_file"])

		read_resp := newRequest(http.MethodGet, url, nil)

		// fmt.Println(i, "- RESPUESTA SOLICITUD READ: ", read_resp)

		for _, data := range read_resp.Data {

			if data["description"] != "perro" {
				log.Fatalln("Error se esperaba en description perro se obtuvo: ", data["description"])
			}
		}

	}

	return true
}
