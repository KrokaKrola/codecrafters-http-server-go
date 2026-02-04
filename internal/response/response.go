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

func (r *Response) Stringify() string {
	statusLine := r.StatusLine.Stringify()

	headers := "\r\n"
	body := ""

	return statusLine + headers + body
}
