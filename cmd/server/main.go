package main

import (
    "github.com/amirhosseinab/stylnginfo/cmd/app"
    "github.com/gorilla/mux"
    "html/template"
    "log"
    "net/http"
    "os"
    "path"
    "time"
)

func main() {
    port := os.Getenv("PORT")
    r := mux.NewRouter()

    r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist", http.FileServer(http.Dir("dist"))))
    r.HandleFunc("/", IndexHandler)

    a := r.PathPrefix("/api").Subrouter()
    a.HandleFunc("/analyze", app.UploadFilesHandler).Methods("POST")

    r.NotFoundHandler = http.HandlerFunc(IndexHandler)

    srv := &http.Server{
        Handler:      r,
        Addr:         ":" + port,
        WriteTimeout: 60 * time.Second,
        ReadTimeout:  60 * time.Second,
    }

    log.Fatal(srv.ListenAndServe())

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    wd, _ := os.Getwd()
    t := template.Must(template.New("index.html").ParseFiles(path.Join(wd, "dist/index.html")))
    t.Execute(w, nil)
}
