package file

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const error_message_big_file = "error archivo muy grande. tamaño máximo admitido: %v kb"

// Create upload files http handler
func (f File) Create(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("TAMAÑO CONTENIDO: ", r.ContentLength)
	// for i, v := range r.Header {
	// 	fmt.Println(i, v)
	// }
	max_size := int64(float64(f.max_files*f.max_kb_size*1024) * 1.05)
	// fmt.Println("TAMAÑO ACEPTADO: ", max_size)

	r.Body = http.MaxBytesReader(w, r.Body, max_size) // 220 KB

	err := r.ParseMultipartForm(max_size)
	if err != nil {
		if strings.Contains(err.Error(), "multipart") {
			f.error(w, err.Error(), http.StatusNotAcceptable)
		} else {
			f.error(w, fmt.Sprintf(error_message_big_file, f.max_kb_size), http.StatusNotAcceptable)
		}
		// log.Println("Error MultipartForm: ", err)
		return
	}

	var new_data = make(map[string]string)

	for i, field := range f.Object().Fields {
		if i > 0 && i <= 4 { // saltarse id y file_path
			new_data[field.Name] = r.FormValue(field.Name)
		}
	}

	// fmt.Println("FORM VALUES: ", new_data)

	// ej ./app_files/medicalhistory/endoscopia/123344
	upload_folder := f.root_folder + "/" + new_data["module_name"] + "/" + new_data["field_name"] + "/" + new_data["folder_id"]

	files := r.MultipartForm.File[new_data["field_name"]]
	if len(files) == 0 {
		f.error(w, "no hay archivos detectados", http.StatusNotAcceptable)
		return
	}

	// fmt.Println("ARCHIVOS: ", len(files))

	if len(files) > int(f.max_files) {
		f.error(w, fmt.Sprintf("error se pretende subir %v archivos, pero el máximo permitido es: %v", len(files), f.max_files), http.StatusNotAcceptable)
		return
	}

	data_out := []map[string]string{}

	for _, fileHeader := range files {

		// fmt.Println("FILE NAME: ", fileHeader.Filename)
		// fmt.Println("FILE SIZE: ", fileHeader.Size)
		// fmt.Println("FILE TYPE: ", fileHeader.Header.Get("Content-Type"))

		if fileHeader.Size > max_size {
			f.error(w, fmt.Sprintf(error_message_big_file, f.max_kb_size), http.StatusNotAcceptable)
			return
		}

		File, err := fileHeader.Open()
		if err != nil {
			f.error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		defer File.Close()

		buff := make([]byte, 512)
		_, err = File.Read(buff)
		if err != nil {
			f.error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		extension := ArchiveType(&buff)

		// fmt.Println("EXTENSION: ", extension, "LEN: ", len(extension))

		if !strings.Contains(f.extensions, extension) {
			f.error(w, "formato de archivo no valido solo se admiten: "+f.extensions, http.StatusBadRequest)
			return
		}
		extension = filepath.Ext(fileHeader.Filename)

		new_file_name := getNewID()
		new_data["id_file"] = new_file_name
		new_data["file_path"] = upload_folder + "/" + new_file_name + extension

		if len(fileHeader.Filename) > 5 {
			// agregamos como descripción el nombre que trae el archivo quitando la extension
			new_data["description"] = fileHeader.Filename[:len(fileHeader.Filename)-len(extension)]
		}

		//validar data
		if mg, ok := f.Object().ValidateData(true, new_data); !ok {
			f.error(w, mg, http.StatusNotAcceptable)
			return
		}

		// volver lectura de archivo desde el principio
		_, err = File.Seek(0, io.SeekStart)
		if err != nil {
			f.error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = os.MkdirAll(upload_folder, os.ModePerm)
		if err != nil {
			f.error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		dst, err := os.Create(fmt.Sprintf("%v/%v%s", upload_folder, new_file_name, extension))
		if err != nil {
			f.error(w, err.Error(), http.StatusBadRequest)
			return
		}

		defer dst.Close()

		// Copy the uploaded File to the filesystem at the specified destination
		_, err = io.Copy(dst, File)
		if err != nil {
			f.error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// guardar en base de datos
		mg, ok := f.CreateObjects(f.Object().Name, new_data)

		if !ok {
			f.error(w, mg, http.StatusInternalServerError)
			return
		}

		out := map[string]string{
			"id_file":     new_file_name,
			"description": new_data["description"],
		}

		data_out = append(data_out, out)

	}

	f.response(w, http.StatusOK, "create", "Carga exitosa", new_data["module_name"], data_out...)

}
