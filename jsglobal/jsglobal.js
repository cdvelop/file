function NewFrontRequest(props) {
    let { Method, ContentType ,EndPoint, Data, Response } = props;

    const content_options = {
        'JSON': 'application/json',
        'FILE': 'multipart/form-data',
        'FORM': 'application/x-www-form-urlencoded'
    }

    const content_type = content_options[ContentType] || 'application/json';
      
    
    fetch(URL_HTTPS + EndPoint, {
        method: Method,
        headers: {
            "Content-Action": content_type,
        },
        body: JSON.stringify(Data)
    })
        // si respuesta es ok pasamos a json de lo contrario rechazamos
        .then(res => res.ok ? res.json() : Promise.reject(res))
        // caso de éxito enviar a la función de respuesta
        .then(json => Response(json))
        .catch(err => {
            const error_msg = "OCURRIÓ UN ERROR AL ACCEDER A LA API " + EndPoint;
            let message = err.statusText || error_msg;
            ShowMessageToUser("err", message)
        });

}