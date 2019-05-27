/**
非递归实现的一个AVL树

这是对递归版本的优化：

插入：优化插入节点时候回溯计算节点高度和旋转操作，旋转的操作是一个常量（0，1，2），而对高度的调整其实也是一个常量（最多两次），
因为插入一个节点，会有两个结果：

a. 子树依旧保持平衡，无需旋转操作，就只需要调整高度即可，而增加节点可能会导致父节点高度变化，也可能不会变化，如果变化了，就调整继续回溯，如果不变化，直接结束即可
b. 子树不平衡，这时候就需要通过旋转来平衡，可以是一次或者两次旋转，这样会导致子树高度-1，这时候就抵消了增加节点带了的父节点的高度变化，所以旋转后可以直接结束，无需回溯

删除：在找到待删除的节点后，判断子树的数量：

a. 左右子树不全：找到存在的子树，连接到当前节点的父节点即可
b. 左右子树全都有： 这时候需要找一个节点做替换，可以先判断下左右子树的高度，就找子树比较高的那边的节点代替，如果是左子树，则找最大节点，如果是右子树，则找最小节点，找到后交换内容之后按照a去走

删除后平衡父节点，重新计算高度，只有在高度相同并且子树平衡的情况下停止回溯，否则继续回溯，重新计算高度
*/
package dsa

type Node2 struct {
	left   *Node2
	right  *Node2
	parent *Node2
	Elem   int
	height int
}

type AVLTree2 struct {
	root *Node2
}

// 获取当前节点的高度
func (avl *AVLTree2) height(node *Node2) int {
	if node != nil {
		return node.height
	}
	return -1
}

// 取最大值
func (avl *AVLTree2) max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 计算树高
func (avl *AVLTree2) maxHeight(node *Node2) int {
	lh := avl.height(node.left)
	rh := avl.height(node.right)
	return avl.max(lh, rh) + 1
}

// 找寻树的最小节点
func (avl *AVLTree2) minNode(node *Node2) *Node2 {
	t := node
	for t.left != nil {
		t = t.left
	}
	return t
}

// 找寻树的最大节点
func (avl *AVLTree2) maxNode(node *Node2) *Node2 {
	t := node
	for t.right != nil {
		t = t.right
	}
	return t
}

// 计算左右子树高度差
func (avl *AVLTree2) nodeDelta(node *Node2) (delta int) {
	lh := avl.height(node.left)
	rh := avl.height(node.right)
	delta = lh - rh
	return
}

// 计算树高和左右子树高度差
func (avl *AVLTree2) nodeInfo(node *Node2) (height, delta int) {
	lh := avl.height(node.left)
	rh := avl.height(node.right)
	delta = lh - rh
	height = avl.max(lh, rh) + 1
	return
}

// 重新将旋转后的子树连接到原来的树上
func (avl *AVLTree2) childReplace(oldNode, newNode, parent *Node2) {
	if parent == nil {
		avl.root = newNode
	} else {
		if parent.left == oldNode {
			parent.left = newNode
		} else {
			parent.right = newNode
		}
	}
	if newNode != nil {
		newNode.parent = parent
	}
}

func (avl *AVLTree2) singleRotateWithLeft(node *Node2) *Node2 {
	t := node.left
	node.left = t.right
	avl.childReplace(node, t, node.parent)
	if t.right != nil {
		t.right.parent = node
	}
	t.right = node
	node.parent = t
	node.height = avl.maxHeight(node)
	t.height = avl.maxHeight(t)
	return t
}

func (avl *AVLTree2) singleRotateWithRight(node *Node2) *Node2 {
	t := node.right
	node.right = t.left
	avl.childReplace(node, t, node.parent)
	if t.left != nil {
		t.left.parent = node
	}
	t.left = node
	node.parent = t
	node.height = avl.maxHeight(node)
	t.height = avl.maxHeight(t)
	return t
}

func (avl *AVLTree2) doubleRotateWithLeft(node *Node2) *Node2 {
	node.left = avl.singleRotateWithRight(node.left)
	return avl.singleRotateWithLeft(node)
}

