package dsa

import (
	"math"
)

const (
	ColorRed = iota
	ColorBlack
)

type RBTNode struct {
	Elem  int
	left  *RBTNode
	right *RBTNode
	color int8
}

type RBTree struct {
	root    *RBTNode
	nilNode *RBTNode
}

// 初始化红黑树
func NewRBTree() *RBTree {
	rb := &RBTree{}
	nilnode := rb.newNode(math.MaxInt64, ColorBlack)
	nilnode.left = nilnode
	nilnode.right = nilnode
	rb.nilNode = nilnode
	rb.root = rb.newNode(math.MinInt64, ColorBlack)
	return rb
}

func (rb *RBTree) newNode(x int, color int8) *RBTNode {
	return &RBTNode{
		Elem:  x,
		left:  rb.nilNode,
		right: rb.nilNode,
		color: color,
	}
}

func (rb *RBTree) singleRotateWithLeft(n *RBTNode) *RBTNode {
	t := n.left
	n.left = t.right
	t.right = n
	return t
}

func (rb *RBTree) singleRotateWithRight(n *RBTNode) *RBTNode {
	t := n.right
	n.right = t.left
	t.left = n
	return t
}

func (rb *RBTree) doubleRotateWithLeft(n *RBTNode) *RBTNode {
	n.left = rb.singleRotateWithRight(n.left)
	return rb.singleRotateWithLeft(n)
}

func (rb *RBTree) doubleRotateWithRight(n *RBTNode) *RBTNode {
	n.right = rb.singleRotateWithLeft(n.right)
	return rb.singleRotateWithRight(n)
}

func (rb *RBTree) rotate(x int, parent *RBTNode) *RBTNode {
	if x > parent.Elem {
		if x > parent.right.Elem {
			parent.right = rb.singleRotateWithRight(parent.right)
		} else {
			parent.right = rb.singleRotateWithLeft(parent.right)
		}
		return parent.right
	} else {
		if x < parent.left.Elem {
			parent.left = rb.singleRotateWithLeft(parent.left)
		} else {
			parent.left = rb.singleRotateWithRight(parent.left)
		}
		return parent.left
	}
}

// 节点调整
func (rb *RBTree) reorient(x int, n, p, gp, ggp *RBTNode) {
	// color flip
	n.color = ColorRed
	n.left.color = ColorBlack
	n.right.color = ColorBlack

	// 如果父节点也是红色，则需要旋转调整
	if p.color == ColorRed {
		gp.color = ColorRed

		// 当前节点和父节点以及祖父节点不在一条线，需要旋转两次
		if (x < gp.Elem) != (x < p.Elem) {
			p = rb.rotate(x, gp)
		}
		n = rb.rotate(x, ggp)
		n.color = ColorBlack
	}
	// 最后调整根节点为黑色
	rb.root.right.color = ColorBlack
}

func (rb *RBTree) getRoot() *RBTNode {
	return rb.root.right
}

// 自顶向下的插入节点
func (rb *RBTree) Insert(x int) {
	var ggp *RBTNode
	n, p, gp := rb.root, rb.root, rb.root
	rb.nilNode.Elem = x
	for n.Elem != x {
		ggp, gp, p = gp, p, n
		if x > n.Elem {
			n = n.right
		} else {
			n = n.left
		}

		// 当节点的两个子节点都是红色的时候，将子节点改成黑色，自己变成红色
		if n.left.color == ColorRed && n.right.color == ColorRed {
			rb.reorient(x, n, p, gp, ggp)
		}
	}

	if n != rb.nilNode {
		// 重复的元素
		return
	}

	n = rb.newNode(x, ColorRed)
	if x > p.Elem {
		p.right = n
	} else {
		p.left = n
	}

	// 节点插入后需要再次调整
	rb.reorient(x, n, p, gp, ggp)
}

