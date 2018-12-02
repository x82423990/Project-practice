package main

type tree struct {
	value       int
	left, right *tree
}

func sort(root *tree, value []int) {

	for _, k := range value {
		add(root, k)
	}
	appendValues(value[:0], root)
}


func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}



func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
func main() {
	var root *tree
	value := []int{2, 4, 5, 6, 7, 90, 23, 14, 88}
	sort(root,value)

}
