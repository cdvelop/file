package file

import (
	"fmt"

	"github.com/cdvelop/model"
)

func (f File) FilePath(u *model.User, params map[string]string) (string, error) {
	// fmt.Println("parámetros FilePath recibidos: ", params)

	data, err := f.db.ReadObjectsInDB(f.Name, params)
	if err != nil {
		return "", err
	}

	if len(data) != 1 {
		return "", fmt.Errorf("parámetros incorrectos al recuperar archivo")
	}

	return data[0][f.FieldFilePath], nil
}

func (f File) Read(u *model.User, params ...map[string]string) ([]map[string]string, error) {

	// fmt.Println("parámetros leer recibidos:", params)

	for _, v := range params {
		v["SELECT"] = f.FieldModuleName + "," + f.FieldName + "," + f.FieldFolderId + "," + f.FieldDescription
	}

	data, err := f.db.ReadObjectsInDB(f.Name, params...)
	if err != nil {
		return nil, err
	}

	return data, nil
}
