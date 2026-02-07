package response

import "github.com/codecrafters-io/http-server-starter-go/internal/http"

type Response struct {
	StatusLine *StatusLine
	Headers    *http.Headers
	Body       string
}

func NewResponse(statusLine *StatusLine, headers *http.Headers, body string) *Response {
	return &Response{
		StatusLine: statusLine,
		Headers:    headers,
		Body:       body,
	}
}

func NewResponseByStatusCode(statusCode http.Status) *Response {
	sLine := &StatusLine{
		httpVersion:  http.Version11,
		statusCode:   statusCode,
		reasonPhrase: statusCode.Reason(),
	}

	return NewResponse(sLine, http.NewHeaders(), "")
}

func (r *Response) String() string {
	statusLine := r.StatusLine.String()

	headers := r.Headers.String()
	body := r.Body

	return statusLine + headers + body
}
