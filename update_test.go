package file_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/cdvelop/model"
	"github.com/cdvelop/testools"
)

func (d *dataTest) updateTest(r model.Response, t *testing.T) {
	t.Run("UPDATE Test:", func(t *testing.T) {

		// fmt.Println("DATA A ACTUALIZAR: ", rq)
		for _, data := range r.Data {
			data[d.file.FieldDescription] = "perro"
		}

		d.Endpoint = "/update/" + d.Api()

		d.Data = r.Data

		responses, _, err := d.CutPost()
		if err != nil {
			log.Fatal(err)
		}

		for i, resp := range responses {
			testools.CheckTest("update", "update", resp.Action, resp)

			if resp.Data[i][d.file.FieldDescription] != "perro" {
				fmt.Printf("se esperaba en description [perro] se obtuvo: [%v]", resp.Data[i][d.file.FieldDescription])
				log.Fatal()
			}
		}
	})
}