// 1. 如果n节点是红色的，直接下降一层；
// 2. 如果N是黑色的，并且两个子节点都是黑色的，则根据兄弟节点t子节点的颜色进行下一步
// 2-1. 如果t的两个节点都是黑色，直接翻转颜色
// 2-2. 如果t的节点不全是黑色的，需要进行旋转具体需要根据红色节点的位置
// 3. 如果n,p都是黑色的，那么t是红的，这时候需要旋转p,t
func (rb *RBTree) reorient2(n, t, p, gp *RBTNode) (*RBTNode, *RBTNode, *RBTNode, *RBTNode) {
	if p.color == ColorBlack && t.color == ColorRed {
		var tmp *RBTNode
		if t.Elem > p.Elem {
			tmp = rb.singleRotateWithRight(p)
			t.color = ColorBlack
			p.color = ColorRed
			t = p.right
		} else {
			tmp = rb.singleRotateWithLeft(p)
			t.color = ColorBlack
			p.color = ColorRed
			t = p.left
		}
		if p.Elem > gp.Elem {
			gp.right = tmp
		} else {
			gp.left = tmp
		}
		gp = tmp
	}

	if p.color == ColorRed && t.color == ColorBlack {
		if n.left.color == ColorBlack && n.right.color == ColorBlack {
			if t.left.color == ColorBlack && t.right.color == ColorBlack {
				p.color = ColorBlack
				n.color = ColorRed
				t.color = ColorRed
			} else {
				// 兄弟节点有一个红儿子，则需要进行选中，具体用哪种选择需要根绝红色节点的位置
				if t.Elem > p.Elem {
					var tmp *RBTNode
					if t.left.color == ColorRed {
						tmp = rb.doubleRotateWithRight(p)
					} else {
						tmp = rb.singleRotateWithRight(p)
					}
					if p.Elem > gp.Elem {
						gp.right = tmp
					} else {
						gp.left = tmp
					}
					t.right.color = ColorBlack
					p.color = ColorBlack
					n.color = ColorRed
					t = p.right
				} else {
					var tmp *RBTNode
					if t.right.color == ColorRed {
						tmp = rb.doubleRotateWithLeft(p)
					} else {
						tmp = rb.singleRotateWithLeft(p)
					}

					if p.Elem > gp.Elem {
						gp.right = tmp
					} else {
						gp.left = tmp
					}
					t.left.color = ColorBlack
					p.color = ColorBlack
					n.color = ColorRed
					t.color = ColorRed
					t = p.left
				}
			}
		}
	}

	return n, t, p, gp
}

func (rb *RBTree) delete(x int, n, t, p, gp *RBTNode) (*RBTNode, *RBTNode, *RBTNode, *RBTNode) {
	rb.nilNode.Elem = x
	for n.Elem != x {
		p, gp = n, p
		if x > n.Elem {
			t = p.left
			n = p.right
		} else {
			t = p.right
			n = p.left
		}

		if n.color == ColorBlack {
			n, t, p, gp = rb.reorient2(n, t, p, gp)
		}
	}
	return n, t, p, gp
}

// 自顶向下（top-down）的删除节点
func (rb *RBTree) Delete(x int) {
	rb.root.right.color = ColorRed
	n, t, p, gp := rb.root.right, rb.root.right, rb.root, rb.root
	var pre *RBTNode
	// 不断寻找待删除节点，直到叶子节点
	for {
		n, t, p, gp = rb.delete(x, n, t, p, gp)
		if pre != nil {
			pre.Elem = n.Elem
		}
		pre = n
		if n.left != rb.nilNode {
			node := rb.findMax(n.left)
			x = node.Elem
		} else if n.right != rb.nilNode {
			node := rb.findMin(n.right)
			x = node.Elem
		} else {
			break
		}
	}

	if n == rb.nilNode {
		return
	}

	// 这里找到的节点一定是红色的叶子叶子节点
	if p.right == n {
		p.right = rb.nilNode
	} else {
		p.left = rb.nilNode
	}

	// 重置根节点为黑色
	rb.root.right.color = ColorBlack
}

func (rb *RBTree) findMin(node *RBTNode) *RBTNode {
	for node.left != rb.nilNode {
		node = node.left
	}
	return node
}

func (rb *RBTree) findMax(node *RBTNode) *RBTNode {
	for node.right != rb.nilNode {
		node = node.right
	}
	return node
}

// 中序遍历
func (rb *RBTree) InOrder(node *RBTNode, result *[]int) {
	if node != rb.nilNode {
		rb.InOrder(node.left, result)
		*result = append(*result, node.Elem)
		rb.InOrder(node.right, result)
	}
}

// 检查红黑树适合满足条件
func (rb *RBTree) CheckRBTree() bool {
	// 空树也合法
	if rb.root.right == rb.nilNode {
		return true
	}
	// 根节点需要时黑的
	if rb.root.right.color != ColorBlack {
		return false
	}

	var res []int
	if !rb.checkRBTree(rb.root.right, 0, &res) {
		return false
	}

	for i := 1; i < len(res); i++ {
		if res[i] != res[i-1] {
			return false
		}
	}

	return true
}
func (rb *RBTree) checkRBTree(n *RBTNode, sum int, res *[]int) bool {
	if n == rb.nilNode {
		*res = append(*res, sum)
		return true
	}
	if n.color == ColorBlack {
		sum++
	} else {
		if n.left.color == ColorRed || n.right.color == ColorRed {
			return false
		}
	}

	return rb.checkRBTree(n.left, sum, res) && rb.checkRBTree(n.right, sum, res)
}
