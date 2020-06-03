package core

type HTTPClient interface {
	Execute(request Request) Result
}
