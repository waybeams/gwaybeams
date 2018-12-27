package ctrl

import (
	"github.com/waybeams/examples/todo/model"
	"github.com/waybeams/waybeams/pkg/ctrl"
	"github.com/waybeams/waybeams/pkg/opts"
	"github.com/waybeams/waybeams/pkg/spec"
)

func ItemList(appModel *model.App, options ...spec.Option) spec.ReadWriter {
	return ctrl.VBox(
		opts.Key("Todo Items"),
		opts.MinHeight(300),
		opts.FlexWidth(1),
		opts.BgColor(0xeeeeeeff),
		opts.Childrenf(func() []spec.ReadWriter {
			result := []spec.ReadWriter{}
			for index, itemModel := range appModel.CurrentItems() {
				result = append(result, ItemSpec(itemModel, index))
			}
			return result
		}),
		opts.Bag(options...),
	)
}
