package dsa

import (
	"math"
	"math/rand"
	"time"
)

type SLNode struct {
	Elem   int
	next   *SLNode
	bottom *SLNode
}

type SkipList struct {
	levels []*SLNode
}

func NewSkiplist() *SkipList {
	sl := &SkipList{}
	sl.levels = append(sl.levels, &SLNode{Elem: math.MinInt64})
	return sl
}

func (sl *SkipList) Find(x int) bool {
	height := len(sl.levels)
	p := sl.levels[height-1]
	n := p.next

	for {
		if n == nil || n.Elem > x {
			if p.bottom == nil {
				break
			}
			p = p.bottom
			n = p.next
			continue
		}

		if n.Elem == x {
			return true
		}

		p = n
		n = n.next
	}
	return false
}

func (sl *SkipList) Insert(x int) {
	height := len(sl.levels)
	var ps []*SLNode
	p := sl.levels[height-1]
	n := p.next

	for {
		if n == nil || n.Elem > x {
			if p.bottom == nil {
				break
			}
			ps = append(ps, p)
			p = p.bottom
			n = p.next
			continue
		}
		if n.Elem == x {
			return
		}
		p = n
		n = n.next
	}

	node := &SLNode{Elem: x}
	node.next = n
	p.next = node

	counter := 0
	for {
		rand.Seed(time.Now().UnixNano())
		random := rand.Intn(2)
		if random == 0 {
			break
		}
		counter++
		tmp := &SLNode{Elem: x}
		tmp.bottom = node
		if counter < height {
			tmp.next = ps[counter-1].next
			ps[counter-1].next = tmp
			node = tmp
		} else {
			header := &SLNode{Elem: math.MinInt64}
			header.bottom = sl.levels[height-1]
			header.next = tmp
			sl.levels = append(sl.levels, header)
			break
		}
	}
}

func (sl *SkipList) Delete(x int) {
	height := len(sl.levels)
	p := sl.levels[height-1]
	n := p.next

	for {
		if n == nil || n.Elem > x {
			if p.bottom == nil {
				break
			}
			p = p.bottom
			n = p.next
			continue
		}
		if n.Elem == x {
			// 找到待删除的节点
			p.next = n.next
			if p.Elem == math.MinInt64 && p.next == nil {
				sl.levels = sl.levels[:len(sl.levels)-1]
			}
			if p.bottom == nil {
				break
			} else {
				// 底部还有节点就继续删除
				p = p.bottom
				n = p.next
				continue
			}
		}
		p = n
		n = n.next
	}
}
