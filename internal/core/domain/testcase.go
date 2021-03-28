package domain

type TestStatusCode int

const (
	OK             TestStatusCode = 200
	WARN           TestStatusCode = 300
	ERROR          TestStatusCode = 400
	INTERNAL_ERROR TestStatusCode = 500
)

type TestResults struct {
	StatusCode TestStatusCode
	Stdout     string
	Stderr     string
}
