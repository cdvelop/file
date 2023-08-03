package file

func (f File) Update(data ...map[string]string) ([]map[string]string, error) {

	return f.db.UpdateObjectsInDB(f.Name, data...)
}
