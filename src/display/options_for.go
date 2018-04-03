package display

func mergeSelectOptions(result, next SelectOptions) SelectOptions {
	for key, value := range next {
		result[key] = value
	}
	return result
}

func selectorMatches(key string, d Displayable, parent Displayable) bool {
	if key == "*" {
		return true
	}
	return false
}

func OptionsFor(d Displayable, parent Displayable) []ComponentOption {
	optionsMap := d.GetSelectOptions()
	current := parent
	for current != nil {
		optionsMap = mergeSelectOptions(optionsMap, current.GetSelectOptions())
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
