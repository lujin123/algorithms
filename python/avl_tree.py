# ref
# 1. http://blog.chinaunix.net/uid-25324849-id-2182877.html
# 2. https://segmentfault.com/a/1190000007054898
# 3. http://www.cnblogs.com/wuditju/p/5999411.html
# 4. http://blog.csdn.net/sysu_arui/article/details/7897017


class Node(object):
    def __init__(self, data=None):
        self.left = None
        self.right = None
        self.data = data
        self.height = 0


class AVLTree(object):
    def __init__(self, root=None):
        self._root = root
        self._count = 0

    def _new_node(self, data=None):
        self._count += 1
        return Node(data)

    @staticmethod
    def height(node):
        return node.height if node else -1

    @property
    def root(self):
        return self._root

    @property
    def count(self):
        return self._count

    def _max_height(self, node):
        return max(self.height(node.left), self.height(node.right))+1

    def left_left_rotate(self, node):
        """
        1. 对结点的左儿子的左子树进行一次插入，需要进行一次右旋
        （左旋还是右旋不知道该怎么叫，其实是以这个结点的左子节点为中心顺时针的旋转一下...）
            3
           /     右旋         2
          2    -------->    / \
         /                 1   3
        1
        新增一个3就需要像上面一样旋转一下
        :param node:
        :return:
        """
        root = node.left
        node.left = root.right
        root.right = node
        node.height = self._max_height(node)
        root.height = self._max_height(root)
        return root

    def right_right_rotate(self, node):
        """
        2. 对结点的右儿子的右子树进行一次插入，需要进行一次左旋
            1
            \         左旋       2
             2     -------->   / \
              \               1   3
               3
        :param node:
        :return:
        """
        root = node.right
        node.right = root.left
        root.left = node
        node.height = self._max_height(node)
        root.height = self._max_height(root)
        return root

    def left_right_rotate(self, node):
        """
        3. 对左儿子的右子树进行一次插入，这个时候需要选先进行一次左旋，再进行一次右旋
                10                       10
               / \                      /  \                   6
              5   19       左旋         6   19     右旋        /  \
             / \        ---------->   / \      --------->   5   10
            2   6                    5   7                 /    / \
                 \                  /                     2    7   19
                  7                2
        7为新加入的结点
        :param node:
        :return:
        """
        node.left = self.right_right_rotate(node.left)
        return self.left_left_rotate(node)

    def right_left_rotate(self, node):
        """
        4. 对右儿子的左子树进行一次插入，这个时候需要先进行一次右旋，再进行一次左旋
                 6                    6
               /  \                  / \                   9
              5   10      右旋       5  9        左旋       / \
                 /  \  --------->     / \     -------->   6  10
                9   11               8   10              / \  \
               /                          \             5  8  11
              8                           11
        8为新加入的结点
        :param node:
        :return:
        """
        node.right = self.left_left_rotate(node.right)
        return self.right_right_rotate(node)

    def insert(self, data):
        node = self._new_node(data)
        if self._root is None:
            self._root = node
        else:
            self._root = self._add_node(self._root, node)

    def _add_node(self, root, node):
        if not root:
            root = node
        elif node.data < root.data:
            root.left = self._add_node(root.left, node)
            if self.height(root.left) - self.height(root.right) >= 2:
                root = self.left_left_rotate(root) if node.data < root.left.data else self.left_right_rotate(root)
        elif node.data > root.data:
            root.right = self._add_node(root.right, node)
            if self.height(root.right) - self.height(root.left) >= 2:
                root = self.right_left_rotate(root) if node.data < root.right.data else self.right_right_rotate(root)
        root.height = self._max_height(root)
        return root

    def remove(self, data):
        self._root = self._remove(self._root, data)

    def _remove(self, root, data):
        """
        删除的步骤：
        1. 如果待删除的结点是一个叶子结点，直接删除即可
        2. 如果待删除的结点有一个子树，直接用用这个子树代替需要删除的结点即可
        3. 如果左右子树都存在，那么就查找到左子树的最大结点或者右子树的最小结点，用这个结点替换需要删除的结点，
        需要注意的是删除找到的结点的删除需要递归的删除，因为还是需要平衡树的高度
        :param root:
        :param data:
        :return:
        """
        if not root:
            raise KeyError("Error, data not in tree")
        elif data < root.data:
            root.left = self._remove(root.left, data)
            if self.height(root.right) - self.height(root.left) >= 2:
                if self.height(root.right.right) > self.height(root.right.left):
                    root = self.right_right_rotate(root)
                else:
                    root = self.right_left_rotate(root)
            root.height = self._max_height(root)
        elif data > root.data:
            root.right = self._remove(root.right, data)
            if self.height(root.left) - self.height(root.right) >= 2:
                if self.height(root.left.left) > self.height(root.left.right):
                    root = self.left_left_rotate(root)
                else:
                    root = self.left_right_rotate(root)
            root.height = self._max_height(root)
        elif root.left and root.right:
            if root.left.height <= root.right.height:
                min_node = self.tree_min(root.right)
                root.data = min_node.data
                root.right = self._remove(root.right, root.data)
            else:
                max_node = self.tree_max(root.left)
                root.data = max_node.data
                root.left = self._remove(root.left, root.data)
            root.height = self._max_height(root)
        else:
            root = root.right if root.right else root.left
        return root

    @staticmethod
    def tree_min(root):
        tmp_node = root
        while tmp_node.left:
            tmp_node = tmp_node.left
        return tmp_node

    @staticmethod
    def tree_max(root):
        tmp_node = root
        while tmp_node:
            tmp_node = tmp_node.right
        return tmp_node

    def pre_order(self, root):
        if root:
            print(root.data, end=" ")
            self.pre_order(root.left)
            self.pre_order(root.right)


def main():
    avl = AVLTree()
    for i in [1, 2, 9, 8, 3, 6]:
        avl.insert(i)

    print('\n========before remove==========')
    avl.pre_order(avl.root)
    print('\n========after remove==========')
    # avl.remove(0)
    avl.remove(3)
    avl.pre_order(avl.root)

if __name__ == '__main__':
    main()
    # 运行结果如下：
    # == == == == before remove == == == == ==
    # 3 2 1 8 6 9
    # == == == == after remove == == == == ==
    # 6 2 1 8 9

