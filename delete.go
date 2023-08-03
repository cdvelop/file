package file

import (
	"fmt"
	"os"
)

func (f File) Delete(params ...map[string]string) ([]map[string]string, error) {

	// fmt.Println("parámetros Delete recibidos:", params)

	recover_data, err := f.db.DeleteObjectsInDB(f.Name, params...)
	if err != nil {
		return nil, err
	}

	// fmt.Println("DATA RECOBRADA DESPUÉS DE BORRAR: ", recover_data)

	for _, data := range recover_data {

		// Borrar archivos desde hdd
		err := os.Remove(data[f.FieldFilePath])
		if err != nil {
			return nil, fmt.Errorf("archivo %s fue eliminado de la db pero no del hdd %s", data[f.FieldName], err)
		}
	}

	return recover_data, nil
}
