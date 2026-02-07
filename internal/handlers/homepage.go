package handlers

import (
	"github.com/codecrafters-io/http-server-starter-go/internal/http"
	"github.com/codecrafters-io/http-server-starter-go/internal/request"
	"github.com/codecrafters-io/http-server-starter-go/internal/response"
	"github.com/codecrafters-io/http-server-starter-go/internal/router"
)

func HandleHomepage() router.HandlerFunc {
	return func(req *request.Request) *response.Response {
		return response.NewResponseByStatusCode(http.StatusOK)
	}
}
