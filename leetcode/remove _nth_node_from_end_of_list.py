# Created by lujin at 2/3/2017
#
# Remove Nth Node From End of List
#
# Description:
# Given a linked list, remove the nth node from the end of list and return its head.
#
# For example,
#
#    Given linked list: 1->2->3->4->5, and n = 2.
#
#    After removing the second node from the end, the linked list becomes 1->2->3->5.
# Note:
# Given n will always be valid.
# Try to do this in one pass.
#


class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def removeNthFromEnd(self, head, n):
        """
        :type head: ListNode
        :type n: int
        :rtype: ListNode
        """
        if not head:
            return head

        p1 = head
        p2 = None
        while p1:
            if n <= 0:
                p2 = p2.next if p2 else head
            else:
                n -= 1
            p1 = p1.next
        if not p2:
            head = head.next
        elif p2.next:
            p2.next = p2.next.next
        else:
            head = None
        return head

    def test(self):
        a = [1, 2, 3, 4, 5]
        # a = [1]
        # a = [1, 2]
        head = None
        p = head
        for i in a:
            node = ListNode(i)
            if not head:
                head = node
            else:
                p.next = node
            p = node
        res = self.removeNthFromEnd(head, 2)
        return res


if __name__ == '__main__':
    s = Solution()
    s.test()
