package app

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "mime/multipart"
    "net/http"
)

func UploadFilesHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseMultipartForm(32 << 20)
    files := r.MultipartForm.File["files"]
    var fs []*File

    for _, file := range files {
        fs = append(fs, NewFile(GetFileInfo(file)))
    }

    result := AnalyzeFiles(fs...)

    b, err := json.Marshal(result)
    if err != nil {
        log.Fatal(err)
    }

    w.WriteHeader(http.StatusOK)
    w.Write(b)
}

func GetFileInfo(file *multipart.FileHeader) (fileName, fileType, content string) {
    f, err := file.Open()
    if err != nil {
        log.Fatal(err)
    }

    fileName = file.Filename
    fileType = file.Header.Get("Content-Type")
    b, err := ioutil.ReadAll(f)
    if err != nil {
        log.Fatal(err)
    }
    return fileName, fileType, string(b)
}

func NewFile(name, fileType, content string) *File {
   types := map[string]FileType{
       "text/html": FileTypeHTML,
       "text/css":  FileTypeCSS,
   }
   return &File{Name: name, Content: content, FileType: types[fileType]}
}
