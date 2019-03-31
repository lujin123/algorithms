package dsa

import (
	"math"
)

type MinHeap struct {
	size     int64
	cap      int64
	elements []int64
}

// 最小堆的插入操作（上滤）：先将元素加载最后，然后和父节点逐个比较，比父节点小就交换位置，知道根节点
func (h *MinHeap) Insert(element int64) {
	h.elements = append(h.elements, element)
	h.size++
	for i := h.size - 1; h.elements[i/2] > h.elements[i]; i /= 2 {
		h.elements[i], h.elements[i/2] = h.elements[i/2], h.elements[i]
	}
}

// 最小堆删除操作：将最后一个元素放到根节点，之后逐个和左右子节点比较交换位置，直到不大于子节点为止
// 如果是链表实现的，找到最后一个位置比较麻烦，可能需要加一个最后一个元素的指针之类的东西
func (h *MinHeap) Delete() (ans int64) {
	ans = h.elements[1]
	h.elements[1] = h.elements[h.size-1]
	h.size--
	var i int64
	for i = 2; i < h.size; i *= 2 {
		if h.elements[i/2] > h.elements[i] {
			h.elements[i/2], h.elements[i] = h.elements[i], h.elements[i/2]
		} else if h.elements[i/2] > h.elements[i+1] {
			h.elements[i/2], h.elements[i+1] = h.elements[i+1], h.elements[i/2]
		} else {
			break
		}
	}
	return
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		elements: []int64{math.MinInt64},
		cap:      math.MaxInt64,
		size:     1,
	}
}
