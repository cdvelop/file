package file_test

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/cdvelop/model"
)

func readTest(endpoint string, server *httptest.Server, response model.Response) {
	for _, dta_resp := range response.Data {

		// Construir la URL de la solicitud GET con el parámetro "id_file"
		url := fmt.Sprintf(server.URL+endpoint+"?id=%s", dta_resp["id_file"])

		// fmt.Println("URL: ", url)

		// Hacer la solicitud GET con http.Get
		get_response, err := http.Get(url)
		if err != nil {
			// Manejar errores de conexión
			fmt.Println("Error al hacer la solicitud GET:", err)
			return
		}
		defer get_response.Body.Close()

		// Leer la respuesta recibida
		_, err = io.ReadAll(get_response.Body)
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println("*** RESPUESTA SOLICITUD READ: ")
		// fmt.Println("CÓDIGO ESTATUS:", get_response.StatusCode)

		if get_response.StatusCode != 200 {
			log.Fatalln("se esperaba código 200 respuesta: ", get_response.StatusCode)
		}

		// Procesar la respuesta recibida
		// fmt.Println("TAMAÑO ARCHIVO RESPUESTA GET: ", len(bodyBytes))

	}
}
