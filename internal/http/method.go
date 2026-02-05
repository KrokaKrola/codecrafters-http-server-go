package http

import "fmt"

type Method string

func (m Method) String() string {
	return string(m)
}

func ParseMethod(s string) (Method, error) {
	switch s {
	case "GET":
		return Method("GET"), nil
	case "POST":
		return Method("POST"), nil
	case "PUT":
		return Method("PUT"), nil
	case "DELETE":
		return Method("DELETE"), nil
	default:
		return Method(""), fmt.Errorf("invalid HTTP method: %s", s)
	}
}

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	PATCH  Method = "PATCH"
	DELETE Method = "DELETE"
)
