package file_test

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	json "github.com/fxamacker/cbor/v2"

	"github.com/cdvelop/file"
	"github.com/cdvelop/model"
)

var (
	dataHttp = map[string]struct {
		endpoint   string   //ej: upload/files download/files
		method     string   //ej: "PUT","GET"
		files      []string //ej: "gatito.jpg, perro.png"
		extensions string   //ej: ".jpg, .png"
		max_files  int64
		max_size   int64
		expected   string
	}{
		"gatito 220kb y dino 36kb ok":                         {"/file", http.MethodPost, []string{"dino.png", "gatito.jpeg"}, ".png, .jpg", 2, 262, "create"},
		"gatito 220kb ok":                                     {"/file", http.MethodPost, []string{"gatito.jpeg"}, ".jpg", 1, 220, "create"},
		"tamaño gatito 220kb y permitido 200 se espera error": {"/file", http.MethodPost, []string{"gatito.jpeg"}, ".jpg", 1, 200, "error"},
	}
)

func Test_ServeHTTP(t *testing.T) {
	//TEST.....
	DeleteUploadTestFiles()

	err := file.CreateFolderIfNotExist(root_test_folder)
	if err != nil {
		fmt.Println("Error:", err)
	}

	const field_name = "endoscopia"
	const module_name = "medical_history"

	for prueba, data := range dataHttp {
		t.Run((prueba), func(t *testing.T) {

			mux := http.NewServeMux()

			h := file.New(root_test_folder, data.extensions, data.max_files, data.max_size)

			mux.HandleFunc(data.endpoint, h.ServeHTTP)

			server := httptest.NewServer(mux)
			defer server.Close()

			c := http.DefaultClient

			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)

			for _, file_name := range data.files {

				// abrimos el archivo local para la prueba
				File, err := os.Open(path_files + file_name)
				if err != nil {
					log.Fatal(err)
				}
				defer File.Close()

				// creamos formato de envió de archivo
				part, err := writer.CreateFormFile(field_name, file_name)
				if err != nil {
					log.Fatalln(err)
				}
				_, err = io.Copy(part, File)
				if err != nil {
					log.Fatal(err)
				}

				// escribimos en memoria el campo del formulario
				err = writer.WriteField(field_name, file_name)
				if err != nil {
					log.Fatal(err)
				}
			}

			// Agregamos los parámetros al formulario

			for key, value := range map[string]string{
				"module_name": module_name,
				"field_name":  field_name,
				"folder_id":   randomNumber(),
			} {
				err = writer.WriteField(key, value)
				if err != nil {
					log.Fatal(err)
				}
			}

			err := writer.Close()
			if err != nil {
				log.Fatal(err)
			}

			// enviamos post con el contenido formulario y cuerpo solicitud
			req, err := http.NewRequest(data.method, server.URL+data.endpoint, body)
			if err != nil {
				log.Fatalf("error %s", err)
			}
			req.Header.Add("Content-Type", writer.FormDataContentType())

			res, err := c.Do(req)
			if err != nil {
				log.Fatalf("error %s", err)
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
			// fmt.Println("*** RESPUESTA SOLICITUD CREATE: ", response)

			if response.Type != data.expected {
				log.Fatal(response)
			}

			if response.Type != "error" {

				t.Run("UPDATE Test:", func(t *testing.T) {
					updateTest(data.endpoint, server, response)
				})
				t.Run("READ Test:", func(t *testing.T) {
					readTest(data.endpoint, server, response)
				})

				t.Run("DELETE Test:", func(t *testing.T) {
					deleteTest(data.endpoint, server, response)
				})

			}

		})
	}

}

// https://matt.aimonetti.net/posts/2013-07-golang-multipart-File-upload-example/

type Response struct {
	Type    string
	Data    []map[string]string
	Object  string
	Module  string
	Message string
}
