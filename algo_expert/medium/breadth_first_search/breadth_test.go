package breadthfirstsearch

import (
	"fmt"
	"testing"
)

func TestBreadth(t *testing.T) {
	root := &Node{
		Name: "A",
		Children: []*Node{
			{Name: "B", Children: []*Node{}},
			{Name: "C", Children: []*Node{}},
			{Name: "D", Children: []*Node{}},
		},
	}

	result := root.BreadthFirstSearch(make([]string, 0))
	fmt.Println(result)
}
