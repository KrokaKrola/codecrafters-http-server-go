package handlers

import (
	"bytes"
	"compress/gzip"
	"strconv"

	"github.com/codecrafters-io/http-server-starter-go/internal/http"
	"github.com/codecrafters-io/http-server-starter-go/internal/request"
	"github.com/codecrafters-io/http-server-starter-go/internal/response"
	"github.com/codecrafters-io/http-server-starter-go/internal/router"
)

func HandleEcho() router.HandlerFunc {
	return func(req *request.Request) *response.Response {
		res := response.NewResponseByStatusCode(http.StatusOK)
		res.Headers.Set("Content-Type", "text/plain")

		encodings := http.GetSupportedEncodings(req.Headers)
		if len(encodings) > 0 {
			res.Headers.Set("Content-Encoding", "gzip")

			var buf bytes.Buffer
			zw := gzip.NewWriter(&buf)

			if _, err := zw.Write([]byte(req.Matches[1])); err != nil {
				res := response.NewResponseByStatusCode(500)
				zw.Close()
				return res
			}
			zw.Close()

			res.Headers.Set("Content-Length", strconv.Itoa(buf.Len()))
			res.Body = buf.String()
		} else {
			res.Headers.Set("Content-Length", strconv.Itoa(len(req.Matches[1])))
			res.Body = req.Matches[1]
		}

		return res
	}
}
