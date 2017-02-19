package page

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/go-xiaohei/pugo/app/helper/markdown"
	"github.com/go-xiaohei/pugo/app/model"
	"github.com/go-xiaohei/pugo/app/model/author"
	"github.com/go-xiaohei/pugo/app/model/index"
	"github.com/go-xiaohei/pugo/app/vars"
)

var (
	_ model.Content = (*Page)(nil)
	_ model.Index   = (*Page)(nil)
)

var (
	// ErrPageFrontMetaFail means it can't detect front-matter block in post bytes
	ErrPageFrontMetaFail = errors.New("detect front-matter fail")
	// ErrPageFrontMetaTypeUnknown means it can't parse front-matter block with known types
	ErrPageFrontMetaTypeUnknown = errors.New("can't detect front-matter's format")
	// ErrPageFrontMetaTimeError means wrong time format in front-matter block
	ErrPageFrontMetaTimeError = errors.New("time format error in front-matter")
)

type (
	// Page is an object for one page content
	Page struct {
		Title      string `toml:"title"`
		Desc       string `toml:"desc"`
		Created    string `toml:"date"`
		Updated    string `toml:"update_date"`
		AuthorName string `toml:"author"`
		Sort       int    `toml:"sort"`
		IsDraft    bool   `toml:"draft"`
		IsNode     bool   `toml:"node"`
		Hover      string `toml:"hover"`
		Lang       string `toml:"lang"`
		Template   string `toml:"template"`

		index  []*index.Index
		slug   string
		author *author.Author

		contentBytes []byte
		srcBytes     []byte
		srcFile      string
		created      time.Time
		updated      time.Time

		frontMetaBytes []byte
		frontMetaType  int
	}
)

func (p *Page) detectFrontMeta() error {
	dataSlice := bytes.SplitN(p.srcBytes, vars.FrontMetaBreak, 3)
	if len(dataSlice) != 3 {
		return ErrPageFrontMetaFail
	}
	frontBytes := bytes.TrimSpace(dataSlice[1])
	for t, prefix := range vars.FrontMetaTypes {
		if bytes.HasPrefix(frontBytes, prefix) {
			frontBytes = bytes.TrimPrefix(frontBytes, prefix)
			p.frontMetaBytes = frontBytes
			p.frontMetaType = t
			p.contentBytes = bytes.TrimSpace(dataSlice[2])
			return nil
		}
	}
	return ErrPageFrontMetaTypeUnknown
}

func (p *Page) parseFrontMeta() error {
	var err error
	if err = toml.Unmarshal(p.frontMetaBytes, p); err != nil {
		return err
	}
	if err = p.formatTime(); err != nil {
		return ErrPageFrontMetaTimeError
	}
	return nil
}

func (p *Page) formatTime() error {
	var err error
	if p.Created == "" {
		p.getCreateTime()
	} else {
		for _, layout := range vars.TimeFormatLayout {
			p.created, err = time.Parse(layout, p.Created)
			if err == nil {
				break
			}
		}
		if err != nil {
			return err
		}
	}
	if p.Updated == "" {
		p.Updated = p.Created
		p.updated = p.created
	} else {
		for _, layout := range vars.TimeFormatLayout {
			p.updated, err = time.Parse(layout, p.Updated)
			if err == nil {
				break
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Page) getCreateTime() {
	if p.srcFile != "" {
		if info, _ := os.Stat(p.srcFile); info != nil {
			p.created = info.ModTime()
			p.Created = p.created.Format("2006-01-02 15:04")
		}
	}
}

func (p *Page) render() {
	if !p.IsNode {
		p.contentBytes = markdown.Render(p.contentBytes)
	}
}

func (p *Page) getIndex() {
	if !p.IsNode {
		p.index = index.New(p.contentBytes)
	}
}

// New parses bytes to a *Page
func New(data []byte, slug string, srcFile string) (*Page, error) {
	var (
		err error
		p   = &Page{
			srcBytes: data,
			slug:     slug,
			srcFile:  srcFile,
		}
	)
	if err = p.detectFrontMeta(); err != nil {
		return nil, err
	}
	if err = p.parseFrontMeta(); err != nil {
		return nil, err
	}
	p.render()
	p.getIndex()
	return p, nil
}

// NewFromFile parses file to a *Page
func NewFromFile(file, slug string) (*Page, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	p, err := New(data, slug, file)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Content returns page's content
func (p *Page) Content() []byte {
	return p.contentBytes
}

// ContentHTML returns page's content html type
func (p *Page) ContentHTML() template.HTML {
	return template.HTML(p.contentBytes)
}

// ContentLength return page's content length
func (p *Page) ContentLength() int {
	return len(p.contentBytes)
}

// CreateTime returns page's created time
func (p *Page) CreateTime() time.Time {
	return p.created
}

// UpdateTime returns page's updated time
func (p *Page) UpdateTime() time.Time {
	return p.updated
}

// DstFile returns rendered destination filepath
func (p *Page) DstFile() string {
	return fmt.Sprintf("%s.html", p.slug)
}

// SrcFile returns source filepath
func (p *Page) SrcFile() string {
	return p.srcFile
}

// URL returns site link for this post
func (p *Page) URL() string {
	return fmt.Sprintf("%s.html", p.slug)
}

// Index returns content index for post
func (p *Page) Index() []*index.Index {
	return p.index
}

// Author gets the author pf this post
func (p *Page) Author() *author.Author {
	return p.author
}
