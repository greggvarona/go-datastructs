/*Package trees contains tree structures.*/
package trees

import "github.com/greggvarona/go-datastructs/model"

// TreeNode represents a node in a tree
type TreeNode struct {
	data         interface{}
	neighborhood []model.Node
}

// Data returns the data stored in the node.
func (t *TreeNode) Data() interface{} {
	return t.data
}

// Neighborhood will return an array of related nodes (e.g. child nodes).
func (t *TreeNode) Neighborhood() []model.Node {
	return t.neighborhood
}
