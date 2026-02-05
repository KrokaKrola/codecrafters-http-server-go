package request

import "github.com/codecrafters-io/http-server-starter-go/internal/http"

type StatusLine struct {
	HttpMethod http.Method
	Target     string
	Version    http.Version
}

func NewStatusLine(httpMethod http.Method, target string, version http.Version) *StatusLine {
	return &StatusLine{
		HttpMethod: httpMethod,
		Target:     target,
		Version:    version,
	}
}
