package file

import (
	"github.com/gabriel-vasile/mimetype"
)

func ArchiveType(File *[]byte) string {
	if File != nil {
		mtype := mimetype.Detect(*File)
		// fmt.Println("ARCHIVO: ", mtype.String(), mtype.Extension())

		return mtype.Extension()
	}
	return ""
}
