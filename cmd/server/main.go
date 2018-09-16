package main

import (
    "github.com/gorilla/mux"
    "html/template"
    "net/http"
    "os"
    "path"
)

func main() {
    port := os.Getenv("PORT")
    r := mux.NewRouter()

    r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist", http.FileServer(http.Dir("dist"))))
    r.HandleFunc("/", IndexHandler)
    r.NotFoundHandler = http.HandlerFunc(IndexHandler)

    http.ListenAndServe(":"+port, r)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    wd, _ := os.Getwd()
    t := template.Must(template.New("index.html").ParseFiles(path.Join(wd, "dist/index.html")))
    t.Execute(w, nil)
}
