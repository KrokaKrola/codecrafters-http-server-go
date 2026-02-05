package connection

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/codecrafters-io/http-server-starter-go/internal/http"
	"github.com/codecrafters-io/http-server-starter-go/internal/request"
	"github.com/codecrafters-io/http-server-starter-go/internal/response"
)

type Connection struct {
	conn     net.Conn
	reader   *bufio.Reader
	writer   *bufio.Writer
	dirValue string
}

func NewConnection(conn net.Conn, dirValue string) *Connection {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	return &Connection{
		conn:     conn,
		reader:   reader,
		writer:   writer,
		dirValue: dirValue,
	}
}

func (c *Connection) Handle() {
	req := request.NewRequest(c.reader)
	if err := req.Parse(); err != nil {
		fmt.Printf("error while parsing request: %e\r\n", err)
		return
	}

	res := c.process(req)

	if _, err := c.writer.WriteString(res.String()); err != nil {
		fmt.Printf("error while writing response: %e\r\n", err)
		return
	}

	if err := c.writer.Flush(); err != nil {
		fmt.Printf("error while trying to flush writer: %e\r\n", err)
		return
	}
}

var indexUrlRegexp = regexp.MustCompile(`^/($|index\.html$)`)
var echoUrlRegexp = regexp.MustCompile(`^/echo/(\S+)$`)
var userAgentRegexp = regexp.MustCompile(`^/user-agent(/)?$`)
var filesUrlRegexp = regexp.MustCompile(`^/files/(\S+)$`)

func (c *Connection) process(req *request.Request) *response.Response {
	var resStatusLine *response.StatusLine
	resHeaders := response.NewHeaders()
	resBody := response.NewBody()

	switch {
	case indexUrlRegexp.MatchString(req.StatusLine.Target):
		resStatusLine = response.New200StatusLine(http.Version11)
	case echoUrlRegexp.MatchString(req.StatusLine.Target):
		submatches := echoUrlRegexp.FindStringSubmatch(req.StatusLine.Target)
		resStatusLine = response.New200StatusLine(http.Version11)
		resHeaders.SetHeader("Content-Type", "text/plain")
		resHeaders.SetHeader("Content-Length", strconv.Itoa(len(submatches[1])))
		resBody.SetContent(submatches[1])
	case userAgentRegexp.MatchString(req.StatusLine.Target):
		reqUa, ok := req.Headers.Get("User-Agent")
		if !ok {
			resStatusLine = response.New500StatusLine(http.Version11)
		} else {
			resStatusLine = response.New200StatusLine(http.Version11)
			resHeaders.SetHeader("Content-Type", "text/plain")
			resHeaders.SetHeader("Content-Length", strconv.Itoa(len(reqUa)))
			resBody.SetContent(reqUa)
		}
	case filesUrlRegexp.MatchString(req.StatusLine.Target):
		submatches := filesUrlRegexp.FindStringSubmatch(req.StatusLine.Target)
		path := filepath.Join(c.dirValue, submatches[1])

		if req.StatusLine.HttpMethod == http.GET {
			dat, err := os.ReadFile(path)
			if err != nil {
				resStatusLine = response.New404StatusLine(http.Version11)
			} else {
				resStatusLine = response.New200StatusLine(http.Version11)
				resHeaders.SetHeader("Content-Type", "application/octet-stream")
				resHeaders.SetHeader("Content-Length", strconv.Itoa(len(dat)))
				resBody.SetContent(string(dat))
			}
		} else if req.StatusLine.HttpMethod == http.POST {
			file, err := os.Create(path)
			if err != nil {
				resStatusLine = response.New500StatusLine(http.Version11)
			} else {
				if _, err := file.Write([]byte(req.Body.Content())); err != nil {
					resStatusLine = response.New500StatusLine(http.Version11)
				}

				if err := file.Close(); err != nil {
					resStatusLine = response.New500StatusLine(http.Version11)
				}
			}
			resStatusLine = response.New201StatusLine(http.Version11)
		}
	default:
		resStatusLine = response.New404StatusLine(http.Version11)
	}

	return response.NewResponse(resStatusLine, resHeaders, resBody)
}
