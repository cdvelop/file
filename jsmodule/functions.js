let fileInput = module.querySelector('input[type=File]')


fileInput.addEventListener('change', function (e) {

    uploadFile(e);
});


// console.log("FILE INPUT SELECCIONADO ", fileInput);

function uploadFile() {

    SOCK.binaryType = 'blob';
    // SOCK.binaryType = 'arraybuffer';

    var files = fileInput.files;
    
    for (const File of files) {
        console.log("ARCHIVO A ENVIAR: ", File);
        SOCK.send(File)
    }


};

function LoadNewPictures(input) {

    console.log("BORRANDO IMAGENES ANTERIORES")

    console.log("CARGANDO NUEVAS IMAGENES INPUT: ",input)
}


function UploadNewFiles(input) {
    
    console.log("SUBIR ARCHIVOS INPUT: ",input)
}