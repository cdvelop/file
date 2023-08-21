package file_test

import (
	"log"
	"testing"

	"github.com/cdvelop/model"
	"github.com/cdvelop/testools"
)

func (d *dataTest) create(prueba string, t *testing.T) (responses []model.Response) {

	t.Run((prueba), func(t *testing.T) {
		var err error

		form := map[string]string{
			d.Fields[1].Name: d.ModuleName,
			d.Fields[2].Name: d.field_name,
			d.Fields[3].Name: testools.RandomNumber(),
		}

		body, content_type, err := testools.MultiPartFileForm(path_files, d.Fields[6].Name, d.files, form)
		if err != nil {
			log.Fatal(err)
		}
		d.ContentType = content_type

		// fmt.Println("METHOD: ", d.Method)

		// var code int
		d.Endpoint += "/" + d.ID()

		// fmt.Println("ENDPOINT CREATE: ", d.Endpoint)

		responses, _, err = d.SendRequest(body.Bytes())
		if err != nil {
			log.Fatal(err)
		}

		for _, resp := range responses {
			testools.CheckTest(prueba, d.expected, resp.Action, resp)
		}

	})

	return

}
