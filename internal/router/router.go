package router

import (
	"regexp"

	"github.com/codecrafters-io/http-server-starter-go/internal/http"
	"github.com/codecrafters-io/http-server-starter-go/internal/request"
	"github.com/codecrafters-io/http-server-starter-go/internal/response"
)

type HandlerFunc func(req *request.Request) *response.Response

type Router struct {
	routes []route
}

func NewRouter() *Router {
	return &Router{}
}

type route struct {
	pattern *regexp.Regexp
	method  http.Method
	handler HandlerFunc
}

func (r *Router) Handle(method http.Method, pattern string, handler HandlerFunc) {
	r.routes = append(r.routes, route{
		pattern: regexp.MustCompile(pattern),
		method:  method,
		handler: handler,
	})
}

func (r *Router) Match(req *request.Request) *response.Response {
	for _, rr := range r.routes {
		if req.RequestLine.HttpMethod == rr.method && rr.pattern.MatchString(req.RequestLine.Target) {
			matches := rr.pattern.FindStringSubmatch(req.RequestLine.Target)
			req.Matches = matches
			return rr.handler(req)
		}
	}

	res := response.NewResponseByStatusCode(http.StatusNotFound)

	return res
}
