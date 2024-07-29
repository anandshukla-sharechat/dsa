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

func abs(a float64) float64 {
	if a < 0.0 {
		a = 0 - a
	}
	return a
}

var diff float64 = 0.00001

func solve(adjMap map[string][]Edge, sourceNode, targetNode string) bool {
	if sourceNode == targetNode {
		ls := adjMap[sourceNode]
		for _, val := range ls {
			if val.A == val.B && val.Wt != 1.0 {
				return true
			}
		}
		return false
	}
	q := list.New()
	mp := make(map[string]bool)
	ans := make([]float64, 0)
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
					if val.B == targetNode {
						ans = append(ans, newRatio)
						continue
					}
					q.PushBack(Ratio{A: val.B, R: newRatio})
				}
			}
		}
	}

	if len(ans) >= 2 {
		for i := 0; i < len(ans)-1; i++ {
			if abs(ans[i]-ans[i+1]) >= diff {
				return true
			}
		}
	}

	return false
}

func makegraph(equations [][]string, values []float64) map[string][]Edge {
	mp := make(map[string][]Edge)

	for i, val := range equations {
		mp[val[0]] = append(mp[val[0]], Edge{A: val[0], B: val[1], Wt: values[i]})
		mp[val[1]] = append(mp[val[1]], Edge{A: val[1], B: val[0], Wt: 1 / values[i]})
	}
	return mp
}

func checkContradictions(equations [][]string, values []float64) bool {
	mp := makegraph(equations, values)
	keys := make([]string, 0)
	for k, _ := range mp {
		keys = append(keys, k)
	}
	for i := 0; i < len(keys); i++ {
		for j := i; j < len(keys); j++ {
			ans := solve(mp, keys[i], keys[j])
			if ans {
				return ans
			}
		}
	}

	return false
}

// https://leetcode.com/problems/check-for-contradictions-in-equations/description/
func main() {
	equations, values := [][]string{{"a", "b"}, {"b", "c"}, {"a", "c"}}, []float64{3, 0.5, 1.5}
	ans := checkContradictions(equations, values)
	fmt.Println(ans)
}
