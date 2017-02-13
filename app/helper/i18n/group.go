package i18n

import "sync"

// Group is a group of i18n objects
type Group struct {
	i18nData map[string]*I18n
	lock     sync.Mutex
}

// NewGroup creates a i18n group
func NewGroup() *Group {
	return &Group{
		i18nData: make(map[string]*I18n),
	}
}

// Set sets i18n object to group
func (g *Group) Set(in *I18n) {
	g.lock.Lock()
	g.i18nData[in.Lang] = in
	g.lock.Unlock()
}

// Get gets i18n object by language name
func (g *Group) Get(lang string) *I18n {
	g.lock.Lock()
	defer g.lock.Unlock()
	return g.i18nData[lang]
}

// Len returns the numbers of i18n in this group
func (g *Group) Len() int {
	return len(g.i18nData)
}

// Has returns whether language's i18n object in this group
func (g *Group) Has(lang string) bool {
	return g.i18nData[lang] != nil
}

// Names returns language names in this group
func (g *Group) Names() []string {
	var s []string
	for key := range g.i18nData {
		s = append(s, key)
	}
	return s
}