func (avl *AVLTree2) doubleRotateWithRight(node *Node2) *Node2 {
	node.right = avl.singleRotateWithLeft(node.right)
	return avl.singleRotateWithRight(node)
}

// 重新平衡，往根部回溯
// 1. 如果遇到节点高度没有变化，显然，就可以停止回溯了
// 2. 如果节点的高度有变化，但是子树的高度差小于等于1，那么更新当前根节点的高度之后继续回溯
// 3. 如果节点高度有变化并且子树高度差大于1了，则需要进行旋转，
// 但是旋转过的节点高度必定会-1，因为插入一个节点导致高度+1，所以旋转后节点不需要继续回溯
// 总结：向上回溯的时候挨个更新节点的左右平衡值
// 如果遇到“插入就会导致不平衡”的节点A就旋转并停止回溯
// 如果遇到“插入不会导致子树高度变化“的节点B并停止回溯
// 注意剪枝，这样效率更高，并不需要每次都回溯到根节点
func (avl *AVLTree2) rebalance(node *Node2, x int) {
	for ; node != nil; node = node.parent {
		height, delta := avl.nodeInfo(node)
		// 节点高度没有变化，停止回溯
		if height == node.height {
			break
		}

		if delta > 1 {
			if x < node.left.Elem {
				node = avl.singleRotateWithLeft(node)
			} else {
				node = avl.doubleRotateWithLeft(node)
			}
			break
		}

		if delta < -1 {
			if x > node.right.Elem {
				node = avl.singleRotateWithRight(node)
			} else {
				node = avl.doubleRotateWithRight(node)
			}
			break
		}
		// 更新节点变化，继续回溯
		node.height = height
	}
}

func (avl *AVLTree2) rebalance2(node *Node2) {
	for t := node; t != nil; t = t.parent {
		height, delta := avl.nodeInfo(node)
		if height != t.height {
			t.height = height
		} else if delta >= -1 && delta <= 1 {
			break
		}

		if delta > 1 {
			if avl.nodeDelta(t.left) > 0 {
				t = avl.singleRotateWithLeft(t)
			} else {
				t = avl.doubleRotateWithLeft(t)
			}
		} else if delta < -1 {
			if avl.nodeDelta(t.right) < 0 {
				t = avl.singleRotateWithRight(t)
			} else {
				t = avl.doubleRotateWithRight(t)
			}
		}
	}
}

func (avl *AVLTree2) Add(x int) {
	if avl.root == nil {
		avl.root = &Node2{
			Elem: x,
		}
		return
	}
	for t := avl.root; t != nil; {
		// x exists, do nothing.
		if x == t.Elem {
			return
		}
		if x > t.Elem {
			if t.right == nil {
				avl.insert(t, x)
				break
			}
			t = t.right
		} else {
			if t.left == nil {
				avl.insert(t, x)
				break
			}
			t = t.left
		}
	}
}

func (avl *AVLTree2) insert(node *Node2, x int) {
	child := &Node2{
		Elem:   x,
		parent: node,
	}
	if x > node.Elem {
		node.right = child
	} else {
		node.left = child
	}

	avl.rebalance(node, x)
}

// 每次都找一个叶子节点去替代待删除的节点
func (avl *AVLTree2) Remove(x int) {
	for node := avl.root; node != nil; {
		if x == node.Elem {
			if node.left != nil && node.right != nil {
				var t *Node2
				if avl.nodeDelta(node) > 0 {
					t = avl.maxNode(node.left)
				} else {
					t = avl.minNode(node.right)
				}
				node.Elem = t.Elem
				node = t
			}
			parent := node.parent
			var t *Node2
			if node.left != nil {
				t = node.left
			} else {
				t = node.right
			}
			avl.childReplace(node, t, parent)
			if parent != nil {
				avl.rebalance2(parent)
			}
			break
		} else if x > node.Elem {
			node = node.right
		} else {
			node = node.left
		}
	}
}

func (avl *AVLTree2) InOrder(node *Node2, result *[]int) {
	if node != nil {
		avl.InOrder(node.left, result)
		*result = append(*result, node.Elem)
		avl.InOrder(node.right, result)
	}
}
