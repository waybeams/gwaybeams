package styles

func Styles(values ...interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	for i := 0; i < len(values); i++ {
		value := values[i].([2]interface{})
		name := value[0].(string)
		m[name] = value[1]
	}
	return m
}

// These declarations makes me want to gouge my eyes out.
// I'm trying to get an implementation that takes advantage of the type system
// and validates user input, but reads cleanly and elegantly for authors.
// check out /lang_test.go Style blocks for examples.
func BgColor(color uint) [2]interface{} {
	return [2]interface{}{"BgColor", color}
}

func BackgroundRGBA(r, g, b, a int) {
}

func BorderColor(color uint) [2]interface{} {
	return [2]interface{}{"BorderColor", color}
}

func BorderSize(size int) [2]interface{} {
	return [2]interface{}{"BorderSize", size}
}

/*
func BorderType(style string) StyleDefinition {
	return &styleDefinition{name: BorderTypeStyle, value: style}
}

func Margin(size int) StyleDefinition {
	return &styleDefinition{name: MarginStyle, value: size}
}

func Padding(size int) StyleDefinition {
	return &styleDefinition{name: PaddingStyle, value: size}
}
*/
