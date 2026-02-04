package request

import (
	"bufio"
	"bytes"
)

type Request struct {
	reader     *bufio.Reader
	Body       *Body
	Header     *Header
	StatusLine *StatusLine
}

func NewRequest(reader *bufio.Reader) *Request {
	return &Request{
		reader: reader,
	}
}

func (r *Request) Parse() error {
	line, _, err := r.reader.ReadLine()
	if err != nil {
		return err
	}

	parts := bytes.Split(line, []byte(" "))

	// TODO: validate request line parts
	r.StatusLine = NewStatusLine(
		string(parts[0]),
		string(parts[1]),
		string(parts[2]),
	)

	return nil
}
