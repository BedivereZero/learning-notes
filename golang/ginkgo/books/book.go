package books

import "strings"

type Book struct {
	Title  string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
	Pages  int    `json:"pages,omitempty"`
}

func (b *Book) IsValid() bool {
	var conditions = []bool{
		b.Title != "",
		b.Author != "",
		b.Pages > 0,
	}
	for _, c := range conditions {
		if c {
			continue
		}
		return false
	}
	return true
}

type Category string

const (
	CategoryNovel      Category = "Novel"
	CategoryShortStory Category = "ShortStory"
)

func (b *Book) Category() Category {
	if b.Pages < 300 {
		return CategoryShortStory
	}
	return CategoryNovel
}

func (b *Book) AuthorFirstName() string {
	before, _, found := strings.Cut(b.Author, " ")
	if !found {
		return ""
	}

	return before
}

func (b *Book) AuthorLastName() string {
	_, after, found := strings.Cut(b.Author, " ")
	if !found {
		return b.Author
	}

	return after
}
