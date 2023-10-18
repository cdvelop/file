package file

import (
	"fmt"
	"os"

	"github.com/cdvelop/model"
)

func (f File) Delete(u *model.User, params ...map[string]string) ([]map[string]string, error) {

	// fmt.Println("parámetros Delete recibidos:", params)

	recover_data, err := f.db.DeleteObjectsInDB(f.Object.Table, params...)
	if err != nil {
		return nil, err
	}

	// fmt.Println("DATA RECOBRADA DESPUÉS DE BORRAR: ", recover_data)

	for _, data := range recover_data {

		// Borrar archivos desde hdd
		err := os.Remove(data[f.File_path])
		if err != nil {
			return nil, fmt.Errorf("archivo %s fue eliminado de la db pero no del hdd %s", data[f.Field_name], err)
		}
	}

	return recover_data, nil
}
