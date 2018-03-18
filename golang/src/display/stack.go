package display

type Stack interface {
	Push(entry Displayable)
	Pop() Displayable
	HasNext() bool
}

type stack struct {
	entries []Displayable
}

func (s *stack) Push(entry Displayable) {
	s.entries = append(s.entries, entry)
}

func (s *stack) Pop() Displayable {
	if s.HasNext() {
		result := s.entries[len(s.entries)-1]
		// This syntax just made me throw up in my mouth.
		s.entries = s.entries[:len(s.entries)-1]
		return result
	}
	return nil
}

func (s *stack) HasNext() bool {
	if len(s.entries) > 0 {
		return true
	}
	return false
}

func NewStack() Stack {
	entries := make([]Displayable, 0, 10)
	return &stack{entries: entries}
}
