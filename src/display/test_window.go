package display

var TestWindow = NewComponentFactory(
	"TestWindow",
	NewNanoWindow,
	ID("Test Window"),
	VAlign(AlignCenter),
	HAlign(AlignCenter),
	BgColor(0x333333ff),
	Width(800),
	Height(600),
)
