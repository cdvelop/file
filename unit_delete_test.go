package file_test

import (
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/cdvelop/model"
)

func deleteTest(endpoint string, server *httptest.Server, rq model.Response) {

	// Crear una solicitud DELETE

	response := newRequest(http.MethodDelete, server.URL+endpoint, &rq)

	// fmt.Println("*** RESPUESTA SOLICITUD DELETE: ")
	// Procesar la respuesta recibida
	// fmt.Println("respuesta delete:", response)
	if response.Message != "Archivo(s) borrado(s) con éxito." {
		log.Fatalln("Se esperaba Archivo(s) borrado(s) con éxito. se obtuvo: ", response.Message)
	}

}
