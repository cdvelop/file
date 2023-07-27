package file

import (
	"fmt"
)

func (config) SelectedTargetChanges() string {
	return "LoadNewPictures(input);"
}

func (config) InputValueChanges() string {
	return "UploadNewFiles(input);"
}

func (config) FieldAddEventListener(field_name string) string {
	return fmt.Sprintf(`input_%v.addEventListener("change", UploadNewFiles);`, field_name)
}
