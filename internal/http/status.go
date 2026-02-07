package http

import "strconv"

type Status int

const (
	StatusOK          Status = 200
	StatusCreated     Status = 201
	StatusNotFound    Status = 404
	StatusServerError Status = 500
)

var Reasons = map[Status]string{
	StatusOK:          "OK",
	StatusCreated:     "Created",
	StatusNotFound:    "Not Found",
	StatusServerError: "Internal Server Error",
}

func (s Status) String() string {
	return strconv.Itoa(int(s))
}

func (s Status) Reason() string {
	return Reasons[s]
}
