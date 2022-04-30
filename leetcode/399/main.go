package leetcode

// https://leetcode.com/problems/evaluate-division
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    disjointSet := newDisjointSet()
    set := Set{}
    result := DivisionResult{}
    for idx, equation := range equations {
        source := equation[0]
        destination := equation[1]
        result.SetResult(source, destination, values[idx])
        set.Add(source)
        set.Add(destination)
        disjointSet.Union(source, destination)
    }
    var answer []float64
    for _, query := range queries {
        source := query[0]
        destination := query[1]
        if !set.IsAllExist(source, destination) || !disjointSet.IsConnected(source, destination) {
            answer = append(answer, -1)
            continue
        }
        if val, ok := result.GetResult(source, destination); ok {
            answer = append(answer, val)
            continue
        }
        fillResult(source, result)
        val, _ := result.GetResult(source, destination)
        answer = append(answer, val)
    }
    return answer
}

func fillResult(source string, result DivisionResult) {
    var queue []string
    queue = enqueue(queue, source)
    result.SetResult(source, source, 1)
    visited := map[string]bool{}
    for len(queue) > 0 {
        current := queue[0]
        queue = dequeue(queue)
        currentResult, _ := result.GetResult(source, current)

        connectedVariables, values := result.GetConnectedVariables(current)
        for idx, variable := range connectedVariables {
            if visited[variable] {
                continue
            }
            result.SetResult(source, variable, currentResult*values[idx])
            queue = enqueue(queue, variable)
            visited[variable] = true
        }
    }
}

func enqueue(queue []string, val string) []string {
    return append(queue, val)
}

func dequeue(queue []string) []string {
    return queue[1:]
}

type Set map[string]bool

func (s Set) Add(key string) {
    s[key] = true
}

func (s Set) IsAllExist(keys ...string) bool {
    exist := true
    for _, key := range keys {
        _, ok := s[key]
        exist = exist && ok
    }
    return exist
}

type DivisionResult map[string]map[string]float64

func (r DivisionResult) GetConnectedVariables(source string) ([]string, []float64) {
    r.ensureMapExist(source)
    var destinations []string
    var values []float64
    for destination, value := range r[source] {
        destinations = append(destinations, destination)
        values = append(values, value)
    }
    return destinations, values
}

func (r DivisionResult) SetResult(source, destination string, value float64) {
    r.ensureMapExist(source)
    r.ensureMapExist(destination)
    r[source][destination] = value
    r[destination][source] = 1 / value
}

func (r DivisionResult) GetResult(source, destination string) (float64, bool) {
    r.ensureMapExist(source)
    r.ensureMapExist(destination)
    if val, ok := r[source][destination]; ok {
        return val, true
    }
    if val, ok := r[destination][source]; ok {
        return 1.0 / val, true
    }
    return -1, false
}

func (r DivisionResult) ensureMapExist(key string) {
    if _, ok := r[key]; !ok {
        r[key] = map[string]float64{}
    }
}

type DisjointSet struct {
    parents map[string]string
}

func (s *DisjointSet) IsConnected(a, b string) bool {
    return s.FindParent(a) == s.FindParent(b)
}

func (s *DisjointSet) Union(a, b string) {
    s.ensureKey(a)
    s.ensureKey(b)
    s.parents[s.FindParent(a)] = s.FindParent(b)
}

func (s *DisjointSet) FindParent(id string) string {
    s.ensureKey(id)
    if s.parents[id] != id {
        s.parents[id] = s.FindParent(s.parents[id])
    }
    return s.parents[id]
}

func (s *DisjointSet) ensureKey(key string) {
    if _, ok := s.parents[key]; !ok {
        s.parents[key] = key
    }
}

func newDisjointSet() *DisjointSet {
    return &DisjointSet{parents: map[string]string{}}
}
