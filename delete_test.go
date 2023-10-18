package file_test

import (
	"log"
	"testing"

	"github.com/cdvelop/model"
	"github.com/cdvelop/testools"
)

func (d dataTest) deleteTest(r model.Response, t *testing.T) {
	t.Run("DELETE Test:", func(t *testing.T) {
		// Crear una solicitud DELETE

		d.Endpoint = "/delete/" + d.file.Object.Name

		d.Data = r.Data

		responses, _, err := d.CutPost()
		if err != nil {
			log.Fatal(err)
		}

		for _, resp := range responses {
			testools.CheckTest("update", "delete", resp.Action, resp)
		}
	})
}
