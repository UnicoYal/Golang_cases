package pointerargs

// Parent is a parent struct.
type Parent struct {
	Name     string
	Children []Child
}

// Child is a child struct.
type Child struct {
	Name string
	Age  int
}

func CopyParent(p *Parent) Parent {
	if p != nil {
		copy_name := p.Name
		copy_children := make([]Child, len(p.Children))
		copy(copy_children, p.Children)
		return Parent{copy_name, copy_children}
	}
	return Parent{}
}