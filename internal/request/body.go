package request

type Body struct {
	content string
}

func NewBody() *Body {
	return &Body{}
}

func (b *Body) SetContent(content string) {
	b.content = content
}

func (b *Body) Content() string {
	return b.content
}
