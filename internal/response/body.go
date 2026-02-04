package response

type Body struct {
	content []byte
}

func NewBody() *Body {
	return &Body{}
}
