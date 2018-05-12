package control

import (
	"errors"
	"events"
	"ui"
	"views"
)

func (c *Control) Bubble(event events.Event) {
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

func (c *Control) DataAsString(key string) string {
	value := c.Data(key)
	if value == nil {
		return ""
	}
	return value.(string)
}

func (c *Control) Data(key string) interface{} {
	model := c.Model()
	if model.Data != nil {
		return model.Data[key]
	}
	return nil
}

func (c *Control) Draw(surface ui.Surface) {
	local := surface.GetOffsetSurfaceFor(c)
	c.View()(local, c)
	c.DrawChildren(surface)
}

func (c *Control) DrawChildren(surface ui.Surface) {
	for _, child := range c.Children() {
		// Create an surface delegate that includes an appropriate offset
		// for each child and send that to the Child's Draw() method.
		child.Draw(surface)
	}
}

func (c *Control) GetDefaultView() ui.RenderHandler {
	return views.RectangleView
}

func (c *Control) InvalidNodes() []ui.Displayable {
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

func (c *Control) Invalidate() {
	// NOTE(lbayes): This is not desired behavior, but it's what we've got right now.
	if c.Parent() != nil {
		c.Parent().InvalidateChildren()
	} else {
		// Invalidate the root node
		c.InvalidateChildrenFor(c)
	}
}

func (c *Control) InvalidateChildren() {
	c.InvalidateChildrenFor(c)
}

func (c *Control) InvalidateChildrenFor(d ui.Displayable) {
	// Late binding to find root at the time of invalidation.
	if c.Parent() != nil {
		c.Root().InvalidateChildrenFor(d)
		return
	}
	c.dirtyNodes = append(c.dirtyNodes, d)
}

func (c *Control) PushTrait(selector string, opts ...ui.Option) error {
	traitOptions := c.TraitOptions()
	if traitOptions[selector] != nil {
		return errors.New("duplicate trait selector found with:" + selector)
	}
	traitOptions[selector] = opts
	return nil
}

func (c *Control) SetData(key string, value interface{}) {
	model := c.Model()
	if model.Data == nil {
		model.Data = make(map[string]interface{})
	}
	model.Data[key] = value
}

func (c *Control) SetText(text string) {
	c.Model().Text = text
}

func (c *Control) SetTitle(title string) {
	c.Model().Title = title
}

func (c *Control) SetView(view ui.RenderHandler) {
	c.view = view
}

func (c *Control) Text() string {
	return c.Model().Text
}

func (c *Control) Title() string {
	return c.Model().Title
}

func (c *Control) TraitOptions() ui.TraitOptions {
	if c.traitOptions == nil {
		c.traitOptions = make(map[string][]ui.Option)
	}
	return c.traitOptions
}

func (c *Control) PushUnsub(unsub events.Unsubscriber) {
	c.unsubs = append(c.unsubs, unsub)
}

func (c *Control) UnsubAll() {
	for _, unsub := range c.unsubs {
		unsub()
	}
}

func (c *Control) View() ui.RenderHandler {
	if c.view == nil {
		return c.GetDefaultView()
	}
	return c.view
}
