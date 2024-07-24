package main

import (
	"container/list"
	"fmt"
)

type Edge struct {
	A  string
	B  string
	Wt float64
	// A -> B ie A is X times B
}

type Ratio struct {
	A string
	R float64
}

func solve(adjMap map[string][]Edge, sourceNode, targetNode string) float64 {
	if sourceNode == targetNode {
		_, ok := adjMap[sourceNode]
		if ok {
			return 1
		} else {
			return -1
		}
	}
	q := list.New()
	mp := make(map[string]bool)
	q.PushBack(Ratio{A: sourceNode, R: 1})

	for q.Len() > 0 {
		sz := q.Len()
		for ; sz > 0; sz-- {
			el := q.Remove(q.Front()).(Ratio)
			mp[el.A] = true
			listA := adjMap[el.A]
			for _, val := range listA {
				if !mp[val.B] {
					newRatio := el.R * val.Wt
					q.PushBack(Ratio{A: val.B, R: newRatio})
					if val.B == targetNode {
						return newRatio
					}
				}
			}
		}
	}

	return -1.0
}

func makegraph(equations [][]string, values []float64) map[string][]Edge {
	mp := make(map[string][]Edge)

	for i, val := range equations {
		mp[val[0]] = append(mp[val[0]], Edge{A: val[0], B: val[1], Wt: values[i]})
		mp[val[1]] = append(mp[val[1]], Edge{A: val[1], B: val[0], Wt: 1 / values[i]})
	}
	return mp
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	mp := makegraph(equations, values)
	ans := make([]float64, len(queries))

	for i, v := range queries {
		ans[i] = solve(mp, v[0], v[1])
	}
	return ans
}

func main() {
	equations, values, queries := [][]string{{"a", "b"}, {"b", "c"}}, []float64{2.0, 3.0}, [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
	ans := calcEquation(equations, values, queries)
	fmt.Println(ans)
}
