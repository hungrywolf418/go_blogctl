package model

import (
	"github.com/blend/go-sdk/selector"
	"github.com/wcharczuk/blogctl/pkg/constants"
)

// Posts is a list of posts.
type Posts []*Post

// First returns the first post in the list.
// It returns an empty post if the list is empty.
func (p Posts) First() (output Post) {
	if len(p) > 0 {
		output = *p[0]
	}
	return
}

// TableRows returns table rows for the given slice of posts.
func (p Posts) TableRows() []PostTableRow {
	output := make([]PostTableRow, len(p))
	for index := range p {
		output[index] = p[index].TableRow()
	}
	return output
}

// FilterBySelector filters the posts by a selector.
func (p Posts) FilterBySelector(sel selector.Selector) []*Post {
	var output []*Post
	for _, post := range p {
		if sel.Matches(post.Labels()) {
			output = append(output, post)
		}
	}
	return output
}

// Sort returns a sorter.
func (p Posts) Sort(key string, ascending bool) *PostSorter {
	return &PostSorter{
		Posts:     []*Post(p),
		SortKey:   key,
		Ascending: ascending,
	}
}

// PostSorter sorts a set of posts by a given sort key.
type PostSorter struct {
	Posts     []*Post
	SortKey   string
	Ascending bool
}

// Len implements sorter.
func (p PostSorter) Len() int {
	return len(p.Posts)
}

// Swap implements sorter.
func (p PostSorter) Swap(i, j int) {
	p.Posts[i], p.Posts[j] = p.Posts[j], p.Posts[i]
}

// Less implements sorter.
func (p PostSorter) Less(i, j int) bool {
	ip, jp := p.Posts[i], p.Posts[j]

	var output bool
	switch p.SortKey {
	case constants.PostSortKeyPosted:
		output = ip.Meta.Posted.After(jp.Meta.Posted)
	case constants.PostSortKeyCapture:
		it, jt := ip.Image.Exif.CaptureDate, jp.Image.Exif.CaptureDate
		if it.IsZero() {
			it = ip.Meta.Posted
		}
		if jt.IsZero() {
			jt = jp.Meta.Posted
		}
		output = it.After(jt)
	case constants.PostSortKeyIndex:
		output = ip.Index < jp.Index
	default:
		output = ip.Meta.Posted.After(jp.Meta.Posted)
	}

	if p.Ascending {
		return !output
	}
	return output
}
