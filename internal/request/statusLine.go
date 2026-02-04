package request

type StatusLine struct {
	HttpMethod string
	Target     string
	Version    string
}

func NewStatusLine(httpMethod string, target string, version string) *StatusLine {
	return &StatusLine{
		HttpMethod: httpMethod,
		Target:     target,
		Version:    version,
	}
}
