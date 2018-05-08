package ui

type Option func(d Displayable)

// DelimiterOption is a Noop Flag option that separates declaration options from
// instance options.
func DelimiterOption(d Displayable) {}
