package display

import "errors"

type DisplayStack interface {
	Push(entry Displayable) error
	Pop() Displayable
	Peek() Displayable
	HasNext() bool
}

type displayStack struct {
	entries []Displayable
}

func (s *displayStack) Push(entry Displayable) error {
	if entry == nil {
		return errors.New("display.DisplayStack does not accept nil entries")
	}
	s.entries = append(s.entries, entry)
	return nil
}

func (s *displayStack) lastEntry() Displayable {
	return s.entries[len(s.entries)-1]
}

func (s *displayStack) Pop() Displayable {
	if s.HasNext() {
		result := s.lastEntry()
		// This syntax just made me throw up in my mouth.
		s.entries = s.entries[:len(s.entries)-1]
		return result
	}
	return nil
}

func (s *displayStack) Peek() Displayable {
	if s.HasNext() {
		return s.lastEntry()
	}
	return nil
}

func (s *displayStack) HasNext() bool {
	if len(s.entries) > 0 {
		return true
	}
	return false
}

func NewDisplayStack() DisplayStack {
	entries := make([]Displayable, 0, 10)
	return &displayStack{entries: entries}
}
