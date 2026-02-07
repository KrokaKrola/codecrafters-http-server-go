package request

import "github.com/codecrafters-io/http-server-starter-go/internal/http"

type RequestLine struct {
	HttpMethod http.Method
	Target     string
	Version    http.Version
}

func NewRequestLine(httpMethod http.Method, target string, version http.Version) *RequestLine {
	return &RequestLine{
		HttpMethod: httpMethod,
		Target:     target,
		Version:    version,
	}
}
