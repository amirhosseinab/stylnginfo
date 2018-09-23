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

    AnalyzedData struct {
        HTMLFiles []*HTMLFile `json:"htmlFiles"`
    }

    HTMLFile struct {
        Name     string     `json:"name"`
        Content  string     `json:"-"`
        CSSFiles []*CSSFile `json:"cssFiles"`
    }

    CSSFile struct {
        Name      string      `json:"name"`
        Content   string      `json:"-"`
        Selectors []*Selector `json:"selectors"`
    }

    Selector struct {
        Name     string `json:"name"`
        Selected bool   `json:"selected"`
    }

    File struct {
        Name     string   `json:"name"`
        Content  string   `json:"-"`
        FileType FileType `json:"fileType"`
    }
)

func NewFile(name, fileType, content string) *File {
    types := map[string]FileType{
        "text/html": FileTypeHTML,
        "text/css":  FileTypeCSS,
    }
    return &File{Name: name, Content: content, FileType: types[fileType]}
}

func AnalyzeFiles(files ...*File) *AnalyzedData {
    result := &AnalyzedData{}
    var cssFiles []*CSSFile

    for _, f := range files {
        if f.FileType == FileTypeCSS {
            cf := &CSSFile{Name: f.Name, Content: f.Content, Selectors: extractSelectors(f)}
            cssFiles = append(cssFiles, cf)
        }
    }

    for _, f := range files {
        if f.FileType == FileTypeHTML {

            hf := &HTMLFile{Name: f.Name, Content: f.Content}

            doc, err := goquery.NewDocumentFromReader(strings.NewReader(hf.Content))
            if err != nil {
                log.Fatal(err)
            }

            for _, cf := range cssFiles {

                var ts map[string]bool
                ts = make(map[string]bool)
                cssFile := &CSSFile{Name: cf.Name, Selectors: []*Selector{}}

                for _, s := range cf.Selectors {
                    doc.Find(s.Name).Each(func(_ int, selection *goquery.Selection) {
                        if _, ok := ts[s.Name]; !ok {
                            cssFile.Selectors = append(cssFile.Selectors, s)
                            ts[s.Name] = true
                        }
                    })
                }

                if len(cssFile.Selectors) > 0 {
                    hf.CSSFiles = append(hf.CSSFiles, cssFile)
                }
            }
            result.HTMLFiles = append(result.HTMLFiles, hf)
        }
    }
    return result
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
                result = append(result, &Selector{Name: s, Selected: false})
            }
        }
    }
    return result
}
