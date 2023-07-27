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

	"github.com/cdvelop/api"
	"github.com/cdvelop/cutkey"
	"github.com/cdvelop/file"
	"github.com/cdvelop/gotools"
	"github.com/cdvelop/model"
	"github.com/cdvelop/sqlite"
)

var (
	dataHttp = map[string]struct {
		field_name string   //ej: endoscopia, voucher, foto_mascota, foto_usuario
		endpoint   string   //ej: /create/
		method     string   //ej: "POST","GET"
		files      []string //ej: "gatito.jpg, perro.png"
		file_type  string   //ej: imagen,video,document,pdf
		max_files  string
		max_size   string
		expected   string
	}{
		"crear 2 archivos gatito 220kb y dino 36kb": {"endoscopia", "/create/", http.MethodPost, []string{"dino.png", "gatito.jpeg"}, "imagen", "2", "262", "create"},
		// "gatito 220kb ok":                                     {"foto_mascota","file/upload", http.MethodPost, []string{"gatito.jpeg"}, "imagen", "1", "220", "create"},
		// "tamaño gatito 220kb y permitido 200 se espera error": {"foto_mascota","file/upload", http.MethodPost, []string{"gatito.jpeg"}, "imagen", "1", "200", "error"},
	}
)

func Test_ServeHTTP(t *testing.T) {
	//TEST.....
	DeleteUploadTestFiles()

	err := gotools.CreateFolderIfNotExist(root_test_folder)
	if err != nil {
		fmt.Println("Error:", err)
	}

	for prueba, data := range dataHttp {
		t.Run((prueba), func(t *testing.T) {

			db := sqlite.NewConnection(root_test_folder, "stored_files_index.db", false)

			module := model.Module{
				Name:    "medical_history",
				Title:   "Modulo Testing",
				Areas:   []byte{},
				Objects: []*model.Object{},
			}

			h := file.New(&module, db, "root_folder:"+root_test_folder, "name:"+data.field_name, "max_files:"+data.max_files, "max_kb_size:"+data.max_size)
			// h.AddModule(&module)

			fmt.Println("NOMBRE DEL API: ", h.Object.Api())

			cut := cutkey.Add(&h.Object)

			api_conf := api.Add([]*model.Module{&module})

			mux := api_conf.ServeMuxAndRoutes()

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
				part, err := writer.CreateFormFile(h.Fields[6].Name, file_name)
				if err != nil {
					log.Fatalln(err)
				}
				_, err = io.Copy(part, File)
				if err != nil {
					log.Fatal(err)
				}

				// escribimos en memoria el campo del formulario
				err = writer.WriteField(h.Fields[6].Name, file_name)
				if err != nil {
					log.Fatal(err)
				}
			}

			// Agregamos los parámetros al formulario
			for key, value := range map[string]string{
				h.Fields[1].Name: module.Name,
				h.Fields[2].Name: data.field_name,
				h.Fields[3].Name: randomNumber(),
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
			req, err := http.NewRequest(data.method, server.URL+data.endpoint+h.Api(), body)
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

			for _, response := range cut.DecodeResponses(resp) {

				// fmt.Println("*** RESPUESTA SOLICITUD CREATE: ", response)

				if response.Action != data.expected {
					log.Fatal(response)
				}

				// if response.Action != "error" {

				// t.Run("UPDATE Test:", func(t *testing.T) {
				// 	updateTest(data.file_name, server, response)
				// })
				// t.Run("READ Test:", func(t *testing.T) {
				// 	readTest(data.file_name, server, response)
				// })

				// t.Run("DELETE Test:", func(t *testing.T) {
				// 	deleteTest(data.file_name, server, response)
				// })

				// }

			}

		})
	}

}

// https://matt.aimonetti.net/posts/2013-07-golang-multipart-File-upload-example/

type Response struct {
	Action  string
	Data    []map[string]string
	Object  string
	Module  string
	Message string
}
