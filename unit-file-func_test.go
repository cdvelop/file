package file_test

import (
	"log"
	"os"
	"testing"

	"github.com/cdvelop/file"
)

const path_files = "./files_test/"

func Test_FileType(t *testing.T) {

	files, err := os.ReadDir(path_files)
	if err != nil {
		log.Fatalln(err)
	}

	for _, archive := range files {
		// fmt.Println(archive.MainName(), archive.IsDir())

		// Open File
		f, err := os.Open(path_files + archive.Name())
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()

	}
}

var (
	dataFiles = map[string]struct {
		name_file string
		expected  string
	}{
		"archivo excel ok": {"document-excel.xlsx", ".xlsx"},
		"archivo texto ok": {"document-text.txt", ".txt"},
		"archivo word ok":  {"document-word.docx", ".docx"},
		"archivo pdf ok":   {"document.pdf", ".pdf"},

		"archivo rar ok": {"File.rar", ".rar"},
		"archivo zip ok": {"File.zip", ".zip"},

		"jpep ok": {"gatito.jpeg", ".jpg"},
		"png ok":  {"dino.png", ".png"},

		"video mp4 ok": {"video.mp4", ".mp4"},
	}
)

func Test_GetFileType(t *testing.T) {
	for prueba, data := range dataFiles {
		t.Run((prueba), func(t *testing.T) {

			full_path := path_files + data.name_file

			archive, err := os.ReadFile(full_path)
			if err != nil {
				log.Println(err)
				t.Fail()
				return
			}

			extension := file.ArchiveType(&archive)

			if extension != data.expected {
				t.Fatalf("err respuesta fue: [%v] se esperaba [%v]", extension, data.expected)
				t.Fail()
			}
		})
	}
}
