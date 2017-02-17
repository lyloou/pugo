package vars

const (
	Version     = "1.0.0"
	VersionDate = "2017-02-22 12:10"
)

const (
	FrontMatterTOML = iota + 1
)

var (
	FrontMatterBreak      = []byte("```")
	FrontMatterTOMLPrefix = []byte("toml")
)

// MetaFiles lists meta files
var MetaFiles = []string{
	"meta.toml",
}

// TimeFormatLayout lists supported time layouts
var TimeFormatLayout = []string{
	"2006-01-02",
	"2006-01-02 15:04",
	"2006-01-02 15:05:05",
}

// FrontMatterTypes lists front matter types and its prefix
var FrontMatterTypes = map[int][]byte{
	FrontMatterTOML: FrontMatterTOMLPrefix,
}
