package file

import (
	_ "embed"
	"fmt"
)

//go:embed jsmodule/functions.js
var functions string

func (File) AttachJsFunctions() string {
	return functions
}

func (File) SelectedTargetChanges() string {
	return "LoadNewPictures(input);"
}

func (File) InputValueChanges() string {
	return "UploadNewFiles(input);"
}

func (File) FieldAddEventListener(field_name string) string {
	return fmt.Sprintf(`input_%v.addEventListener("change", UploadNewFiles);`, field_name)
}

func (File) FieldRemoveEventListener(field_name string) string {
	return fmt.Sprintf(`input_%v.removeEventListener("change", UploadNewFiles);`, field_name)
}
