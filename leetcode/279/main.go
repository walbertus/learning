package leetcode

// https://leetcode.com/problems/perfect-squares/
func numSquares(n int) int {
    q := newQueue()
    q.enqueue(0)
    distance := make(map[int]int)
    distance[0] = 0
    found := false
    for q.length() > 0 && !found {
        now := q.dequeue()
        currentDistance := distance[now]
        for i := 1; i <= 100; i++ {
            next := now + (i * i)
            if next > n {
                continue
            }
            if _, ok := distance[next]; !ok {
                distance[next] = currentDistance + 1
                q.enqueue(next)
            }
            if next == n {
                found = true
                break
            }
        }
    }
    return distance[n]
}

type queue struct {
    container []int
}

func (q *queue) enqueue(element int) {
    q.container = append(q.container, element)
}

func (q *queue) dequeue() int {
    element := q.container[0]
    q.container = q.container[1:]
    return element
}

func (q *queue) length() int {
    return len(q.container)
}

func newQueue() *queue {
    return &queue{
        container: []int{},
    }
}
