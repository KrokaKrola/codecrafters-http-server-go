package response

import "fmt"

type StatusLine struct {
	httpVersion  string
	statusCode   string
	reasonPhrase string
}

func (s *StatusLine) SetHttpVersion(httpVersion string) {
	s.httpVersion = httpVersion
}

func (s *StatusLine) SetStatusCode(statusCode string) {
	s.statusCode = statusCode
}

func (s *StatusLine) SetReasonPhrase(reasonPhrase string) {
	s.reasonPhrase = reasonPhrase
}

func (s *StatusLine) Stringify() string {
	result := fmt.Sprintf("%s %s", s.httpVersion, s.statusCode)
	if len(s.reasonPhrase) > 0 {
		result += fmt.Sprintf(" %s", s.reasonPhrase)
	}

	return result + "\r\n"
}

func NewStatusLine(httpVersion string) *StatusLine {
	return &StatusLine{
		httpVersion:  httpVersion,
		statusCode:   "",
		reasonPhrase: "",
	}
}

func New200StatusLine(httpVersion string) *StatusLine {
	return &StatusLine{
		httpVersion:  httpVersion,
		statusCode:   "200",
		reasonPhrase: "OK",
	}
}

func New404StatusLine(httpVersion string) *StatusLine {
	return &StatusLine{
		httpVersion:  httpVersion,
		statusCode:   "404",
		reasonPhrase: "Not Found",
	}
}
