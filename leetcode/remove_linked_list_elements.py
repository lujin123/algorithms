# Created by lujin at 3/3/
#
# 203. Remove Linked List Elements
#
# Description:
#
# Remove all elements from a linked list of integers that have value val.
#
# Example
# Given: 1 --> 2 --> 6 --> 3 --> 4 --> 5 --> 6, val = 6
# Return: 1 --> 2 --> 3 --> 4 --> 5
#


class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def removeElements(self, head, val):
        """
        删除链表中节点值为指定元素的节点

        :type head: ListNode
        :type val: int
        :rtype: ListNode
        """
        while head and head.val == val:
            head = head.next

        if not head:
            return head

        p = head

        while p.next:
            if p.next.val == val:
                p.next = p.next.next
            else:
                p = p.next

        return head

    def test(self):
        a = [1, 2, 6, 3, 4, 6, 1]
        # a = [6, 6, 1]
        # a = [6]
        # a = []
        # a = [1, 6, 6, 1]
        head = None
        p = None
        for i in a:
            node = ListNode(i)
            if head:
                p.next = node
            else:
                head = node
            p = node
        self.removeElements(head, 6)

if __name__ == '__main__':
    Solution().test()
