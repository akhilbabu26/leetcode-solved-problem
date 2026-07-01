func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)

	dist := nearestThief(grid)

	lo, hi := 0, 2*n

	ans := 0

	for lo <= hi {
		mid := (lo + hi) / 2

		if canReach(dist, mid) {
			ans = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return ans
}

func canReach(dist [][]int, limit int) bool {
	n := len(dist)

	if dist[0][0] < limit {
		return false
	}

	dir := [][]int{{1,0},{-1,0},{0,1},{0,-1}}

	vis := make([][]bool,n)
	for i:=range vis{
		vis[i]=make([]bool,n)
	}

	q := [][]int{{0,0}}
	vis[0][0]=true

	for len(q)>0{
		x,y:=q[0][0],q[0][1]
		q=q[1:]

		if x==n-1 && y==n-1{
			return true
		}

		for _,d:=range dir{
			nx:=x+d[0]
			ny:=y+d[1]

			if nx<0||ny<0||nx>=n||ny>=n{
				continue
			}

			if vis[nx][ny] || dist[nx][ny]<limit{
				continue
			}

			vis[nx][ny]=true
			q=append(q,[]int{nx,ny})
		}
	}

	return false
}

func nearestThief(grid [][]int) [][]int {
	n := len(grid)

	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}

	type Pair struct {
		x, y int
	}

	q := []Pair{}

	// Multi-source BFS: start from all thieves.
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dist[i][j] = 0
				q = append(q, Pair{i, j})
			}
		}
	}

	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for head := 0; head < len(q); head++ {
		cur := q[head]

		for _, d := range dirs {
			nx := cur.x + d[0]
			ny := cur.y + d[1]

			if nx < 0 || ny < 0 || nx >= n || ny >= n {
				continue
			}

			if dist[nx][ny] != -1 {
				continue
			}

			dist[nx][ny] = dist[cur.x][cur.y] + 1
			q = append(q, Pair{nx, ny})
		}
	}

	return dist
}