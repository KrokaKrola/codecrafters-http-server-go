package response

import "strings"

type Headers struct {
	content map[string]string
}

func (h *Headers) SetHeader(name string, value string) {
	name = strings.TrimSpace(name)
	lowered := strings.ToLower(name)
	h.content[lowered] = value
}

func (h *Headers) String() string {
	result := ""

	for key, val := range h.content {
		result += capitalize(key) + ": " + val + "\r\n"
	}

	return result + "\r\n"
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
