package display

import "fmt"

func Render(root Displayable, surface Surface) error {

	nodeHandler := func(node Displayable) {
		fmt.Println("HELLO!")
	}

	PostOrderVisit(root, nodeHandler)

	return nil
}

func layout(root Displayable) error {
	return nil
}
