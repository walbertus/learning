package leetcode

// https://leetcode.com/problems/coin-change
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dist := []int{}
	for i := 0; i <= 10000; i++ {
		dist = append(dist, -1)
	}
	queue := Queue{}

	dist[amount] = 0
	queue = enqueue(queue, amount)

	for size(queue) != 0 {
		now := 0
		now, queue = dequeue(queue)

		for _, coin := range coins {
			next := now - coin
			if next == 0 {
				return dist[now] + 1
			}
			if next < 0 {
				continue
			}
			if dist[next] != -1 {
				continue
			}
			dist[next] = dist[now] + 1
			queue = enqueue(queue, next)
		}
	}
	return -1
}

type Queue []int

func enqueue(q Queue, value int) Queue {
	return append(q, value)
}

func dequeue(q Queue) (int, Queue) {
	result := q[0]
	return result, q[1:]
}

func size(q Queue) int {
	return len(q)
}
