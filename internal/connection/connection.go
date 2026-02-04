package connection

import (
	"bufio"
	"fmt"
	"net"

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

	stringifiedResponse := res.Stringify()

	fmt.Println("response", stringifiedResponse)

	if _, err := c.writer.WriteString(stringifiedResponse); err != nil {
		return
	}

	if err := c.writer.Flush(); err != nil {
		return
	}
}

func (c *Connection) process(req *request.Request) *response.Response {
	var resStatusLine *response.StatusLine

	switch req.StatusLine.Target {
	case "/":
		resStatusLine = response.New200StatusLine("HTTP/1.1")
	case "/index.html":
		resStatusLine = response.New200StatusLine("HTTP/1.1")
	default:
		resStatusLine = response.New404StatusLine("HTTP/1.1")
	}

	resHeaders := response.NewHeaders()
	resBody := response.NewBody()

	return response.NewResponse(resStatusLine, resHeaders, resBody)
}
