/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
   if node == nil {
        return nil
    }

    visited := make(map[*Node]*Node)

    var dfs func(*Node) *Node

    dfs = func(curr *Node) *Node {

        if clone, exists := visited[curr]; exists {
            return clone
        }

        clone := &Node{
            Val: curr.Val,
        }

        visited[curr] = clone

        for _, nei := range curr.Neighbors {
            clone.Neighbors = append(clone.Neighbors, dfs(nei))
        }

        return clone
    }

    return dfs(node) 
}