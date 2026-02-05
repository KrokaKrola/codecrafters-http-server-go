package request

import "strings"

type Headers struct {
	content map[string]string
}

func (h *Headers) Set(name string, value string) {
	name = strings.TrimSpace(name)
	lowered := strings.ToLower(name)
	h.content[lowered] = value
}

func (h *Headers) Get(name string) (string, bool) {
	name = strings.TrimSpace(name)
	lowered := strings.ToLower(name)
	value, ok := h.content[lowered]
	return value, ok
}

func NewHeaders() *Headers {
	return &Headers{
		content: make(map[string]string),
	}
}
