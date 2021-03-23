

 * `Binary Heap` - array representation of a heap-ordered complete binary tree. 

    **Representation**: No links required, parent of node 'k' is at k/2, children of 'k' are at 2k and 2k+1. Keys in nodes, parent's key no smaller than children's keys. 

    **Operations**: 
    * remove violations: swim or promotion, sink or demotion
    * insert O(log(n)): add node at the end and swim, 
    * del max O(log(n)): exchange with the last element and sink down
    * max O(1): first element

data/pairs - https://www.hackerrank.com/challenges/array-pairs/problem
num/max_subarray.go
weekOfCode/code36/cut