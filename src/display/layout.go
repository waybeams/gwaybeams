package display

type LayoutType int

const (
	StackLayoutType = iota
	FlowLayoutType
	RowLayoutType
)

type Layout interface {
}
