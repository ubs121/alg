from tree_bst import BstNode

# AA-Tree balancing: right rotation
def aa_skew(node):
    if None in [node, node.left]: return node
    if node.left.lvl != node.lvl: return node  # no need skew

    lft = node.left
    node.left = lft.right
    lft.right = node
    return lft

# AA-Tree balancing: Left rotation & level increase
def aa_split(node):
    if None in [node, node.right, node.right.right]: return node
    if node.right.right.lvl != node.lvl: return node

    rgt = node.right
    node.right = rgt.left
    rgt.left = node
    rgt.lvl += 1
    return rgt

# AA-Tree balancing: insert a node
def aa_insert(node, key, val):
    if node is None: 
        return BstNode(key, val)

    if node.key == key: node.val = val
    elif key < node.key: node.left = aa_insert(node.left, key, val)
    else: node.right = aa_insert(node.right, key, val)

    node = aa_skew(node)  # in case it's backward
    node = aa_split(node)  # in case it's overfull
    return node

tree=None
tree=aa_insert(tree, "d","")
tree=aa_insert(tree, "e","")
tree=aa_insert(tree, "f","")
tree=aa_insert(tree, "a","")
tree=aa_insert(tree, "b","")
tree=aa_insert(tree, "c","")

print(tree)