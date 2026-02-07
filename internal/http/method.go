package http

import (
	"fmt"
)

type Method string

func (m Method) String() string {
	return string(m)
}

func ParseMethod(s string) (Method, error) {
	m := Method(s)
	switch m {
	case GET, POST, PUT, PATCH, DELETE:
		return m, nil
	default:
		return "", fmt.Errorf("invalid HTTP method: %s", s)
	}
}

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	PATCH  Method = "PATCH"
	DELETE Method = "DELETE"
)
