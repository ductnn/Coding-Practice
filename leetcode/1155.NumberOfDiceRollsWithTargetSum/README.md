## Solutions

**Solution: Dynamic Programming**

We define $f[i][j]$ as the number of ways to get a sum of $j$ using $i$ dice.
Then, we can obtain the following state transition equation:

$$
f[i][j] = \sum_{h=1}^{\min(j, k)} f[i-1][j-h]
$$

where $h$ represents the number of points on the $i$-th die.

Initially, we have $f[0][0] = 1$, and the final answer is $f[n][target]$.

The time complexity is $O(n \times k \times target)$, and the space complexity
is $O(n \times target)$.

We notice that the state $f[i][j]$ only depends on $f[i-1][]$, so we can use a
rolling array to optimize the space complexity to $O(target)$.
