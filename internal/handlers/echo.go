package handlers

import (
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
		} else {
			res.Headers.Set("Content-Length", strconv.Itoa(len(req.Matches[1])))
			res.Body = req.Matches[1]
		}

		return res
	}
}
