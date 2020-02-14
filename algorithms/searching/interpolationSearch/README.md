# Interpolation Search

**Interpolation search** is an algorithm for searching for a key in an array that has been ordered by numerical values assigned to the keys (key values).

For example we have a sorted array of n uniformly distributed values arr[], and we need to write a function to search for a particular element x in the array.

**Linear Search** finds the element in O(n) time, **Jump Search** takes O(√ n) time and **Binary Search** take O(Log n) time.

The **Interpolation Search** is an improvement over Binary Search for instances, where the values in a sorted array are uniformly distributed.
Binary Search always goes to the middle element to check. On the other hand, interpolation search may go to different locations according to the value of
the key being searched. For example, if the value of the key is closer to the last element, interpolation search is likely to start search toward the end side. *- source:* [GeeksForGeeks](https://www.geeksforgeeks.org/interpolation-search/), [Wikipedia](https://en.wikipedia.org/wiki/Interpolation_search)
> [implementation](https://github.com/Samueljoli/golang-algorithms/blob/master/algorithms/searching/interpolationSearch/implementation.go)

## Complexity
**Time complexity:** `O (log(log(n))`

