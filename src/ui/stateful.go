package ui

type Stateful interface {
	OnState(name string, options ...Option)
	ApplyCurrentState() error
	HasState(name string) bool
	SetState(name string)
	State() string
}
