package ui

type Stateful interface {
	OnState(name string, options ...Option)
	ApplyCurrentState()
	HasState(name string) bool
	OptionsForState(stateName string) []Option
	SetState(name string)
	State() string
}
