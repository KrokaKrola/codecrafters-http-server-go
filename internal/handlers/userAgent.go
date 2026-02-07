package handlers

import (
	"strconv"

	"github.com/codecrafters-io/http-server-starter-go/internal/http"
	"github.com/codecrafters-io/http-server-starter-go/internal/request"
	"github.com/codecrafters-io/http-server-starter-go/internal/response"
	"github.com/codecrafters-io/http-server-starter-go/internal/router"
)

func HandleUserAgent() router.HandlerFunc {
	return func(req *request.Request) *response.Response {
		reqUa, ok := req.Headers.Get("User-Agent")
		if !ok {
			return response.NewResponseByStatusCode(http.StatusServerError)
		}

		res := response.NewResponseByStatusCode(http.StatusOK)

		res.Headers.Set("Content-Type", "text/plain")
		res.Headers.Set("Content-Length", strconv.Itoa(len(reqUa)))
		res.Body = reqUa

		return res
	}
}
