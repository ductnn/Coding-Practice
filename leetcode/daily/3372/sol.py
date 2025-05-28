# https://leetcode.com/problems/maximize-the-number-of-target-nodes-after-connecting-trees-i/description/

def build_graph(edges, n):
    graph = [[] for _ in range(n)]
    for a, b in edges:
        graph[a].append(b)
        graph[b].append(a)
    return graph

def dfs(graph, node, parent, depth):
    if depth < 0:
        return 0
    count = 1  # Đếm chính nó
    for neighbor in graph[node]:
        if neighbor != parent:
            count += dfs(graph, neighbor, node, depth - 1)
    return count

def max_target_nodes(edges1, edges2, k):
    n1 = len(edges1) + 1
    n2 = len(edges2) + 1
    graph1 = build_graph(edges1, n1)
    graph2 = build_graph(edges2, n2)

    # Bước 1: Tìm số lượng lớn nhất các nút có thể reach được trong cây 2 với (k - 1) bước
    max_targets_tree2 = 0
    for i in range(n2):
        reachable = dfs(graph2, i, -1, k - 1)
        max_targets_tree2 = max(max_targets_tree2, reachable)

    # Bước 2: Với mỗi nút trong cây 1, tính tổng số target nodes
    result = []
    for i in range(n1):
        reachable_in_tree1 = dfs(graph1, i, -1, k)
        total = reachable_in_tree1 + max_targets_tree2
        result.append(total)
    return result

# === MAIN TEST CASE ===
if __name__ == "__main__":
    edges1 = [[0,1],[0,2],[2,3],[2,4]]
    edges2 = [[0,1],[0,2],[0,3],[2,7],[1,4],[4,5],[4,6]]
    k = 2

    output = max_target_nodes(edges1, edges2, k)
    print("Output:", output)  # Expected: [9, 7, 9, 8, 8]
