package books

type Book struct {
	Title  string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
	Pages  int    `json:"pages,omitempty"`
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
