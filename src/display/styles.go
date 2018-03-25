package display

const DefaultStyleFontSize = 12
const DefaultStyleFontFace = "sans"
const DefaultStyleFontColor = 0x000

type StyleDefinition interface {
	Selector(sel StyleSelector) error
	GetSelector() StyleSelector

	BgColor(color uint)
	BorderColor(color uint)
	BorderSize(size int)
	FontColor(color uint)
	FontFace(face string)
	FontSize(size int)
	GetBgColor() uint
	GetBorderColor() uint
	GetBorderSize() int
	GetFontColor() uint
	GetFontFace() string
	GetFontSize() int
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

func (s *styleDefinition) BorderColor(color uint) {
	s.setValueAt("borderColor", color)
}

func (s *styleDefinition) GetBorderColor() uint {
	return s.getUintValueAt("borderColor")
}

func (s *styleDefinition) BorderSize(size int) {
	s.setValueAt("borderSize", size)
}

func (s *styleDefinition) GetBorderSize() int {
	return s.getIntValueAt("borderSize")
}

// Create a new StyleDefinition for a given component
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

// TODO(lbayes): Parse the string selector into some structured type
type StyleSelector string

func Style(b Builder, styles ...StyleOption) error {
	return nil
}

func validateSelector(expr StyleSelector) error {
	return nil
}
