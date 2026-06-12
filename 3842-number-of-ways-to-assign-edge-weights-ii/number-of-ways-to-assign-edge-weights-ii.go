const MOD = 1000000007

type Solution struct{}

func assignEdgeWeights(edges [][]int, queries [][]int) []int {
	n := len(edges) + 1

	graph := make([][]int, n+1)
	for _, e := range edges {
		u, v := e[0], e[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	LOG := 17
	for (1 << LOG) <= n {
		LOG++
	}

	depth := make([]int, n+1)
	parent := make([][]int, LOG)
	for i := range parent {
		parent[i] = make([]int, n+1)
	}

	var dfs func(int, int)
	dfs = func(u, p int) {
		parent[0][u] = p

		for _, v := range graph[u] {
			if v == p {
				continue
			}
			depth[v] = depth[u] + 1
			dfs(v, u)
		}
	}

	dfs(1, 0)

	for k := 1; k < LOG; k++ {
		for v := 1; v <= n; v++ {
			if parent[k-1][v] != 0 {
				parent[k][v] = parent[k-1][parent[k-1][v]]
			}
		}
	}

	lca := func(a, b int) int {
		if depth[a] < depth[b] {
			a, b = b, a
		}

		diff := depth[a] - depth[b]

		for k := 0; k < LOG; k++ {
			if (diff & (1 << k)) != 0 {
				a = parent[k][a]
			}
		}

		if a == b {
			return a
		}

		for k := LOG - 1; k >= 0; k-- {
			if parent[k][a] != parent[k][b] {
				a = parent[k][a]
				b = parent[k][b]
			}
		}

		return parent[0][a]
	}

	pow2 := make([]int, n)
	pow2[0] = 1
	for i := 1; i < n; i++ {
		pow2[i] = (pow2[i-1] * 2) % MOD
	}

	ans := make([]int, len(queries))

	for i, q := range queries {
		u, v := q[0], q[1]

		p := lca(u, v)
		dist := depth[u] + depth[v] - 2*depth[p]

		if dist == 0 {
			ans[i] = 0
		} else {
			ans[i] = pow2[dist-1]
		}
	}

	return ans
}