package request

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/internal/http"
)

type Request struct {
	reader     *bufio.Reader
	Body       *Body
	Headers    *Headers
	StatusLine *StatusLine
}

func NewRequest(reader *bufio.Reader) *Request {
	return &Request{
		reader: reader,
	}
}

func (r *Request) Parse() error {
	statusLine, err := r.parseStatusLine()
	if err != nil {
		return err
	}
	r.StatusLine = statusLine

	headers, err := r.parseHeaders()
	if err != nil {
		return err
	}
	r.Headers = headers

	body, err := r.parseBody()
	if err != nil {
		return err
	}
	r.Body = body

	return nil
}

func (r *Request) parseStatusLine() (*StatusLine, error) {
	line, _, err := r.reader.ReadLine()
	if err != nil {
		return nil, err
	}

	parts := bytes.Split(line, []byte(" "))

	method, err := http.ParseMethod(string(parts[0]))
	if err != nil {
		return nil, err
	}

	version, err := http.ParseVersion(string(parts[2]))
	if err != nil {
		return nil, err
	}

	return NewStatusLine(
		method,
		string(parts[1]),
		version,
	), nil
}

func (r *Request) parseHeaders() (*Headers, error) {
	headers := NewHeaders()

	for {
		line, _, err := r.reader.ReadLine()

		if err != nil {
			return nil, err
		}

		if bytes.Equal(line, []byte("")) {
			break
		}

		parts := bytes.Split(line, []byte(": "))
		key := string(parts[0])
		value := strings.Trim(string(parts[1]), " ")
		headers.Set(key, value)
	}

	return headers, nil
}

func (r *Request) parseBody() (*Body, error) {
	body := NewBody()
	cl, ok := r.Headers.Get("Content-Length")
	if !ok {
		body.SetContent("")
		return body, nil
	}

	cLen, err := strconv.ParseInt(cl, 10, 64)
	if err != nil {
		return nil, err
	}

	content := make([]byte, cLen)

	if _, err = r.reader.Read(content); err != nil {
		return nil, err
	}

	body.SetContent(string(content))

	return body, nil
}
