package leetcode

// https://leetcode-cn.com/problems/rotate-list/
/**
61. 旋转链表

给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL
示例 2:

输入: 0->1->2->NULL, k = 4
输出: 2->0->1->NULL
解释:
向右旋转 1 步: 2->0->1->NULL
向右旋转 2 步: 1->2->0->NULL
向右旋转 3 步: 0->1->2->NULL
向右旋转 4 步: 2->0->1->NULL
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//  思路：首先统计下链表长度n，之后计算实际需要移动的位置k=k%n，
// 将链表头尾相连成为了循环链表，从头开始循环n-k次找到节点断开作为头结点，
// 前一个节点末尾置空即可
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	n := 1
	cursor := head

	for cursor.Next != nil {
		cursor = cursor.Next
		n++
	}

	loop := n - (k % n)
	tail := cursor
	cursor.Next = head
	cursor = head

	for i := 0; i < loop; i++ {
		cursor = cursor.Next
		tail = tail.Next
	}
	tail.Next = nil
	return cursor
}

func rotateRight2(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	n := 1
	cursor := head

	for cursor.Next != nil {
		cursor = cursor.Next
		n++
	}

	cursor.Next = head
	loop := n - (k % n)

	cursor = head

	for i := 0; i < loop-1; i++ {
		cursor = cursor.Next
	}
	head = cursor.Next
	cursor.Next = nil

	return head
}
