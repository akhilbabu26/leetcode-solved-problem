func validPath(n int, edges [][]int, source int, destination int) bool {
	graph := make(map[int][]int)

	for _, edge := range edges {
		a := edge[0]
		b := edge[1]

		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	visited := make(map[int]bool)

	return dfs(graph, source, destination, visited)
}

func dfs(graph map[int][]int, current int, destination int, visited map[int]bool) bool {

	if current == destination {
		return true
	}

	visited[current] = true

	for _, neighbor := range graph[current] {

		if !visited[neighbor] {

			if dfs(graph, neighbor, destination, visited) {
				return true
			}
		}
	}

	return false
}