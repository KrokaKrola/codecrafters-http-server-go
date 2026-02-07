package response

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/internal/http"
)

type StatusLine struct {
	httpVersion  http.Version
	statusCode   http.Status
	reasonPhrase string
}

func (s *StatusLine) String() string {
	var result strings.Builder
	fmt.Fprintf(&result, "%s %s", s.httpVersion.String(), s.statusCode)

	if len(s.reasonPhrase) > 0 {
		fmt.Fprintf(&result, " %s", s.reasonPhrase)
	}

	result.WriteString("\r\n")

	return result.String()
}
