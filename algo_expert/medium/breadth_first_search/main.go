package breadthfirstsearch

// Do not edit the class below except
// for the breadthFirstSearch method.
// Feel free to add new properties
// and methods to the class.
type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) BreadthFirstSearch(array []string) []string {
    res := make([]string, 0)
    queue := make([]*Node, 0)
    queue = append(queue, n)

    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        res = append(res, node.Name)

        children := node.Children
        for _, child := range children {
            queue = append(queue, child)
        }
    }
	return res
}
