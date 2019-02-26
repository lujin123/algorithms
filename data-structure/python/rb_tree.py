# ref
# 1. http://www.cnblogs.com/yangecnu/p/Introduce-Red-Black-Tree.html


class Node(object):
    def __init__(self, key, value):
        self.left = None
        self.right = None
        self.key = key
        self.value = value
        self.red = True


class RBTree(object):
    def __init__(self, root=None):
        self._root = root
        self._parent = None

    def _new_node(self, key, value):
        return Node(key, value)

    def _is_red(self, node):
        return node.red if node else False

    def rotate_left(self, node):
        root = node.right
        node.right = root.left
        root.left = node

        root.red = node.red
        node.red = True
        return root

    def rotate_right(self, node):
        root = node.left
        node.left = root.right
        root.right = node

        root.red = node.red
        node.red = True
        return root

    def flip_color(self, node):
        node.red = True
        node.left.red = False
        node.right.red = False
        return node

    def insert(self, key, value):
        self._root = self._insert(self._root, key, value)
        self._root.red = False

    def _insert(self, root, key, value):
        """
        1. 如果节点的右子节点为红色，且左子节点为黑色，则进行左旋操作
        2. 如果节点的左子节点为红色，并且左子节点的左子节点也为红色，则进行右旋操作
        3. 如果节点的左右子节点均为红色，则执行FlipColor操作，提升中间结点
        :param root:
        :param key:
        :param value:
        :return:
        """
        if not root:
            return self._new_node(key, value)
        if key < root.key:
            root.left = self._insert(root.left, key, value)
        elif key > root.key:
            root.right = self._insert(root.right, key, value)
        else:
            root.value = value
        if self._is_red(root.right) and not self._is_red(root.left):
            root = self.rotate_left(root)
        if self._is_red(root.left) and self._is_red(root.left.left):
            root = self.rotate_right(root)
        if self._is_red(root.left) and self._is_red(root.right):
            root = self.flip_color(root)
        return root

    def transplant(self, node_p, node_s):
        if node_s.left and node_s.right:
            node_v = self.min_node(node_s.right)
            node_v.left = node_s.left

        else:
            node_v = node_s.left if node_s.left else node_s.right

        if not node_p:
            self._root = node_v
        elif node_p.left.key == node_s.key:
            node_p.left = node_v
        elif node_p.right.key == node_s.key:
            node_p.right = node_v

    def find(self, key):
        node = self._root
        self._parent = None
        while node:
            if node.key > key:
                self._parent = node
                node = node.left
            elif node.key < key:
                self._parent = node
                node = node.right
            else:
                return node
        self._parent = None
        return None

    def remove(self, key):
        """
        x: 待删除节点
        w: x的兄弟节点
        p: x的父节点

        1. x的兄弟节点w是red
         w的两个子节点必定是black,这样可以交换一下 p 和 w 的颜色，然后对p节点进行左旋转

        2. x的兄弟节点w是black并且w的两个孩子也都是black

        :param key:
        :return:
        """
        node = self.find(key)
        parent = self._parent
        if not node:
            return
        if node.left and node.right:
            pass
        else:
            tmp_node = node.left if node.left else node.right

            pass

    def _remove(self, root, key):
        if not root:
            raise KeyError("Error, data not in tree")
        if key > root.key:
            pass

    @staticmethod
    def min_node(node):
        tmp_node = node
        while tmp_node.left:
            tmp_node = tmp_node.left
        return tmp_node

    @staticmethod
    def max_node(node):
        tmp_node = node
        while tmp_node.right:
            tmp_node = tmp_node.right
        return tmp_node


def main():
    # m = [1, 2, 9, 8, 3, 6]
    m = [1, 2, 3, 6, 8, 9]
    rb_tree = RBTree()
    for k in m:
        rb_tree.insert(k, None)
    # for k, v in m.items():
    #     rb_tree.insert(k, v)

    # print(rb_tree.find(99))
    print(rb_tree)
    pass

if __name__ == '__main__':
    main()
