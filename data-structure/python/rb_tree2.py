"""
1. Every node is either red or black.
2. The root is black.
3. Every leaf ( NIL ) is black.
4. If a node is red, then both its children are black.
5. For each node, all simple paths from the node to descendant leaves contain the same number of black nodes.
"""
RED = 0
BLACK = 1


class Node(object):
    def __init__(self, key, value, color=RED):
        self.key = key
        self.value = value
        self.left = None
        self.right = None
        self.parent = None
        self.color = color


class RBTree(object):
    def __init__(self, root=None):
        self._root = root

    def _new_node(self, key, value):
        return Node(key, value)

    def rotate_left(self, node):
        root = node.right

        node.right = root.left
        if root.left:
            root.left.parent = node

        root.parent = node.parent
        if not node.parent:
            self._root = root
        elif node == node.parent.left:
            node.parent.left = root
        else:
            node.parent.right = root

        root.left = node
        node.parent = root

    def rotate_right(self, node):
        root = node.left

        node.left = root.right
        if root.right:
            root.right.parent = node

        root.parent = node.parent
        if not node.parent:
            self._root = root
        elif node == node.parent.left:
            node.parent.left = root
        else:
            node.parent.right = root

        root.right = node
        node.parent = root

    # def _flip_color(self, node):
    #     node.parent.color = BLACK

    def _find_pos(self, key):
        node = None
        tmp = self._root
        while tmp:
            node = tmp
            tmp = tmp.right if key > tmp.key else tmp.left
        return node

    def insert(self, key, value):
        parent = self._find_pos(key)
        node = self._new_node(key, value)
        node.parent = parent
        if not parent:
            self._root = node
        elif key > parent.key:
            parent.right = node
        else:
            parent.left = node
        self._insert_fix_up(node)

    def _insert_fix_up(self, node):
        while node.parent and node.parent.color == RED:
            if node.parent == node.parent.parent.left:
                uncle = node.parent.parent.right
                if uncle and uncle.color == RED:
                    node.parent.color = BLACK
                    uncle.color = BLACK
                    node.parent.parent.color = RED
                    node = node.parent.parent
                    continue
                elif node == node.parent.right:
                    self.rotate_left(node.parent)
                    node = node.left
                node.parent.color = BLACK
                node.parent.parent.color = RED
                self.rotate_right(node.parent.parent)
            elif node.parent == node.parent.parent.right:
                uncle = node.parent.parent.left
                if uncle and uncle.color == RED:
                    node.parent.color = BLACK
                    uncle.color = BLACK
                    node.parent.parent.color = RED
                    node = node.parent.parent
                    continue
                elif node == node.parent.left:
                    # node = node.parent
                    self.rotate_right(node.parent)
                    node = node.right
                node.parent.color = BLACK
                node.parent.parent.color = RED
                self.rotate_left(node.parent.parent)

        self._root.color = BLACK


def main():
    m = {1: 'a', 2: 'b', 9: 'c', 8: 'd', 3: 'e', 6: 'f'}
    rb_tree = RBTree()
    for k, v in m.items():
        rb_tree.insert(k, v)

    print(rb_tree)
    pass

if __name__ == '__main__':
    main()
