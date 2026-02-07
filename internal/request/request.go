package request

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"strconv"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/internal/http"
)

const BytesIn50MB = 50 * 1000 * 1000

type Request struct {
	reader      *bufio.Reader
	Body        string
	Headers     *http.Headers
	RequestLine *RequestLine
	Ctx         context.Context
	Matches     []string
}

func NewRequest(reader *bufio.Reader, ctx context.Context) *Request {
	return &Request{
		reader: reader,
		Ctx:    ctx,
	}
}

func (r *Request) Parse() error {
	requestLine, err := r.parseRequestLine()
	if err != nil {
		return err
	}
	r.RequestLine = requestLine

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

func (r *Request) parseRequestLine() (*RequestLine, error) {
	line, _, err := r.reader.ReadLine()
	if err != nil {
		return nil, err
	}

	parts := bytes.Split(line, []byte(" "))

	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid request line")
	}

	method, err := http.ParseMethod(string(parts[0]))
	if err != nil {
		return nil, err
	}

	version, err := http.ParseVersion(string(parts[2]))
	if err != nil {
		return nil, err
	}

	return NewRequestLine(
		method,
		string(parts[1]),
		version,
	), nil
}

const MaxNumberOfHeaders = 100

func (r *Request) parseHeaders() (*http.Headers, error) {
	headers := http.NewHeaders()

	for {
		line, _, err := r.reader.ReadLine()

		if err != nil {
			return nil, err
		}

		if bytes.Equal(line, []byte("")) {
			break
		}

		parts := bytes.SplitN(line, []byte(": "), 2)
		if len(parts) < 2 {
			slog.Warn("invalid header structure", "line", line)
			continue
		}

		key := string(parts[0])
		value := strings.Trim(string(parts[1]), " ")
		headers.Set(key, value)

		if headers.Len() >= MaxNumberOfHeaders {
			return nil, fmt.Errorf("exceeded max number of headers for one request")
		}
	}

	return headers, nil
}

func (r *Request) parseBody() (string, error) {
	cl, ok := r.Headers.Get("Content-Length")
	if !ok {
		return "", nil
	}

	cLen, err := strconv.ParseInt(cl, 10, 64)
	if err != nil {
		return "", err
	}

	if cLen > BytesIn50MB {
		return "", fmt.Errorf("body size exceeds server max limit of 50mb")
	}

	content := make([]byte, cLen)

	if _, err = io.ReadFull(r.reader, content); err != nil {
		return "", err
	}

	return string(content), nil
}
