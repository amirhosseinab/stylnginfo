package app

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

    CSSFile struct {
        File      *File    `json:"file"`
        Selectors []string `json:"selectors"`
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
        switch f.FileType {
        case FileTypeCSS:

        }
    }
}
