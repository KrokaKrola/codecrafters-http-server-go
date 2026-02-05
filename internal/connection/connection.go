package connection

import (
	"bufio"
	"fmt"
	"net"
	"regexp"
	"strconv"

	"github.com/codecrafters-io/http-server-starter-go/internal/http"
	"github.com/codecrafters-io/http-server-starter-go/internal/request"
	"github.com/codecrafters-io/http-server-starter-go/internal/response"
)

type Connection struct {
	conn   net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
}

func NewConnection(conn net.Conn) *Connection {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	return &Connection{
		conn:   conn,
		reader: reader,
		writer: writer,
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
var echoUrlRegexp = regexp.MustCompile(`^/echo/(\w+)$`)
var userAgentRegexp = regexp.MustCompile(`^/user-agent(/)?$`)

func (c *Connection) process(req *request.Request) *response.Response {
	var resStatusLine *response.StatusLine
	var resHeaders *response.Headers
	var resBody *response.Body

	switch {
	case indexUrlRegexp.MatchString(req.StatusLine.Target):
		resStatusLine = response.New200StatusLine(http.Version11)
		resHeaders = response.NewHeaders()
		resBody = response.NewBody()
	case echoUrlRegexp.MatchString(req.StatusLine.Target):
		submatches := echoUrlRegexp.FindStringSubmatch(req.StatusLine.Target)
		resStatusLine = response.New200StatusLine(http.Version11)
		resHeaders = response.NewHeaders()
		resHeaders.SetHeader("Content-Type", "text/plain")
		resHeaders.SetHeader("Content-Length", strconv.Itoa(len(submatches[1])))
		resBody = response.NewBody()
		resBody.SetContent(submatches[1])
	case userAgentRegexp.MatchString(req.StatusLine.Target):
		reqUa, ok := req.Headers.Get("User-Agent")
		if !ok {
			resStatusLine = response.New500StatusLine(http.Version11)
		} else {
			resStatusLine = response.New200StatusLine(http.Version11)
			resHeaders = response.NewHeaders()
			resHeaders.SetHeader("Content-Type", "text/plain")
			resHeaders.SetHeader("Content-Length", strconv.Itoa(len(reqUa)))
			resBody = response.NewBody()
			resBody.SetContent(reqUa)
		}
	default:
		resStatusLine = response.New404StatusLine(http.Version11)
		resHeaders = response.NewHeaders()
		resBody = response.NewBody()
	}

	return response.NewResponse(resStatusLine, resHeaders, resBody)
}
