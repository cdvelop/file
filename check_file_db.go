package file

import (
	"fmt"
	"log"

	"github.com/cdvelop/dbtools"
	"github.com/cdvelop/objectdb"
	"github.com/cdvelop/sqlite"
)

func (f *File) checkDataBase() {

	dba := sqlite.NewConnection(root_folder, "stored_files_index.db", false)

	db := objectdb.Get(dba)

	f.Connection = db

	if !dba.TableExist(f.Object().Name, f.Connection) {
		db.Set(dba)
		// defer db.Close()
		fmt.Println("NO EXISTE TABLA: ", f.Object().Name)

		m := dbtools.NewOperationDB(db.DB, dba)

		if !m.CreateAllTABLES(f.Object()) {
			log.Fatalf("Error No se logro crear tabla: %v", f.Object().Name)
		}
	}

}
