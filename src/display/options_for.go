package display

func mergeSelectOptions(result, next TraitOptions) TraitOptions {
	for key, value := range next {
		result[key] = value
	}
	return result
}

func selectorMatches(key string, d Displayable, parent Displayable) bool {
	if key == d.GetTypeName() || key == "*" {
		return true
	}
	return false
}

func OptionsFor(d Displayable, parent Displayable) []ComponentOption {
	optionsMap := d.GetTraitOptions()
	current := parent
	for current != nil {
		optionsMap = mergeSelectOptions(optionsMap, current.GetTraitOptions())
		current = current.GetParent()
	}

	result := []ComponentOption{}
	for key, value := range optionsMap {
		if selectorMatches(key, d, parent) {
			result = append(result, value...)
		}
	}

	return result
}