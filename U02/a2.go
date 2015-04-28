package "main"

type List iterface {
	New() List
}

type node struct {
	item Int
	prev* node
	next* node
}

func New() List {
	return &node{}
}

