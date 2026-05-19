func findCircleNum(isConnected [][]int) int {
    n := len(isConnected)
	visited := make([]bool, n)
	provinces := 0

	for city := 0; city < n; city++ {
		if !visited[city] {
			provinces++
			DFS(isConnected, city, visited)
		}
	}

	return provinces
}

func DFS(isConnected [][]int, city int, visited []bool) {
	visited[city] = true
	for neighbor := 0; neighbor < len(isConnected); neighbor++ {
		if isConnected[city][neighbor] == 1 && !visited[neighbor] {
			DFS(isConnected, neighbor, visited)
		}
	}
}