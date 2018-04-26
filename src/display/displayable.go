package display

import "errors"

// RenderHandler is a function type that will draw component state onto the provided
// Surface.
type RenderHandler func(s Surface, d Displayable) error

// Displayable entities can be composed, scaled, positioned, and drawn.
// This is the uber-interface for all Visual Elements in Waybeams.
type Displayable interface {
	Emitter
	Composable
	Layoutable
	Styleable
	Focusable
	Stateful
	Updateable

	// Text and Title are both kind of weird for the general
	// component case... Need to think more about this.
	Data() interface{}
	Draw(s Surface)
	InvalidNodes() []Displayable
	Invalidate()
	InvalidateChildren()
	InvalidateChildrenFor(d Displayable)
	PushTrait(sel string, opts ...ComponentOption) error
	PushUnsubscriber(Unsubscriber)
	SetData(data interface{})
	SetText(text string)
	SetTitle(title string)
	SetView(view RenderHandler)
	Text() string
	Title() string
	TraitOptions() TraitOptions
	UnsubAll()
	View() RenderHandler
}

func (c *Component) Data() interface{} {
	return c.Model().Data
}

func (c *Component) Draw(surface Surface) {
	local := surface.GetOffsetSurfaceFor(c)
	c.View()(local, c)
	c.DrawChildren(surface)
}

func (c *Component) DrawChildren(surface Surface) {
	for _, child := range c.Children() {
		// Create an surface delegate that includes an appropriate offset
		// for each child and send that to the Child's Draw() method.
		child.Draw(surface)
	}
}

func (c *Component) GetDefaultView() RenderHandler {
	return RectangleView
}

func (c *Component) InvalidNodes() []Displayable {
	nodes := c.dirtyNodes
	results := []Displayable{}
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
		// } else {
		// Invalidate the root node
		// c.InvalidateChildrenFor(c)
	}
}

func (c *Component) InvalidateChildren() {
	c.InvalidateChildrenFor(c)
}

func (c *Component) InvalidateChildrenFor(d Displayable) {
	// Late binding to find root at the time of invalidation.
	if c.Parent() != nil {
		c.Root().InvalidateChildrenFor(d)
		return
	}
	c.dirtyNodes = append(c.dirtyNodes, d)
}

func (c *Component) PushTrait(selector string, opts ...ComponentOption) error {
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

func (c *Component) SetView(view RenderHandler) {
	c.view = view
}

func (c *Component) Text() string {
	return c.Model().Text
}

func (c *Component) Title() string {
	return c.Model().Title
}

func (c *Component) TraitOptions() TraitOptions {
	if c.traitOptions == nil {
		c.traitOptions = make(map[string][]ComponentOption)
	}
	return c.traitOptions
}

func (c *Component) UnsubAll() {
	for _, unsub := range c.unsubs {
		unsub()
	}
}

func (c *Component) PushUnsubscriber(unsub Unsubscriber) {
	c.unsubs = append(c.unsubs, unsub)
}

func (c *Component) View() RenderHandler {
	if c.view == nil {
		return c.GetDefaultView()
	}
	return c.view
}
