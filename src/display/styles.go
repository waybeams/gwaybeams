package display

import "errors"

type StyleKey int

const (
	BgColorKey = iota
)

func StyleKeyToName(key StyleKey) (string, error) {
	switch key {
	case BgColorKey:
		return "BgColor", nil
	}

	return "", errors.New("StyleKey requested, but not found")
}

type Attrs struct {
	BgColor  uint
	FontSize int
	FontFace string
	Select   string
}

type StyleOption func(d Displayable) error

func BgColor(color uint) StyleOption {
	return func(d Displayable) error {
		// d.SetStyle(BgColorKey, color)
		return nil
	}
}

func Style(b Builder, styles ...StyleOption) error {
	return nil
}
