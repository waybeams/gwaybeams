package comp

import (
	"errors"
	"events"
	"ui"
	"views"
)

func (c *Component) Bubble(event events.Event) {
	c.Emit(event)

	current := c.Parent()
	for current != nil {
		if event.IsCancelled() {
			return
		}
		current.Emit(event)
		current = current.Parent()
	}
}

func (c *Component) Data() interface{} {
	return c.Model().Data
}

func (c *Component) Draw(surface ui.Surface) {
	local := surface.GetOffsetSurfaceFor(c)
	c.View()(local, c)
	c.DrawChildren(surface)
}

func (c *Component) DrawChildren(surface ui.Surface) {
	for _, child := range c.Children() {
		// Create an surface delegate that includes an appropriate offset
		// for each child and send that to the Child's Draw() method.
		child.Draw(surface)
	}
}

func (c *Component) GetDefaultView() ui.RenderHandler {
	return views.RectangleView
}

func (c *Component) InvalidNodes() []ui.Displayable {
	nodes := c.dirtyNodes
	results := []ui.Displayable{}
	for nIndex, node := range nodes {
		ancestorFound := false
		for aIndex, possibleAncestor := range nodes {
			if node != nil && aIndex != nIndex && node.IsContainedBy(possibleAncestor) {
				ancestorFound = true
				break
			}
		}
		if !ancestorFound {
			results = append(results, node)
		}
	}

	return results
}

func (c *Component) Invalidate() {
	// NOTE(lbayes): This is not desired behavior, but it's what we've got right now.
	if c.Parent() != nil {
		c.Parent().InvalidateChildren()
	} else {
		// Invalidate the root node
		c.InvalidateChildrenFor(c)
	}
}

func (c *Component) InvalidateChildren() {
	c.InvalidateChildrenFor(c)
}

func (c *Component) InvalidateChildrenFor(d ui.Displayable) {
	// Late binding to find root at the time of invalidation.
	if c.Parent() != nil {
		c.Root().InvalidateChildrenFor(d)
		return
	}
	c.dirtyNodes = append(c.dirtyNodes, d)
}

func (c *Component) PushTrait(selector string, opts ...ui.Option) error {
	traitOptions := c.TraitOptions()
	if traitOptions[selector] != nil {
		return errors.New("duplicate trait selector found with:" + selector)
	}
	traitOptions[selector] = opts
	return nil
}

func (c *Component) SetData(data interface{}) {
	c.Model().Data = data
}

func (c *Component) SetText(text string) {
	c.Model().Text = text
}

func (c *Component) SetTitle(title string) {
	c.Model().Title = title
}

func (c *Component) SetView(view ui.RenderHandler) {
	c.view = view
}

func (c *Component) Text() string {
	return c.Model().Text
}

func (c *Component) Title() string {
	return c.Model().Title
}

func (c *Component) TraitOptions() ui.TraitOptions {
	if c.traitOptions == nil {
		c.traitOptions = make(map[string][]ui.Option)
	}
	return c.traitOptions
}

func (c *Component) PushUnsub(unsub events.Unsubscriber) {
	c.unsubs = append(c.unsubs, unsub)
}

func (c *Component) UnsubAll() {
	for _, unsub := range c.unsubs {
		unsub()
	}
}

func (c *Component) View() ui.RenderHandler {
	if c.view == nil {
		return c.GetDefaultView()
	}
	return c.view
}
