# Trees

## Binary Tree

A binary tree is a hierarchical data structure where each node has at most 
two children, referred to as the left child and the right child.

### Key Properties:

* Root: The topmost node of the tree.
* Edge: The link from a parent node to a child node.
* Leaf: A node with no children.
* Subtree: A node and all its descendants.

### Common Tree Traversal Methods:

#### In-order Traversal:

* Visits nodes in the order: Left -> Root -> Right.
* For a binary search tree, this traversal visits nodes in ascending order.

#### Pre-order Traversal:

* Visits nodes in the order: Root -> Left -> Right.
* Useful for creating a copy of the tree or a prefix expression.

#### Post-order Traversal:

* Visits nodes in the order: Left -> Right -> Root.
* Useful for deleting a tree from the bottom up.

#### Level-order Traversal (Breadth-First Search):

* Visits nodes level by level, from left to right.
* Typically implemented using a queue.

### Common Binary Tree Types:

* Full Binary Tree: Every node has either 0 or 2 children.
* Complete Binary Tree: Every level is completely filled, except possibly the 
  last one, which is filled from left to right.
* Perfect Binary Tree: All internal nodes have two children, and all leaf 
  nodes are at the same level.
* Binary Search Tree (BST): A binary tree where for each node, the value of 
  all nodes in the left subtree is less than the node's value, and the value of 
  all nodes in the right subtree is greater.

## Tree Traversal Methods: DFS vs. BFS

Tree traversal is the process of visiting each node in a tree data structure 
exactly once. The order in which nodes is visited determines the type of 
traversal.

#### Depth-First Search (DFS)

DFS is an algorithm for traversing or searching tree or graph data structures. 
The algorithm starts at the root node and explores as far as possible along 
each branch before backtracking. There are three common ways to perform DFS 
on a tree:

* **Pre-order:** Root -> Left -> Right
* **In-order:** Left -> Root -> Right
* **Post-order:** Left -> Right -> Root

The most common implementation of DFS is **recursive**, which implicitly uses 
the call stack to manage the nodes to visit.

#### Iterative Depth-First Search

This is a non-recursive implementation of DFS that uses an **explicit stack** 
data structure. You push the root onto the stack, and then in a loop, you pop a 
node, process it, and push its children onto the stack (often right child 
first, then left, to mimic the recursive order). It achieves the same result 
as the recursive version and is useful for avoiding stack overflow issues on 
very deep trees.

#### Breadth-First Search (BFS)

BFS is another graph traversal algorithm that explores all the neighbor nodes 
at the present depth prior to moving on to the nodes at the next depth level. 
It systematically explores a tree level by level.

BFS is typically implemented with a **queue** data structure. You start by 
adding the root node to the queue. While the queue is not empty, you remove a 
node from the front, process it, and then add its children to the back of the 
queue.

### Binary Tree vs Binary Search Tree

* **Binary Tree:** Each node has at most two children. There are no specific 
  ordering rules for the nodes, and its primary use is for general hierarchical 
  data organization.
* **Binary Search Tree (BST):** Each node has at most two children, but it 
  follows a strict ordering property. All values in a node's left subtree are 
  less than the node's value, and all values in its right subtree are greater. 
  Its primary use is for efficient searching, insertion, and deletion.
