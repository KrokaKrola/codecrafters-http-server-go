package response

type Headers struct {
	content map[string]string
}

func NewHeaders() *Headers {
	return &Headers{}
}
