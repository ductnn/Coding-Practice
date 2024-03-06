# [1750. Minimum Length of String After Deleting Similar Ends](https://leetcode.com/problems/minimum-length-of-string-after-deleting-similar-ends)

## Solutions

### Solution 1: Two pointers

Define two pointers $i$ and $j$ to point to the head and tail of the string
$s$ respectively, then move them to the middle until the characters pointed to
by $i$ and $j$ are not equal

=> $\max(0, j - i + 1)$ is the answer.

The time complexity is $O(n)$ and the space complexity is $O(1)$. Where $n$ is
the length of the string $s$.
