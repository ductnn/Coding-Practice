# https://leetcode.com/problems/maximize-the-number-of-target-nodes-after-connecting-trees-i/description/

from typing import List

class Solution:
    def dfs(self, graph, node, parent, depth):
        if depth < 0:
            return 0
        count = 1
        for neighbor in graph[node]:
            if neighbor != parent:
                count += self.dfs(graph, neighbor, node, depth - 1)
        return count

    def maxTargetNodes(self, edges1: List[List[int]], edges2: List[List[int]], k: int) -> List[int]:
        # Find the number of nodes by getting the maximum node index + 1
        n1 = max(max(edge[0], edge[1]) for edge in edges1) + 1 if edges1 else 0
        n2 = max(max(edge[0], edge[1]) for edge in edges2) + 1 if edges2 else 0

        graph1 = [[] for _ in range(n1)]
        graph2 = [[] for _ in range(n2)]

        for edge in edges1:
            graph1[edge[0]].append(edge[1])
            graph1[edge[1]].append(edge[0])
        
        for edge in edges2:
            graph2[edge[0]].append(edge[1])
            graph2[edge[1]].append(edge[0])
        
        max_targets_tree2 = 0
        for i in range(n2):
            reachable = self.dfs(graph2, i, -1, k - 1)
            max_targets_tree2 = max(max_targets_tree2, reachable)
        
        result = []
        for i in range(n1):
            reachable_in_tree1 = self.dfs(graph1, i, -1, k)
            total = reachable_in_tree1 + max_targets_tree2
            result.append(total)
        return result

# === MAIN TEST CASE ===
if __name__ == "__main__":
    sol = Solution()
    edges1 = [[0,1],[0,2],[2,3],[2,4]]
    edges2 = [[0,1],[0,2],[0,3],[2,7],[1,4],[4,5],[4,6]]
    k = 2
    print(sol.maxTargetNodes(edges1, edges2, k))
