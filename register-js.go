package file

import (
	"fmt"
)

func (File) SelectedTargetChanges() string {
	return "LoadNewPictures(input);"
}

func (File) InputValueChanges() string {
	return "UploadNewFiles(input);"
}

func (File) FieldAddEventListener(field_name string) string {
	return fmt.Sprintf(`input_%v.addEventListener("change", UploadNewFiles);`, field_name)
}
