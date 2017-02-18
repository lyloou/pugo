package vars

const (
	Version     = "1.0.0"
	VersionDate = "2017-02-22 12:10"
)

const (
	FrontMetaTOML = iota + 1
)

var (
	FrontMetaBreak      = []byte("```")
	FrontMetaTOMLPrefix = []byte("toml")
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

// FrontMetaTypes lists front matter types and its prefix
var FrontMetaTypes = map[int][]byte{
	FrontMetaTOML: FrontMetaTOMLPrefix,
}
