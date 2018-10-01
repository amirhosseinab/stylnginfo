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
        HTMLFiles           []*File     `json:"htmlFiles"`
        CSSFiles            []*File     `json:"cssFiles"`
        Selectors           []*Selector `json:"selectors"`
        FileLinks           []*Link     `json:"fileLinks,omitempty"`
        CSSToSelectorLinks  []*Link     `json:"cssToSelectorLinks,omitempty"`
        HTMLToSelectorLinks []*Link     `json:"htmlToSelectorLinks,omitempty"`
    }

    File struct {
        Name     string   `json:"name"`
        Content  string   `json:"-"`
        FileType FileType `json:"fileType"`
    }

    Link struct {
        Source string `json:"source"`
        Target string `json:"target"`
    }

    Selector struct {
        Name         string         `json:"name"`
        Declarations []*Declaration `json:"declarations,omitempty"`
        MediaQuery   string         `json:"mediaQuery,omitempty"`
        CSSFileName  string         `json:"cssFileName"`
    }

    Declaration struct {
        Property string `json:"property"`
        Value    string `json:"value"`
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
    ad := &AnalyzedData{}

    for _, f := range files {
        if f.FileType == FileTypeHTML {
            ad.HTMLFiles = append(ad.HTMLFiles, f)
        }
        if f.FileType == FileTypeCSS {
            ad.CSSFiles = append(ad.CSSFiles, f)

            selectors, links := ExtractSelectorsAndLinks(f)
            ad.Selectors = append(ad.Selectors, selectors...)

            ad.CSSToSelectorLinks = append(ad.CSSToSelectorLinks, links...)
        }
    }

    for _, h := range ad.HTMLFiles {
       var flmap, hsmap map[string]bool
       flmap = make(map[string]bool)
       hsmap = make(map[string]bool)

       for _, s := range ad.Selectors {
           doc, err := goquery.NewDocumentFromReader(strings.NewReader(h.Content))
           if err != nil {
               log.Fatal(err)
           }
           doc.Find(s.Name).Each(func(_ int, selection *goquery.Selection) {
               if _, ok := flmap[h.Name+s.CSSFileName]; !ok {
                   ad.FileLinks = append(ad.FileLinks, &Link{Source: h.Name, Target: s.CSSFileName})
                   flmap[h.Name+s.CSSFileName] = true
               }
               if _, ok := hsmap[h.Name+s.Name]; !ok {
                   ad.HTMLToSelectorLinks = append(ad.HTMLToSelectorLinks, &Link{Source: h.Name, Target: s.Name})
                   hsmap[h.Name+s.Name] = true
               }
           })
       }
    }

    return ad
}


func ExtractSelectorsAndLinks(file *File) ([]*Selector, []*Link) {
    var selectors []*Selector
    var links []*Link

    var smap map[string]bool
    smap = make(map[string]bool)

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
            if _, ok := smap[s]; !ok {
                selectors = append(selectors, &Selector{Name: s, CSSFileName: file.Name})
                links = append(links, &Link{Source: file.Name, Target: s})
                smap[s] = true
            }
        }
    }
    return selectors, links
}
