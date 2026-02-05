package response

import (
	"fmt"

	"github.com/codecrafters-io/http-server-starter-go/internal/http"
)

type StatusLine struct {
	httpVersion  http.Version
	statusCode   http.Status
	reasonPhrase string
}

func (s *StatusLine) SetHttpVersion(httpVersion http.Version) {
	s.httpVersion = httpVersion
}

func (s *StatusLine) SetStatusCode(statusCode http.Status) {
	s.statusCode = statusCode
}

func (s *StatusLine) SetReasonPhrase(reasonPhrase string) {
	s.reasonPhrase = reasonPhrase
}

func (s *StatusLine) String() string {
	result := fmt.Sprintf("%s %s", s.httpVersion, s.statusCode)
	if len(s.reasonPhrase) > 0 {
		result += fmt.Sprintf(" %s", s.reasonPhrase)
	}

	return result + "\r\n"
}

func NewStatusLine(httpVersion http.Version) *StatusLine {
	return &StatusLine{
		httpVersion: httpVersion,
	}
}

func New200StatusLine(httpVersion http.Version) *StatusLine {
	return &StatusLine{
		httpVersion:  httpVersion,
		statusCode:   http.StatusOK,
		reasonPhrase: http.StatusOK.Reason(),
	}
}

func New201StatusLine(httpVersion http.Version) *StatusLine {
	return &StatusLine{
		httpVersion:  httpVersion,
		statusCode:   http.StatusCreated,
		reasonPhrase: http.StatusCreated.Reason(),
	}
}

func New404StatusLine(httpVersion http.Version) *StatusLine {
	return &StatusLine{
		httpVersion:  httpVersion,
		statusCode:   http.StatusNotFound,
		reasonPhrase: http.StatusNotFound.Reason(),
	}
}

func New500StatusLine(httpVersion http.Version) *StatusLine {
	return &StatusLine{
		httpVersion:  httpVersion,
		statusCode:   http.ServerError,
		reasonPhrase: http.ServerError.Reason(),
	}
}
