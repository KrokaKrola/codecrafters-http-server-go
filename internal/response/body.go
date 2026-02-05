package response

type Body struct {
	content string
}

func (b *Body) String() string {
	return b.content
}

func (b *Body) SetContent(content string) {
	b.content = content
}

func NewBody() *Body {
	return &Body{}
}
