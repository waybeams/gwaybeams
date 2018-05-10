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
