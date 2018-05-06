package ui

type Stateful interface {
	OnState(name string, options ...Option)
	ApplyCurrentState()
	HasState(name string) bool
	SetState(name string)
	State() string
}
