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
    FileTypeScss
    FileTypeLess
)

type (
    FileType int

    ScrutinyData struct {
        HTMLFiles []*HTMLFile `json:"htmlFiles"`
        CSSFiles  []*CSSFile  `json:"cssFiles"`
    }

    HTMLFile struct {
        *File
        CSSFiles []*CSSFile `json:"cssFiles,omitempty"`
    }

    CSSFile struct {
        *File
        Selectors []*Selector `json:"selectors"`
        IsUsed    bool        `json:"isUsed"`
    }

    LessFile struct {
        *File
        ImportedFiles []*LessFile
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

func ScrutinizeCSSFiles(indexFile *File, files []*File) []*CSSFile {
    var fs []*CSSFile
    used := getUsedCSSFileNames(indexFile)

    for _, f := range files {
        _, isUsed := used[f.Name]
        fs = append(fs, &CSSFile{
            File:      f,
            Selectors: extractSelectorsAndLinks(f),
            IsUsed:    isUsed,
        })
    }
    return fs
}
func ScrutinizeLessFiles(files []*File) []*LessFile {
    var fs []*LessFile
    for _, f := range files {
        fs = append(fs, &LessFile{
            File:          f,
            ImportedFiles: getImportedLessFiles(f, files),
        })
    }
    return fs
}

func getImportedLessFiles(file *File, files []*File) []*LessFile {
    re, err := regexp.Compile(`(?m)^@import\s+"(.*)";\s?$`)
    if err != nil {
        log.Fatal(err)
    }
    var lessFiles []*LessFile
    fs := re.FindAllStringSubmatch(file.Content, -1)
    if fs != nil {
        for _, f := range fs {
            var file *File
            fileName := f[1]

            if !strings.HasSuffix(fileName, ".less") {
                fileName = fileName + ".less"
            }
            for _, tf := range files {
                if tf.Name == fileName {
                    file = tf
                    break
                }
            }
            lessFiles = append(lessFiles, &LessFile{
                File:          file,
                ImportedFiles: getImportedLessFiles(file, files),
            })
        }
    }
    return lessFiles
}

func getUsedCSSFileNames(indexFile *File) map[string]bool {
    var fileNames map[string]bool
    fileNames = make(map[string]bool)

    reLinks, err := regexp.Compile(`(?m)^\s*<link\s.*?rel\s?=\s?"stylesheet"\s+\/>\s*$`)
    if err != nil {
        log.Fatal(err)
    }
    links := reLinks.FindAllString(indexFile.Content, -1)
    rePath, err := regexp.Compile(`\w+\/.*?\.css`)
    for _, l := range links {
        p := rePath.FindString(l)
        fn := p[strings.LastIndex(p, "/")+1:]
        fileNames[fn] = true
    }

    return fileNames
}

func AnalyzeFiles(files ...*File) *ScrutinyData {
    ad := &ScrutinyData{}
    for _, f := range files {
        if f.FileType == FileTypeHTML {
            ad.HTMLFiles = append(ad.HTMLFiles, &HTMLFile{File: f})
        }
        if f.FileType == FileTypeCSS {
            cf := &CSSFile{File: f}
            cf.Selectors = extractSelectorsAndLinks(f)
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

func extractSelectorsAndLinks(file *File) []*Selector {
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
