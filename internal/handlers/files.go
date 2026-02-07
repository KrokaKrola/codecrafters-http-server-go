package handlers

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/codecrafters-io/http-server-starter-go/internal/http"
	"github.com/codecrafters-io/http-server-starter-go/internal/request"
	"github.com/codecrafters-io/http-server-starter-go/internal/response"
	"github.com/codecrafters-io/http-server-starter-go/internal/router"
)

func HandleFilesGet(dirValue string) router.HandlerFunc {
	return func(req *request.Request) *response.Response {
		path := filepath.Join(dirValue, req.Matches[1])

		dat, err := os.ReadFile(path)
		if err != nil {
			return response.NewResponseByStatusCode(http.StatusNotFound)
		}

		res := response.NewResponseByStatusCode(http.StatusOK)
		res.Headers.Set("Content-Type", "application/octet-stream")
		res.Headers.Set("Content-Length", strconv.Itoa(len(dat)))
		res.Body = string(dat)

		return res
	}
}

func HandleFilesPost(dirValue string) router.HandlerFunc {
	return func(req *request.Request) *response.Response {
		path := filepath.Join(dirValue, req.Matches[1])
		file, err := os.Create(path)

		if err != nil {
			return response.NewResponseByStatusCode(http.StatusServerError)
		}

		if _, err := file.Write([]byte(req.Body)); err != nil {
			return response.NewResponseByStatusCode(http.StatusServerError)
		}

		if err := file.Close(); err != nil {
			return response.NewResponseByStatusCode(http.StatusServerError)
		}

		return response.NewResponseByStatusCode(http.StatusCreated)
	}
}
