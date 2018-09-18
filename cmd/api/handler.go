package api

import "net/http"

func UploadFilesHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseMultipartForm(32 << 20)
    files := r.MultipartForm.File["files"]
    for _, file := range files {
        //f, err := file.Open()
        //if err != nil {
        //    log.Fatal(err)
        //}

        switch file.Header.Get("Content-Type"){
        case "text/css":
        case "text/html":
        }

        //defer f.Close()

        //b, err := ioutil.ReadAll(f)
        //if err != nil {
        //    log.Fatal(err)
        //}

    }

}