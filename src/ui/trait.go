package ui

import (
	"strings"
)

type TraitOptions map[string][]Option

// Trait is a concrete factory function that builds a bag of ComponentOptions
// and applies them to all Selected Components before applying
// instance-specified options.
func Trait(c Context, selector string, opts ...Option) error {
	component := c.Builder().Peek()
	if component == nil {
		panic("Trait definition must be nested inside of a component")
	}

	// TODO(lbayes): There are more questions here than answers.
	// This entire feature is not fleshed out at all and will certainly behave
	// unexpectedly.
	component.PushTrait(selector, opts...)
	return nil
}

func mergeSelectOptions(result, next TraitOptions) TraitOptions {
	for key, value := range next {
		result[key] = value
	}
	return result
}

func QuerySelectorMatches(key string, d Displayable) bool {
	// Return for the "all" selector
	if key == "*" {
		return true
	}

	// Return for TypeName match
	if key == d.TypeName() {
		return true
	}

	var strippedKey = key[1:len(key)]

	// Return for ID match
	if strings.Index(key, "#") == 0 {
		if d.ID() == strippedKey {
			return true
		}

		// Return for ID style, key but no match
		return false
	}

	if strings.Index(key, ".") == 0 {
		// Return for any TraitName match
		for _, name := range d.TraitNames() {
			if strippedKey == name {
				return true
			}
		}
	}

	// The provided selector does not match this instance.
	return false
}

func TraitOptionsFor(d, parent Displayable) []Option {
	optionsMap := d.TraitOptions()
	current := parent
	for current != nil {
		optionsMap = mergeSelectOptions(optionsMap, current.TraitOptions())
		current = current.Parent()
	}

	result := []Option{}
	for key, value := range optionsMap {
		if QuerySelectorMatches(key, d) {
			result = append(result, value...)
		}
	}

	return result
}
