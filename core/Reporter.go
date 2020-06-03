package core

type Reporter interface {
	Execute(result Result)
}
