package trees

import "github.com/greggvarona/go-datastructs/model"

// BinaryTreeNode represents a node in a binary tree
type BinaryTreeNode struct {
	TreeNode
}

// NewBinaryTreeNode creates a new binary tree node with an option of providing the left and right node.
func NewBinaryTreeNode(data interface{}, l *BinaryTreeNode, r *BinaryTreeNode) *BinaryTreeNode {
	btree := &BinaryTreeNode{}

	btree.data = data
	btree.neighborhood = append(btree.neighborhood, l)
	btree.neighborhood = append(btree.neighborhood, r)
	return btree
}

// Left retrieves the left neighbor of a node.
func (btn *BinaryTreeNode) Left() *BinaryTreeNode {
	if n, ok := btn.neighborhood[0].(*BinaryTreeNode); ok {
		return n
	}
	return nil
}

// Right regrieves the right neighbor of a node.
func (btn *BinaryTreeNode) Right() *BinaryTreeNode {
	if n, ok := btn.neighborhood[1].(*BinaryTreeNode); ok {
		return n
	}
	return nil
}

// PreOrder traverses the binary (sub) tree in pre-order.
func PreOrder(btn *BinaryTreeNode, f model.VisitFunction) {
	if btn == nil {
		return
	}

	f(btn)

	if btn.Left() != nil {
		PreOrder(btn.Left(), f)
	}

	if btn.Right() != nil {
		PreOrder(btn.Right(), f)
	}
}

// InOrder traverses the binary (sub) tree in in-order.
func InOrder(btn *BinaryTreeNode, f model.VisitFunction) {
	if btn == nil {
		return
	}

	if btn.Left() != nil {
		InOrder(btn.Left(), f)
	}

	f(btn)

	if btn.Right() != nil {
		InOrder(btn.Right(), f)
	}
}

// PostOrder traverses the binary (sub) tree in post-order.
func PostOrder(btn *BinaryTreeNode, f model.VisitFunction) {
	if btn == nil {
		return
	}

	if btn.Left() != nil {
		PostOrder(btn.Left(), f)
	}

	if btn.Right() != nil {
		PostOrder(btn.Right(), f)
	}

	f(btn)
}

// BinaryTree is a graph with nodes that have 2 neighbors.
type BinaryTree struct {
	Root *BinaryTreeNode
}

// PreOrder traverses the binary tree in pre-order.
func (t *BinaryTree) PreOrder(f model.VisitFunction) {
	if t.Root == nil {
		return
	}
	PreOrder(t.Root, f)
}

// InOrder traverses the binary tree in in-order.
func (t *BinaryTree) InOrder(f model.VisitFunction) {
	if t.Root == nil {
		return
	}
	InOrder(t.Root, f)
}

// PostOrder traverses the binary tree in post-order.
func (t *BinaryTree) PostOrder(f model.VisitFunction) {
	if t.Root == nil {
		return
	}
	PostOrder(t.Root, f)
}
