# Binary Search

In computer science, binary search, also known as half-interval search, logarithmic search, or binary chop, is a search algorithm that finds the position of a target value within a sorted array. Binary search compares the target value to the middle element of the array; if they are unequal, the half in which the target cannot lie is eliminated and the search continues on the remaining half until it is successful. If the search ends with the remaining half being empty, the target is not in the array. *- source:* [wikipedia](https://en.wikipedia.org/wiki/Binary_search_algorithm)

>[implementation](https://github.com/Samueljoli/golang-algorithms/blob/master/algorithms/searching/binarySearch/implementation.go)

![binary search animation](https://github.com/Samueljoli/golang-algorithms/blob/master/assets/binary-search.gif?raw=true)

## Complexity

**Time complexity:** `O( log n )`
