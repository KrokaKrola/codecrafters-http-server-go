package response

type Response struct {
	StatusLine *StatusLine
	Headers    *Headers
	Body       *Body
}

func NewResponse(statusLine *StatusLine, headers *Headers, body *Body) *Response {
	return &Response{
		StatusLine: statusLine,
		Headers:    headers,
		Body:       body,
	}
}

func (r *Response) String() string {
	statusLine := r.StatusLine.String()

	headers := r.Headers.String()
	body := r.Body.String()

	return statusLine + headers + body
}
