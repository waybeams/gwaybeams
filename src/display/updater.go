package display

import (
	"clock"
	"events"
)

// Update the tree from the provided Displayable forward.
// Existing nodes should be reused whenever possible.
func NewUpdater() Builder {
	return &Updater{}
}

type Updater struct {
	childrenTypeMap ChildrenTypeMap
	clock           clock.Clock
	emitter         Emitter
	isDestroyed     bool
	lastError       error
	root            Displayable
	stack           Stack
}

func (u *Updater) getEmitter() Emitter {
	if u.emitter == nil {
		u.emitter = NewEmitter()
	}
	return u.emitter
}

func (u *Updater) getStack() Stack {
	if u.stack == nil {
		u.stack = NewDisplayStack()
	}
	return u.stack
}

func (u *Updater) OnEnterFrame(handler EventHandler) Unsubscriber {
	return u.getEmitter().On(events.EnterFrame, handler)
}

func (u *Updater) Clock() clock.Clock {
	if u.clock == nil {
		// Go get the clock from the provided Root component's builder
		u.clock = clock.New()
	}
	return u.clock
}

func (u *Updater) LastError() error {
	return u.lastError
}

func (u *Updater) Peek() Displayable {
	return u.getStack().Peek()
}

func (u *Updater) getExistingChild(d Displayable, parent Displayable) Displayable {
	updateableChildren := parent.updateableChildren()
	typeName := d.TypeName()
	kids := updateableChildren[typeName]
	var result Displayable
	if kids != nil {
		if len(kids) > 1 {
			result = kids[0]
			updateableChildren[typeName] = kids[1:]
		} else {
			result = kids[0]
			delete(updateableChildren, typeName)
		}
		parent.setUpdateableChildren(updateableChildren)
	}
	return result
}

func (u *Updater) Update(d Displayable) error {
	u.Push(d)
	d.Layout()
	return u.lastError
}

func (u *Updater) Push(d Displayable, options ...ComponentOption) {
	stack := u.getStack()

	// Get the parent element if one exists
	parent := stack.Peek()

	if parent != nil {
		existingChild := u.getExistingChild(d, parent)
		if existingChild == nil {
			parent.AddChild(d)
		} else {
			d = existingChild
		}

		// Clear the composer function before triggering for what might be
		// a second time
		d.Composer(nil)
	} else {
		u.root = d
		d.SetBuilder(u)
	}

	// One of these options might be a Children(func()), which will recurse
	// back into this Push function.
	for _, option := range options {
		err := option(d)
		if err != nil {
			// If an option error is found, bail with it, for now.
			u.lastError = err
			return
		}
	}

	// Push the element onto the stack
	stack.Push(d)

	// Create the children type map for this parent
	d.setUpdateableChildren(u.createChildrenMapFor(d))

	// Process composition function to build children
	u.callComposeFunctionFor(d)

	// Remove the children that haven't been used.
	u.clearChildrenFromTypeMap(d)

	// Pop the element off the stack
	stack.Pop()
}

func (u *Updater) clearChildrenFromTypeMap(d Displayable) {
	types := d.updateableChildren()

	for _, entries := range types {
		for _, entry := range entries {
			entry.Parent().RemoveChild(entry)
		}
	}

	// Remove the children type mape for this parent
	u.childrenTypeMap = nil
}

func (u *Updater) createChildrenMapFor(d Displayable) ChildrenTypeMap {
	types := make(map[string][]Displayable)
	for _, child := range d.Children() {
		typeName := child.TypeName()
		entry := types[typeName]
		if entry == nil {
			types[typeName] = []Displayable{child}
		} else {
			types[typeName] = append(entry, child)
		}
	}
	return types
}

func (u *Updater) callComposeFunctionFor(d Displayable) {
	composeEmpty := d.GetComposeEmpty()
	if composeEmpty != nil {
		composeEmpty()
		return
	}
	composeWithBuilder := d.GetComposeWithBuilder()
	if composeWithBuilder != nil {
		composeWithBuilder(u)
		return
	}
	composeWithComponent := d.GetComposeWithComponent()
	if composeWithComponent != nil {
		composeWithComponent(d)
		return
	}
	composeWithBuilderAndComponent := d.GetComposeWithBuilderAndComponent()
	if composeWithBuilderAndComponent != nil {
		composeWithBuilderAndComponent(u, d)
		return
	}
}

func (u *Updater) Destroy() {
	u.stack = nil
	u.lastError = nil
	u.root = nil
	u.isDestroyed = true
}

func (u *Updater) Listen() {
	var frameHandler = func() bool {
		root := u.root
		if root != nil {
			u.getEmitter().Emit(NewEvent(events.EnterFrame, root, nil))

			if root.ShouldRecompose() {
				root.RecomposeChildren()
			}
			root.Layout()
		}

		return u.isDestroyed
	}
	clock.OnFrame(frameHandler, DefaultFrameRate, u.Clock())
}
