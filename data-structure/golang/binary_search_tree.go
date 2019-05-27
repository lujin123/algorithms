/**
二叉查找树实现
1. Insert
2. Delete
3. Find
4. FindMax
5. FindMin
*/
package dsa

type BSTNode struct {
	Elem  int
	left  *BSTNode
	right *BSTNode
}

type BinarySearchTree struct {
	root *BSTNode
}

func (bst *BinarySearchTree) replace(oldNode, newNode, parent *BSTNode) {
	if parent == nil {
		bst.root = newNode
	} else {
		if parent.left == oldNode {
			parent.left = newNode
		} else {
			parent.right = newNode
		}
	}
}

func (bst *BinarySearchTree) Insert(x int) {
	if bst.root == nil {
		bst.root = &BSTNode{
			Elem: x,
		}
		return
	}

	for t := bst.root; t != nil; {
		if x == t.Elem {
			// x exists, do nothing.
			return
		}
		if x > t.Elem {
			if t.right == nil {
				t.right = &BSTNode{Elem: x}
				break
			}
			t = t.right
		} else {
			if t.left == nil {
				t.left = &BSTNode{Elem: x}
				break
			}
			t = t.left
		}
	}
}

func (bst *BinarySearchTree) Delete(x int) {
	var parent *BSTNode
	for t := bst.root; t != nil; {
		if x == t.Elem {
			if t.left != nil && t.right != nil {
				parent = t
				n := t.left
				flag := false
				for ; n.right != nil; n = n.right {
					if flag {
						parent = n
					} else {
						flag = true
					}
				}

				t.Elem = n.Elem
				t = n
			}
			if t.left != nil {
				bst.replace(t, t.left, parent)
			} else {
				bst.replace(t, t.right, parent)
			}
			break
		} else if x > t.Elem {
			parent = t
			t = t.right
		} else {
			parent = t
			t = t.left
		}
	}
}

func (bst *BinarySearchTree) Find(x int) *BSTNode {
	for t := bst.root; t != nil; {
		if x == t.Elem {
			return t
		} else if x > t.Elem {
			t = t.right
		} else {
			t = t.left
		}
	}
	return nil
}

func (bst *BinarySearchTree) FindMax(node *BSTNode) *BSTNode {
	t := node
	for ; t.right != nil; t = t.right {
	}
	return t
}

func (bst *BinarySearchTree) FindMin(node *BSTNode) *BSTNode {
	t := node
	for ; t.left != nil; t = t.left {
	}
	return t
}

func (bst *BinarySearchTree) InOrder(node *BSTNode, result *[]int) {
	if node != nil {
		bst.InOrder(node.left, result)
		*result = append(*result, node.Elem)
		bst.InOrder(node.right, result)
	}
}
