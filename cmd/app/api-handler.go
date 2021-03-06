package app

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "mime/multipart"
    "net/http"
    "strings"
)

type (
    cssFilesWeightData struct {
        FileName      string `json:"fileName"`
        SelectorCount int    `json:"selectorCount"`
        IsUsed        bool   `json:"isUsed"`
    }
)

func CSSFileHandler(w http.ResponseWriter, r *http.Request) {
    cssFiles := ScrutinizeCSSFiles(readIndexFile(r), readFiles(r))
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

func LessFileHandler(w http.ResponseWriter, r *http.Request) {
    files := readFiles(r)
    d := ScrutinizeLessFiles(files)

    b, err := json.Marshal(d)
    if err != nil {
        log.Fatal(err)
    }

    w.WriteHeader(http.StatusOK)
    w.Write(b)
}

func readIndexFile(r *http.Request) *File {
    var indexFile *File
    _, fh, _ := r.FormFile("indexFile")
    if fh != nil {
        indexFile = NewFile(getFileInfo(fh))
    }
    return indexFile
}

func readFiles(r *http.Request) []*File {
    var files []*File

    r.ParseMultipartForm(8 << 20)
    fs := r.MultipartForm.File["files"]
    for _, f := range fs {
        files = append(files, NewFile(getFileInfo(f)))
    }
    return files
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
    var ft FileType
    types := map[string]FileType{
        "text/html": FileTypeHTML,
        "text/css":  FileTypeCSS,
    }

    ft, ok := types[fileType]
    if !ok && strings.HasSuffix(name, ".less") {
        ft = FileTypeLess
    }
    if !ok && strings.HasSuffix(name, ".scss") {
        ft = FileTypeScss
    }

    return &File{Name: name, Content: content, FileType: ft}
}
