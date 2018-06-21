package fake

type BrowserCanvasContext struct {
}

type BrowserCanvas struct {
}

func (f *BrowserCanvas) GetContext2D() *BrowserCanvasContext {
	return &BrowserCanvasContext{}
}

func NewBrowserCanvas() {
	return &BrowserCanvas{}
}
