package app

import (
    "github.com/PuerkitoBio/goquery"
    "log"
    "regexp"
    "strings"
)

const (
    FileTypeHTML = iota
    FileTypeCSS
)

type (
    FileType int

    Analyzer struct {
        Files    []*File     `json:"files"`
        Selector []*Selector `json:"selectors,omitempty"`
    }

    File struct {
        Name     string   `json:"name"`
        Content  string   `json:"content"`
        FileType FileType `json:"fileType"`
    }

    Selector struct {
        Name     string `json:"name"`
        FileName string `json:"fileName"`
    }
)

func NewAnalyzer(files ...*File) *Analyzer {
    return &Analyzer{Files: files}
}

func NewFile(name, fileType, content string) *File {
    types := map[string]FileType{
        "text/html": FileTypeHTML,
        "text/css":  FileTypeCSS,
    }
    return &File{Name: name, Content: content, FileType: types[fileType]}
}

func (a *Analyzer) Analyze() {
    for _, f := range a.Files {
        if f.FileType == FileTypeCSS {
            a.Selector = append(a.Selector, extractSelectors(f)...)
        }
    }
    for _, f := range a.Files {
        if f.FileType == FileTypeHTML {
            doc, err := goquery.NewDocumentFromReader(strings.NewReader(f.Content))
            if err != nil {
                log.Fatal(err)
            }
            for _, s := range a.Selector {
                doc.Find(s.Name).Each(func(_ int, selection *goquery.Selection) {
                    log.Println(selection.Attr("class"))
                })
            }
        }
    }
}

func extractSelectors(file *File) []*Selector {
    var result []*Selector
    var selectors map[string]interface{}
    selectors = make(map[string]interface{})

    r, err := regexp.Compile("\\.[\\w\\-0-9\\s\\.:,\\*\\>\\(\\)]+{")
    if err != nil {
        log.Fatal(err)
    }
    items := r.FindAllString(file.Content, -1)
    for _, item := range items {
        refinedItem := item
        refinedItem = strings.Trim(refinedItem, "{")

        subItem := strings.Split(refinedItem, ",")
        for _, s := range subItem {
            s = strings.TrimSpace(s)
            if _, ok := selectors[s]; !ok {
                selectors[s] = nil
                result = append(result, &Selector{
                    Name:     s,
                    FileName: file.Name,
                })
            }
        }
    }
    return result
}
