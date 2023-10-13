package file

import "github.com/cdvelop/model"

func (f File) Update(u *model.User, data ...map[string]string) ([]map[string]string, error) {

	return f.db.UpdateObjectsInDB(f.Name, data...)
}
