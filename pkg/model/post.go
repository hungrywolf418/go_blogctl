package model

import (
	"fmt"
	"html/template"
	"path/filepath"
	"strings"
	"time"

	"github.com/wcharczuk/blogctl/pkg/constants"
)

// Post is a single post item.
type Post struct {
	OriginalPath string    `json:"originalPath,omitempty" yaml:"originalPath,omitempty"`
	OutputPath   string    `json:"outputPath,omitempty" yaml:"outputPath,omitempty"`
	Slug         string    `json:"slug,omitempty" yaml:"slug,omitempty"`
	Index        int       `json:"index" yaml:"index"`
	ModTime      time.Time `json:"modTime" yaml:"modTime"`

	Meta  Meta  `json:"meta" yaml:"meta"`
	Text  Text  `json:"text,omitempty" yaml:"text,omitempty"`
	Image Image `json:"image,omitempty" yaml:"image,omitempty"`

	Template *template.Template `json:"-" yaml:"-"`
	Previous *Post              `json:"-" yaml:"-"`
	Next     *Post              `json:"-" yaml:"-"`
}

// Labels returns labels use for filtering with a selector.
func (p Post) Labels() map[string]string {
	output := map[string]string{
		"title":    p.Meta.Title,
		"location": p.Meta.Location,
		"slug":     p.Slug,
		"postType": p.PostType(),
	}
	for _, tag := range p.Meta.Tags {
		output[tag] = "tagged"
	}
	return output
}

// PostType returns a string version of the post type.
func (p Post) PostType() string {
	if p.IsText() {
		return "text"
	}
	return "image"
}

// IsZero returns if the post is set.
func (p Post) IsZero() bool {
	return p.Text.IsZero() && p.Image.IsZero()
}

// IsImage returns if the post is an image post.
func (p Post) IsImage() bool {
	return !p.Image.IsZero()
}

// IsText returns if the post is an text post.
func (p Post) IsText() bool {
	return !p.Text.IsZero()
}

// HasPrevious returns if there is a previous post.
func (p Post) HasPrevious() bool {
	return p.Previous != nil && !p.Previous.IsZero()
}

// HasNext returns if there is a next post.
func (p Post) HasNext() bool {
	return p.Next != nil && !p.Next.IsZero()
}

// TitleOrDefault returns the title for the post.
func (p Post) TitleOrDefault() string {
	return p.Meta.Title
}

// IndexPath is a helper that returns the fully qualified path for the post's index.html.
// It is in the form /Year/Month/Day/Slug/index.html
func (p Post) IndexPath() string {
	return filepath.Join(p.Slug, constants.FileIndex)
}

// ImagePathOriginal returns the fully qualified image source path.
func (p Post) ImagePathOriginal() string {
	return filepath.Join(p.Slug, constants.FileImageOriginal)
}

// ImagePathForSize returns the image source for a given image size in pixels.
func (p Post) ImagePathForSize(size int) string {
	return filepath.Join(p.Slug, fmt.Sprintf(constants.ImageSizeFormat, size))
}

// ImagePathLarge returns the fully qualified image source path.
func (p Post) ImagePathLarge() string {
	return p.ImagePathForSize(constants.SizeLarge)
}

// ImagePathMedium returns the fully qualified image source path.
func (p Post) ImagePathMedium() string {
	return p.ImagePathForSize(constants.SizeMedium)
}

// ImagePathSmall returns the fully qualified image source path.
func (p Post) ImagePathSmall() string {
	return p.ImagePathForSize(constants.SizeSmall)
}

// TableRow returns a post as an ansi table row.
func (p Post) TableRow() PostTableRow {
	return PostTableRow{
		Title:    p.Meta.Title,
		Location: p.Meta.Location,
		Posted:   p.Meta.Posted,
		Slug:     p.Slug,
		Tags:     strings.Join(p.Meta.Tags, ", "),
		PostType: p.PostType(),
	}
}
