package http

import "fmt"

type Version string

func (v Version) String() string {
	return string(v)
}

func ParseVersion(s string) (Version, error) {
	switch s {
	case "HTTP/1.1":
		return Version("HTTP/1.1"), nil
	default:
		return Version(""), fmt.Errorf("invalid http version: %s", s)
	}
}

const (
	Version11 Version = "HTTP/1.1"
)
