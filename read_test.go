package file_test

import (
	"log"
	"testing"

	"github.com/cdvelop/model"
	"github.com/cdvelop/testools"
)

func (d *dataTest) readFileTest(in model.Response, t *testing.T) {
	t.Run("READ FILE:", func(t *testing.T) {

		// fmt.Println("DATA A ENVIAR PARA LECTURA: ", in.Data)
		for _, data := range in.Data {
			d.Endpoint = "/file/" + d.file.Object.Name

			new_data := map[string]string{"id_file": data["id_file"]}

			responses, code, err := d.Get(new_data)

			if err != nil {
				log.Fatal(err)
			}

			for _, resp := range responses {
				testools.CheckTest("read file", 200, code, resp)
			}
		}
	})
}

func (d *dataTest) readTest(in model.Response, t *testing.T) {

	var folders_ids []map[string]string

	t.Run("READ ONE DATA JSON:", func(t *testing.T) {

		// fmt.Println("DATA A ENVIAR PARA LECTURA JSON: ", in.Data)

		for _, data := range in.Data {
			d.Endpoint = "/read/" + d.file.Object.Name

			new_data := map[string]string{"id_file": data["id_file"]}

			responses, code, err := d.Get(new_data)

			if err != nil {
				log.Fatal(err)
			}

			// fmt.Println("RESPUESTA LECTURA JSON: ", responses)

			for i, resp := range responses {

				if id, folder_id := resp.Data[i]["folder_id"]; !folder_id {
					log.Fatalln("se esperaba recuperar folder id")
				} else {
					folders_ids = append(folders_ids, map[string]string{"folder_id": id})
				}

				if path, file_path := resp.Data[i]["file_path"]; file_path {
					log.Fatalln("error de seguridad servidor de archivos. no se espera recibir la ruta del archivo como respuesta de lectura al cliente ej:", path)
				}

				testools.CheckTest("read data json", 200, code, resp)
			}
		}
	})

	t.Run("READ ALL BY FOLDER ID DATA JSON:", func(t *testing.T) {

		// fmt.Println("DATA FOLDER ID PARA LECTURA JSON: ", folders_ids)

		for _, new_data := range folders_ids {
			d.Endpoint = "/read/" + d.file.Object.Name

			responses, code, err := d.Get(new_data)

			if err != nil {
				log.Fatal(err)
			}

			if len(responses) != 1 {
				log.Fatal("error se esperaba 1 respuesta se obtuvo: ", len(responses))
			}

			// fmt.Println("RESPUESTAS LECTURA FOLDER ID JSON: ", len(responses))

			for i, resp := range responses {

				// if _, folder_id := resp.Data[i]["folder_id"]; folder_id {
				// 	log.Fatalln("error no se espera recibir nuevamente el dato folder_id")
				// }

				if path, file_path := resp.Data[i]["file_path"]; file_path {
					log.Fatalln("error de seguridad no se espera recibir la ruta del archivo en el servidor ", path)
				}

				testools.CheckTest("read data json", 200, code, resp)
			}
		}
	})

}
