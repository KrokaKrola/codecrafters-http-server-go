package response

import "fmt"

type StatusLine struct {
	httpVersion  string
	statusCode   string
	reasonPhrase string
}

func (s *StatusLine) Stringify() string {
	result := fmt.Sprintf("%s %s", s.httpVersion, s.statusCode)
	if len(s.reasonPhrase) > 0 {
		result += fmt.Sprintf(" %s", s.reasonPhrase)
	}

	return result + "\r\n"
}

func NewStatusLine(httpVersion string, statusCode string, reasonPhrase string) *StatusLine {
	return &StatusLine{
		httpVersion:  httpVersion,
		statusCode:   statusCode,
		reasonPhrase: reasonPhrase,
	}
}
