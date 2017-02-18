package node

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/go-xiaohei/pugo/app/model/page"
	"github.com/go-xiaohei/pugo/app/model/post"
)

const (
	// NodePost is node type of a post
	NodePost = iota*2 + 1
	// NodePage is node type of a page
	NodePage
	// NodePagedPosts is node type of a paged post-list
	NodePagedPosts
	// NodeTagPosts is node type of a tagged post-list
	NodeTagPosts
	// NodeIndex is node type of index page
	NodeIndex
	// NodeArchive is node type of archive page
	NodeArchive
	// NodeXML is node type of xml page
	NodeXML
	// NodeNil is node type of nil page, a position to hold children
	NodeNil
)

type (
	// Node is a tree representation for site pages
	Node struct {
		Title    string
		URL      string
		Children []*Node
		Sort     int
		Type     int
	}
	nodes []*Node
)

func (ns nodes) Len() int           { return len(ns) }
func (ns nodes) Swap(i, j int)      { ns[i], ns[j] = ns[j], ns[i] }
func (ns nodes) Less(i, j int) bool { return ns[i].Sort < ns[j].Sort }

// NewTree creates new tree with path
func NewTree(p string) *Node {
	return &Node{
		Title: "@root",
		URL:   p,
		Sort:  0,
		Type:  0,
	}
}

// Len returns total nodes counter from this node
func (n *Node) Len() int {
	var counter int
	for _, c := range n.Children {
		counter += c.Len()
	}
	return counter + len(n.Children)
}

// SortChildren sorts children nodes
func (n *Node) SortChildren() {
	if len(n.Children) > 0 {
		sort.Sort(nodes(n.Children))
	}
}

// Print prints the node as tree level
func (n *Node) Print(prefix string) {
	fmt.Printf("%s[%s] %s [%d]\n", prefix, n.Title, n.URL, n.Type)
	for _, c := range n.Children {
		c.Print(prefix + "---")
	}
}

// Add adds node with path, title, node type and sort number
func (n *Node) Add(p, title string, t, sort int) {
	p = strings.TrimPrefix(filepath.ToSlash(p), "/")
	pSlice := strings.SplitN(p, "/", 2)
	if len(pSlice) < 1 || (len(pSlice) == 1 && pSlice[0] == "") {
		return
	}
	// try to find the node in children
	var currentNode *Node
	isFound := false
	for _, c := range n.Children {
		if c.URL == pSlice[0] {
			currentNode = c
			isFound = true
			break
		}
	}
	if !isFound {
		currentNode = &Node{
			Title: "",
			URL:   pSlice[0],
			Sort:  0,
			Type:  NodeNil,
		}
		n.Children = append(n.Children, currentNode)
	}
	if len(pSlice) == 1 || (len(pSlice) > 1 && pSlice[1] == "") {
		currentNode.Title = title
		currentNode.Type = t
		currentNode.Sort = sort
		return
	}
	// currentNode.Title = ""
	currentNode.Add(pSlice[1], title, t, sort)
}

// FillPosts fills posts to nodes
func (n *Node) FillPosts(posts []*post.Post) {
	for _, p := range posts {
		n.Add(p.URL(), p.Title, NodePost, 0)
	}
}

// FillPages fills posts to nodes
func (n *Node) FillPages(pages []*page.Page) {
	for _, p := range pages {
		n.Add(p.URL(), p.Title, NodePage, p.Sort)
	}
}

// FillPagedPosts fills post-list to nodes
func (n *Node) FillPagedPosts(pps []*post.PagedPosts) {
	var name string
	for _, pp := range pps {
		name = fmt.Sprintf("post-page-%d", pp.Pager.Current)
		n.Add(pp.URL(), name, NodePagedPosts, 0)
	}
}

// FillTagPosts fills post-list to nodes
func (n *Node) FillTagPosts(tps map[string]*post.TagPosts) {
	for _, tp := range tps {
		n.Add(tp.URL(), tp.Tag.Name, NodeTagPosts, 0)
	}
}

// FillCommonPages fills common pages to node
func (n *Node) FillCommonPages() {
	n.Add("index.html", "index", NodeIndex, 0)
	n.Add("archive.html", "archive", NodeArchive, 0)
	n.Add("feed.xml", "feed", NodeXML, 0)
	n.Add("sitemap.xml", "sitemap", NodeXML, 0)
}
