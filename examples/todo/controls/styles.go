package controls

import (
	"github.com/waybeams/waybeams/pkg/opts"
	"github.com/waybeams/waybeams/pkg/spec"
)

type Styles struct {
	Box            spec.Option
	Button         spec.Option
	Header         spec.Option
	Main           spec.Option
	selectedFilter spec.Option
}

func (s *Styles) SelectedFilter(isSelected bool) spec.Option {
	if isSelected {
		return s.selectedFilter
	}
	return opts.Empty()
}

func CreateStyles() *Styles {
	boxStyle := opts.Bag(
		opts.BgColor(0xffffffff),
		opts.Padding(10),
		opts.Gutter(10),
	)

	return &Styles{
		Box: boxStyle,
		Button: opts.Bag(
			opts.BgColor(0xf8f8f8ff),
		),
		Header: opts.Bag(
			opts.FontColor(0xaf2f2f26),
			opts.FontFace("Roboto Light"),
			opts.FontSize(100),
		),
		Main: opts.Bag(
			boxStyle,
			opts.FontColor(0x111111ff),
			opts.FontFace("Roboto"),
			opts.FontSize(24),
		),
		selectedFilter: opts.Bag(
			opts.StrokeColor(0x0000ffff),
			opts.StrokeSize(1),
		),
	}
}
