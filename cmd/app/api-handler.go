package app

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "mime/multipart"
    "net/http"
)

type (
    cssFilesWeightData struct {
        FileName      string `json:"fileName"`
        SelectorCount int    `json:"selectorCount"`
        IsUsed        bool   `json:"isUsed"`
    }
)

func CSSFileHandler(w http.ResponseWriter, r *http.Request) {
    cssFiles := ScrutinizeCSSFiles(readFiles(r))
    var d []*cssFilesWeightData

    for _, f := range cssFiles {
        d = append(d, &cssFilesWeightData{
            FileName:      f.Name,
            SelectorCount: len(f.Selectors),
            IsUsed:        f.IsUsed,
        })
    }

    b, err := json.Marshal(d)
    if err != nil {
        log.Fatal(err)
    }

    w.WriteHeader(http.StatusOK)
    w.Write(b)
}

func readFiles(r *http.Request) (indexFile *File, files []*File) {
    r.ParseMultipartForm(8 << 20)
    _, fh, _ := r.FormFile("indexFile")
    indexFile = NewFile(getFileInfo(fh))

    fs := r.MultipartForm.File["files"]
    for _, f := range fs {
        files = append(files, NewFile(getFileInfo(f)))
    }
    return
}

//func UploadFilesHandler(w http.ResponseWriter, r *http.Request) {
//    fs := readFiles(r)
//    result := AnalyzeFiles(fs...)
//
//    b, err := json.Marshal(result)
//    if err != nil {
//        log.Fatal(err)
//    }
//
//    w.WriteHeader(http.StatusOK)
//    w.Write(b)
//}

func getFileInfo(file *multipart.FileHeader) (fileName, fileType, content string) {
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
