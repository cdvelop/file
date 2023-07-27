package file_test

import (
	"fmt"

	"github.com/cdvelop/gotools"
)

const root_test_folder = "./root_test_folder"

func DeleteUploadTestFiles() {

	//calcular tamaño máximo carpeta antes de borrar

	err := gotools.DeleteIfFolderSizeExceeds(root_test_folder, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}
	//  else {
	// 	fmt.Println("El contenido de la carpeta ", root_test_folder, " fue eliminado su tamaño era superior a 100 MB.")
	// }

}
