package spec

import "fmt"

// Apply will call each provided Option with the provided ReadWriter.
func Apply(rw ReadWriter, options ...Option) ReadWriter {
	for _, option := range options {
		option(rw)
	}
	return rw
}

// ApplyAll will take arbitrary slices of Options and will apply each set
// in order from left to right.
func ApplyAll(rw ReadWriter, optionSets ...[]Option) ReadWriter {
	options := []Option{}
	for _, optionSet := range optionSets {
		options = append(options, optionSet...)
	}
	return Apply(rw, options...)
}

func FirstChild(r Reader) ReadWriter {
	return r.Children()[0]
}

// Contains returns true if the ancestor contains the descendant.
func Contains(ancestor, descendant Reader) bool {
	current := descendant.Parent()
	for current != nil {
		if current == ancestor {
			return true
		}
		current = current.Parent()
	}

	return false
}

func Root(rw ReadWriter) ReadWriter {
	parent := rw.Parent()
	if parent != nil {
		return Root(parent)
	}
	return rw
}

func FilteredChildren(rw Reader, shouldInclude func(child Reader) bool) []ReadWriter {
	result := []ReadWriter{}
	for _, child := range rw.Children() {
		if shouldInclude(child) {
			result = append(result, child)
		}
	}
	return result
}

func FirstByKey(rw ReadWriter, key string) ReadWriter {
	if rw.Key() == key {
		return rw
	}
	for _, child := range rw.Children() {
		result := FirstByKey(child, key)
		if result != nil {
			return result
		}
	}
	return nil
}

func Path(r Reader) string {
	parent := r.Parent()
	localPath := "/" + pathPart(r)

	if parent != nil {
		return Path(parent) + localPath
	}
	return localPath
}

func pathPart(r Reader) string {
	// Try ID first
	// id := c.ID()
	// if id != "" {
	// return c.ID()
	// }

	// Empty ID, try Key
	key := r.Key()
	if key != "" {
		return r.Key()
	}

	parent := r.Parent()
	if parent != nil {
		siblings := parent.Children()
		for index, child := range siblings {
			// This comparison is why we can't have nice things.
			// This method does not work on base Spec because the
			// reference provided to the control methods is not the same
			// as the embedding structs that are provided to the Children
			// collection!
			if child == r {
				return fmt.Sprintf("%v%v", r.SpecName(), index)
			}
		}
	}

	// Empty ID and Key, and Parent just use TypeName
	return r.SpecName()
}

func NearestFocusable(r ReadWriter) ReadWriter {
	var candidate ReadWriter = r
	for candidate != nil {
		parent := candidate.Parent()
		if parent == nil || candidate.IsFocusable() {
			return candidate
		}
		candidate = candidate.Parent()
	}
	return nil
}

// ContainsCoordinate returns true if the provided global coordinate falls
// within the boundaries of the provided spec.Reader.
func ContainsCoordinate(r Reader, globalX, globalY float64) bool {
	dX, dY := LocalToGlobal(r, 0, 0)

	return globalX >= dX && globalX <= dX+r.Width() &&
		globalY >= dY && globalY <= dY+r.Height()
}

// CoordToControl will return the deepest Focusable node that contains the
// provided global coordinate.
//
// The search will begin at the provided node (usually root), and at each level,
// will step forward only along the child that contains the coordinate. Once a
// leaf is found, the code will walk back up until the nearest Focusable node
// is returned.
func CoordToControl(r ReadWriter, globalX, globalY float64) ReadWriter {
	result := r

	children := r.Children()
	if len(children) == 0 {
		// We have reached a leaf, now walk back toward root and return the
		// first focusable element we find.
		return NearestFocusable(r)
	}

	for _, child := range children {
		if ContainsCoordinate(child, globalX, globalY) {
			result = CoordToControl(child, globalX, globalY)
			break
		}
	}

	return result
}

// LocalToGlobal returns the corresponding coordinate on the Global stage,
// given the control local coordinates.
func LocalToGlobal(r Reader, localX, localY float64) (float64, float64) {
	parent := r.Parent()
	if parent != nil {
		return LocalToGlobal(parent, localX+r.X(), localY+r.Y())
	}
	return localX, localY
}
