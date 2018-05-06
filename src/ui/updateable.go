package ui

type ChildrenTypeMap map[string][]Displayable

type Updateable interface {
	SetUpdateableChildren(types ChildrenTypeMap)
	UpdateableChildren() ChildrenTypeMap
}
