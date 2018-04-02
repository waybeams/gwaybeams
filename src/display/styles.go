package display

// DefaultStyleFontSize is 12.
const DefaultStyleFontSize = 12

// DefaultStyleFontFace is "sans".
const DefaultStyleFontFace = "sans"

// DefaultStyleFontColor is black.
const DefaultStyleFontColor = 0x000000ff

// StyleDefinition is a bag of style names and values.
type StyleDefinition interface {
	Selector(sel StyleSelector) error
	GetSelector() StyleSelector

	BgColor(color uint)
	StrokeSize(size int)
	FontColor(color uint)
	FontFace(face string)
	FontSize(size int)
	GetBgColor() uint
	GetStrokeSize() int
	GetFontColor() uint
	GetFontFace() string
	GetFontSize() int
	GetStrokeColor() uint
	StrokeColor(color uint)
}

type styleBag map[string]interface{}

type styleDefinition struct {
	displayable Displayable
	selector    StyleSelector
	styles      styleBag
}

func (s *styleDefinition) getBag() styleBag {
	if s.styles == nil {
		s.styles = make(map[string]interface{})
	}
	return s.styles
}

func (s *styleDefinition) Selector(sel StyleSelector) error {
	s.selector = sel
	return nil
}

func (s *styleDefinition) GetSelector() StyleSelector {
	return s.selector
}

func (s *styleDefinition) GetDisplayable() Displayable {
	return s.displayable
}

// Collection of methods to help fight ludicrous duplication caused
// by scalar type enforcement.
func (s *styleDefinition) setValueAt(name string, value interface{}) {
	s.getBag()[name] = value
}

func (s *styleDefinition) hasEntryFor(name string) bool {
	return s.getBag()[name] != nil
}

// Collection of methods to help fight ludicrous duplication caused
// by scalar type enforcement.
func (s *styleDefinition) getUintValueAt(name string) uint {
	bag := s.getBag()
	if bag[name] != nil {
		return s.getBag()[name].(uint)
	}
	return 0
}

// Collection of methods to help fight ludicrous duplication caused
// by scalar type enforcement.
func (s *styleDefinition) getIntValueAt(name string) int {
	if s.hasEntryFor(name) {
		return s.getBag()[name].(int)
	}
	return 0
}

// Collection of methods to help fight ludicrous duplication caused
// by scalar type enforcement.
func (s *styleDefinition) getBoolValueAt(name string) bool {
	if s.hasEntryFor(name) {
		return s.getBag()[name].(bool)
	}
	return false
}

// Collection of methods to help fight ludicrous duplication caused
// by scalar type enforcement.
func (s *styleDefinition) getStringValueAt(name string) string {
	if s.hasEntryFor(name) {
		return s.getBag()[name].(string)
	}
	return ""
}

func (s *styleDefinition) BgColor(color uint) {
	s.setValueAt("bgColor", color)
}

func (s *styleDefinition) GetBgColor() uint {
	return s.getUintValueAt("bgColor")
}

func (s *styleDefinition) FontSize(size int) {
	s.setValueAt("fontSize", size)
}

func (s *styleDefinition) GetFontSize() int {
	return s.getIntValueAt("fontSize")
}

func (s *styleDefinition) FontColor(color uint) {
	s.setValueAt("fontColor", color)
}

func (s *styleDefinition) GetFontColor() uint {
	return s.getUintValueAt("fontColor")
}

func (s *styleDefinition) FontFace(face string) {
	s.setValueAt("fontFace", face)
}

func (s *styleDefinition) GetFontFace() string {
	return s.getStringValueAt("fontFace")
}

func (s *styleDefinition) StrokeColor(color uint) {
	s.setValueAt("strokeColor", color)
}

func (s *styleDefinition) GetStrokeColor() uint {
	return s.getUintValueAt("strokeColor")
}

func (s *styleDefinition) StrokeSize(size int) {
	s.setValueAt("strokeSize", size)
}

func (s *styleDefinition) GetStrokeSize() int {
	return s.getIntValueAt("strokeSize")
}

func NewStyleDefinition() StyleDefinition {
	definition := &styleDefinition{}
	return definition
}

func NewDefaultStyleDefinition() StyleDefinition {
	definition := NewStyleDefinition()
	definition.FontSize(DefaultStyleFontSize)
	definition.FontFace(DefaultStyleFontFace)
	definition.FontColor(DefaultStyleFontColor)
	return definition
}

type StyleOption func(StyleDefinition) error

func BgColor(color uint) StyleOption {
	return func(s StyleDefinition) error {
		s.BgColor(color)
		return nil
	}
}

func FontFace(face string) StyleOption {
	return func(s StyleDefinition) error {
		s.FontFace(face)
		return nil
	}
}

func FontSize(size int) StyleOption {
	return func(s StyleDefinition) error {
		s.FontSize(size)
		return nil
	}
}

func Selector(sel StyleSelector) StyleOption {
	return func(s StyleDefinition) error {
		s.Selector(sel)
		return nil
	}
}

type StyleName string

type StyleSelector string

func Style(b Builder, styles ...StyleOption) error {
	return nil
}

func validateSelector(expr StyleSelector) error {
	return nil
}
