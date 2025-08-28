# Invert Binary Tree

## Big O Complexity of Invert Tree

The InvertTree solution we implemented is a recursive algorithm that 
traverses every node of the binary tree exactly once.

### Time Complexity: O(n)

* The time complexity is O(n), where n is the total number of nodes in the 
  tree.
* This is because the algorithm visits and performs a constant amount of 
  work (swapping children) for each node.
* The total time is directly proportional to the number of nodes. Whether the 
  tree is a balanced tree or a skewed "linked list," every node must be visited 
  to invert the tree completely.

### Space Complexity: O(h)

* The space complexity is O(h), where h is the height of the tree.
* This space is used by the call stack due to the recursive nature of the 
  solution.
* The call stack stores the function calls as the algorithm goes deeper into 
  the tree.
* In the worst-case scenario, for a completely unbalanced (skewed) tree, the 
  height h is equal to the number of nodes n. Therefore, the worst-case space 
  complexity is O(n).
* In the best-case scenario (a perfectly balanced tree), the height is 
  log 2 (n), so the space complexity is O(logn).

