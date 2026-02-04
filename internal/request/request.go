package request

import (
	"bufio"
)

type Request struct {
	reader     *bufio.Reader
	Body       Body
	Header     Header
	StatusLine StatusLine
}

func NewRequest(reader *bufio.Reader) *Request {
	return &Request{
		reader: reader,
	}
}

func (r *Request) Parse() {

}
