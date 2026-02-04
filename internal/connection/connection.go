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
	req.Parse()
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
	fmt.Println("req", req)

	resStatusLine := response.NewStatusLine("HTTP/1.1", "200", "OK")
	resHeaders := response.NewHeaders()
	resBody := response.NewBody()

	return response.NewResponse(resStatusLine, resHeaders, resBody)
}
