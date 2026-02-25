package http

import "fmt"

type Version string

func (v Version) String() string {
	return string(v)
}

func ParseVersion(s string) (Version, error) {
	v := Version(s)
	switch v {
	case Version11:
		return v, nil
	default:
		return Version(""), fmt.Errorf("invalid http version: %s", s)
	}
}

const (
	Version11 Version = "HTTP/1.1"
)
