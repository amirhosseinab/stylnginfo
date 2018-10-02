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
        *File
        CSSFiles []*CSSFile `json:"cssFiles,omitempty"`
    }

    CSSFile struct {
        *File
        Selectors []*Selector `json:"selectors,omitempty"`
    }

    File struct {
        Name     string   `json:"name"`
        Content  string   `json:"-"`
        FileType FileType `json:"fileType"`
    }

    Selector struct {
        Name         string         `json:"name"`
        Declarations []*Declaration `json:"declarations,omitempty"`
        MediaQuery   string         `json:"mediaQuery,omitempty"`
    }

    Declaration struct {
        Property string `json:"property"`
        Value    string `json:"value"`
    }
)

func AnalyzeFiles(files ...*File) *AnalyzedData {
    ad := &AnalyzedData{}
    for _, f := range files {
        if f.FileType == FileTypeHTML {
            ad.HTMLFiles = append(ad.HTMLFiles, &HTMLFile{File: f})
        }
        if f.FileType == FileTypeCSS {
            cf := &CSSFile{File: f}
            cf.Selectors = ExtractSelectorsAndLinks(f)
            ad.CSSFiles = append(ad.CSSFiles, cf)
        }
    }

    for _, hf := range ad.HTMLFiles {
        for _, cf := range ad.CSSFiles {
            var selectors []*Selector
            var sel map[string]bool
            sel = make(map[string]bool)

            doc, err := goquery.NewDocumentFromReader(strings.NewReader(hf.Content))
            if err != nil {
                log.Fatal(err)
            }

            for _, s := range cf.Selectors {
                doc.Find(s.Name).Each(func(_ int, selection *goquery.Selection) {
                    if _, ok := sel[s.Name]; !ok {
                        selectors = append(selectors, s)
                        sel[s.Name] = true
                    }
                })
            }
            if selectors != nil && len(selectors) != 0 {
                hf.CSSFiles = append(hf.CSSFiles, &CSSFile{
                    File:      cf.File,
                    Selectors: selectors,
                })
            }
        }
    }

    return ad
}

func ExtractSelectorsAndLinks(file *File) []*Selector {
    var selectors []*Selector

    var sel map[string]bool
    sel = make(map[string]bool)

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
            if _, ok := sel[s]; !ok {
                selectors = append(selectors, &Selector{Name: s})
                sel[s] = true
            }
        }
    }
    return selectors
}
