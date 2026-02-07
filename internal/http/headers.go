package http

import "strings"

type Headers struct {
	content map[string]string
}

func (h *Headers) Len() int {
	return len(h.content)
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

func (h *Headers) String() string {
	var result strings.Builder

	for key, val := range h.content {
		result.WriteString(capitalize(key))
		result.WriteString(": ")
		result.WriteString(val)
		result.WriteString("\r\n")
	}

	result.WriteString("\r\n")

	return result.String()
}

func capitalize(key string) string {
	if strings.Contains(key, "-") {
		parts := strings.Split(key, "-")
		var result []string

		for _, part := range parts {
			result = append(result, capitalize(part))
		}

		return strings.Join(result, "-")
	}

	return strings.ToUpper(key[0:1]) + key[1:]
}

func NewHeaders() *Headers {
	return &Headers{
		content: make(map[string]string),
	}
}
