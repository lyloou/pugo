package base

import (
	"errors"
	"net/url"
	"path"

	"github.com/go-xiaohei/pugo/app/helper/i18n"
)

type (
	// Nav is item of navigation
	Nav struct {
		Link        string `toml:"link" ini:"link"`
		Title       string `toml:"title" ini:"title"`
		OriginTitle string `toml:"-" ini:"-"`
		IsBlank     bool   `toml:"blank" ini:"blank"`
		Icon        string `toml:"icon" ini:"icon"`
		Hover       string `toml:"hover" ini:"hover"`
		I18n        string `toml:"i18n" ini:"i18n"`
		IsRemote    bool   `toml:"-" ini:"-"`
		// IconClass   string `toml:"icon" ini:"icon"` // deprecated, old icon field name
	}
	// NavGroup is group if items of navigation
	NavGroup []*Nav
)

var (
	// ErrNavInvalid means nav item is invalid
	ErrNavInvalid = errors.New("nav's title or link is blank")
)

// Tr print nav title with i18n helper
func (n *Nav) Tr(in *i18n.I18n) string {
	return in.Tr("nav." + n.I18n)
}

// TrLink print nav link with i18n prefix
func (n *Nav) TrLink(in *i18n.I18n) string {
	if n.IsRemote {
		return n.Link
	}
	if n.I18n == "" {
		return n.Link
	}
	return path.Join(in.Lang, n.Link)
}

// TrTitle print nav title with i18n value.
// If i18n="", use Nav.Title
func (n *Nav) TrTitle(in *i18n.I18n) string {
	if n.I18n == "" {
		return n.Title
	}
	return in.Tr("nav." + n.I18n)
}

// Format formats navs are correct in group
func (ng NavGroup) Format() error {
	for _, n := range ng {
		if n.Link == "" || n.Title == "" {
			return ErrNavInvalid
		}
		if u, _ := url.Parse(n.Link); u != nil && u.Host != "" {
			n.IsRemote = true
		}
		if n.I18n == "" {
			n.I18n = n.Title
		}
		n.OriginTitle = n.Title
	}
	return nil
}
