package dsa

const (
	DeltaHeight = 2
)

type Node struct {
	Left    *Node
	Right   *Node
	Element int
	height  int8
}

type AVLTree struct {
	root *Node
}

func (avl *AVLTree) height(node *Node) int8 {
	if node == nil {
		return -1
	}
	return node.height
}

func (avl *AVLTree) nodeHeight(node *Node) int8 {
	lh := avl.height(node.Left)
	rh := avl.height(node.Right)

	if lh > rh {
		return lh + 1
	}

	return rh + 1
}

func (avl *AVLTree) isBalance(node *Node) bool {
	lh := avl.height(node.Left)
	rh := avl.height(node.Right)
	delta := lh - rh
	if delta < 0 {
		delta = -delta
	}
	return delta < DeltaHeight
}

// 插入左子树的左节点
func (avl *AVLTree) singleRotateWithLeft(root *Node) *Node {
	t := root.Left
	root.Left = t.Right
	t.Right = root

	root.height = avl.nodeHeight(root)
	t.height = avl.nodeHeight(t)
	return t
}

// 插入右子树的右节点
func (avl *AVLTree) singleRotateWithRight(root *Node) *Node {
	t := root.Right
	root.Right = t.Left
	t.Left = root

	root.height = avl.nodeHeight(root)
	t.height = avl.nodeHeight(t)
	return t
}

// 插入左子树的右节点
func (avl *AVLTree) doubleRotateWithLeft(root *Node) *Node {
	root.Left = avl.singleRotateWithRight(root.Left)
	return avl.singleRotateWithLeft(root)
}

// 插入右子树的左节点
func (avl *AVLTree) doubleRotateWithRight(root *Node) *Node {
	root.Right = avl.singleRotateWithLeft(root.Right)
	return avl.singleRotateWithRight(root)
}

// 每次插入一个节点，都会从根节点往下递归，然后找到位置插入再进行平衡的调整，
// 但是也会导致从插入节点返回到根节点的递归调整树的高度
func (avl *AVLTree) insert(x int, node *Node) *Node {
	if node == nil {
		node = &Node{
			Element: x,
		}
	} else if x > node.Element {
		node.Right = avl.insert(x, node.Right)
		if !avl.isBalance(node) {
			if x > node.Right.Element {
				node = avl.singleRotateWithRight(node)
			} else {
				node = avl.doubleRotateWithRight(node)
			}
		}
	} else if x < node.Element {
		node.Left = avl.insert(x, node.Left)
		if !avl.isBalance(node) {
			if x < node.Left.Element {
				node = avl.singleRotateWithLeft(node)
			} else {
				node = avl.doubleRotateWithLeft(node)
			}
		}
	}
	// ignore x==node.Element, already in tree, do nothing.
	node.height = avl.nodeHeight(node)
	return node
}

// 插入节点
func (avl *AVLTree) Insert(element int) {
	avl.root = avl.insert(element, avl.root)
}

// 删除节点
// 1. 如果删除的节点是一个叶子节点，直接删除
// 2. 如果删除的节点不是叶子节点，则找到其左子树的最大节点交换
// 左子树的最大节点是其左子树的最右儿子节点，如果左子树右儿子节点为空，则根节点就是最大的
// 如果交换后还不是叶子节点，则继续交换
func (avl *AVLTree) Delete(element int) {
	avl.root = avl.delete(element, avl.root)
}

func (avl *AVLTree) delete(x int, node *Node) *Node {
	if node == nil {
		panic("data not in tree")
	} else if x > node.Element {
		node.Right = avl.delete(x, node.Right)
		if !avl.isBalance(node) {
			if avl.height(node.Left.Left) > avl.height(node.Left.Right) {
				node = avl.singleRotateWithLeft(node)
			} else {
				node = avl.doubleRotateWithLeft(node)
			}
		}
		node.height = avl.nodeHeight(node)
	} else if x < node.Element {
		node.Left = avl.delete(x, node.Left)
		if !avl.isBalance(node) {
			if avl.height(node.Right.Right) > avl.height(node.Right.Left) {
				node = avl.singleRotateWithRight(node)
			} else {
				node = avl.doubleRotateWithRight(node)
			}
		}
		node.height = avl.nodeHeight(node)
	} else if node.Left != nil && node.Right != nil {
		if avl.height(node.Left) > avl.height(node.Right) {
			maxNode := node.Left
			for ; maxNode.Right != nil; maxNode = maxNode.Right {
			}
			node.Element = maxNode.Element
			node.Left = avl.delete(maxNode.Element, node.Left)
		} else {
			minNode := node.Right
			for ; minNode.Left != nil; minNode = minNode.Left {
			}
			node.Element = minNode.Element
			node.Right = avl.delete(minNode.Element, node.Right)
		}
		node.height = avl.nodeHeight(node)
	} else {
		if node.Left != nil {
			node = node.Left
		} else {
			node = node.Right
		}
	}

	return node
}

func (avl *AVLTree) PreOrder(node *Node, result *[]int) {
	if node != nil {
		*result = append(*result, node.Element)
		avl.PreOrder(node.Left, result)
		avl.PreOrder(node.Right, result)
	}
}

func (avl *AVLTree) InOrder(node *Node, result *[]int) {
	if node != nil {
		avl.InOrder(node.Left, result)
		*result = append(*result, node.Element)
		avl.InOrder(node.Right, result)
	}
}
