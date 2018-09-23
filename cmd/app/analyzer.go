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
        CSSFiles  []*CSSFile  `json:"cssFiles"`
    }

    HTMLFile struct {
        Name             string      `json:"name"`
        Content          string      `json:"-"`
        AppliedSelectors []*Selector `json:"appliedSelectors"`
    }

    CSSFile struct {
        Name      string   `json:"name"`
        Content   string   `json:"-"`
        Selectors []string `json:"selectors"`
    }

    Selector struct {
        SelectorName string `json:"selectorName"`
        CSSFileName  string `json:"cssFileName"`
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

    for _, f := range files {
        if f.FileType == FileTypeCSS {
            cf := &CSSFile{Name: f.Name, Content: f.Content, Selectors: extractSelectors(f)}
            result.CSSFiles = append(result.CSSFiles, cf)
        }
    }

    for _, f := range files {
        if f.FileType == FileTypeHTML {
            hf := &HTMLFile{Name: f.Name, Content: f.Content}

            doc, err := goquery.NewDocumentFromReader(strings.NewReader(hf.Content))
            if err != nil {
                log.Fatal(err)
            }

            var ts map[string]bool
            ts = make(map[string]bool)

            for _, cf := range result.CSSFiles {
                for _, s := range cf.Selectors {
                    doc.Find(s).Each(func(_ int, selection *goquery.Selection) {
                        if _, ok := ts[cf.Name+s]; !ok {
                            selector := &Selector{SelectorName: s, CSSFileName: cf.Name}
                            hf.AppliedSelectors = append(hf.AppliedSelectors, selector)
                            ts[cf.Name+s] = true
                        }
                    })
                }
            }

            result.HTMLFiles = append(result.HTMLFiles, hf)
        }
    }
    return result
}

func extractSelectors(file *File) []string {
    var result []string
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
                result = append(result, s)
            }
        }
    }
    return result
}
