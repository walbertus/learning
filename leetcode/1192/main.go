package leetcode

import (
    "math"
    "time"
)

// https://leetcode.com/problems/critical-connections-in-a-network
func criticalConnections(n int, connections [][]int) [][]int {
    visited := make([]bool, n)
    parent := make([]int, n)
    lowestDiscoveryTime := make([]int64, n)
    discoveryTime := make([]int64, n)
    for i := 0; i < n; i++ {
        visited[i] = false
        parent[i] = -1
        lowestDiscoveryTime[i] = math.MaxInt
        discoveryTime[i] = math.MaxInt
    }

    adjList := make([][]int, n)
    for _, connection := range connections {
        adjList[connection[0]] = append(adjList[connection[0]], connection[1])
        adjList[connection[1]] = append(adjList[connection[1]], connection[0])
    }

    initialNode := 0
    visited[initialNode] = true
    parent[initialNode] = initialNode
    now := getCurrentNanoTime()
    lowestDiscoveryTime[initialNode] = now
    discoveryTime[initialNode] = now

    return traverseNode(initialNode, adjList, visited, parent, lowestDiscoveryTime, discoveryTime)
}

func traverseNode(node int, graph [][]int, visited []bool, parent []int, lowestDiscoveryTime []int64, discoveryTime []int64) [][]int {
    visited[node] = true
    now := getCurrentNanoTime()
    discoveryTime[node] = now
    lowestDiscoveryTime[node] = now
    var criticalBridges [][]int
    for _, neighbor := range graph[node] {
        if visited[neighbor] {
            if parent[node] != neighbor {
                lowestDiscoveryTime[node] = min(lowestDiscoveryTime[node], discoveryTime[neighbor])
            }
            continue
        }
        parent[neighbor] = node
        nextCriticalBridges := traverseNode(neighbor, graph, visited, parent, lowestDiscoveryTime, discoveryTime)
        lowestDiscoveryTime[node] = min(lowestDiscoveryTime[node], lowestDiscoveryTime[neighbor])
        if lowestDiscoveryTime[neighbor] > discoveryTime[node] {
            criticalBridges = append(criticalBridges, []int{node, neighbor})
        }
        criticalBridges = append(criticalBridges, nextCriticalBridges...)
    }
    return criticalBridges
}

func getCurrentNanoTime() int64 {
    return time.Now().UnixNano()
}

func min(a, b int64) int64 {
    if a < b {
        return a
    }
    return b
}
