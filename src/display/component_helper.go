package display

func nearestFocusable(d Displayable) Displayable {
	result := d

	for !result.IsFocusable() {
		parent := result.Parent()
		if parent == nil {
			return result
		}
		result = parent
	}

	return result
}

// CursorPick will return the deepest Focusable node that contains the
// provided global coordinate.
//
// The search will begin at the provided node (usually root), and at each level,
// will step forward only along the child that contains the coordinate. Once a
// leaf is found, the code will walk back up until the nearest Focusable node
// is returned.
func CursorPick(d Displayable, globalX, globalY float64) Displayable {
	result := d

	children := d.Children()
	if len(children) == 0 {
		// We have reached a leaf, now walk back toward root and return the
		// first focusable element we find.
		return nearestFocusable(result)
	}

	for _, child := range children {
		if ContainsCoordinate(child, globalX, globalY) {
			result = CursorPick(child, globalX, globalY)
			break
		}
	}

	return result
}

func ContainsCoordinate(d Displayable, globalX, globalY float64) bool {
	dX, dY := LocalToGlobal(d, 0, 0)

	return globalX >= dX && globalX <= dX+d.Width() &&
		globalY >= dY && globalY <= dY+d.Height()
}

// LocalToGlobal returns the corresponding coordinate on the Global stage,
// given the component local coordinates.
func LocalToGlobal(d Displayable, localX, localY float64) (float64, float64) {
	parent := d.Parent()
	if parent != nil {
		return LocalToGlobal(parent, localX+d.X(), localY+d.Y())
	}
	return localX, localY
}
