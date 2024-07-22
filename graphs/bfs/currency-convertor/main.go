/**

UBER 1st DSA Round - 22/07/2024

We are a currency exchange that maintains the current exchange rates between currencies. A user can come to us with some amount in one currency and request the equivalent amount in a different currency. Given a list of exchange rates between currencies, write a function that calculates currency rate between any 2 currencies.

Example 1 -
(GBP, EUR, 10)     - read as "1 GBP equals 10 EUR"
(EUR, USD, 1.1)    - "1 EUR equals 1.1 USD"
(10 GBP, USD) => ? - "10 GBP equals how many USD?"

Answer: 110


adj list : map[string][]Edge

    data := [][]string{
        {"EUR", "USD", "1.1"},
        {"USD", "INR", "70.0"},
        {"INR", "JPY", "1.5"},
        {"EUR", "GBP", "0.9"},
        {"GBP", "MP", "27.3"},
        {"MP", "JPY", "4.7"},
    }

    (10 EUR, JPY)

*/

package main

import (
	"container/list"
	"errors"
	"fmt"
	"strconv"
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

func solve(adjMap map[string][]Edge, sourceNode, targetNode string) (float64, error) {

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
						return newRatio, nil
					}
				}
			}
		}
	}

	return -1, errors.New("error : info not found for currency")

}

func main() {
	mp := make(map[string][]Edge)
	data := [][]string{
		{"EUR", "USD", "1.1"},
		{"USD", "INR", "70.0"},
		{"INR", "JPY", "1.5"},
		{"EUR", "GBP", "0.9"},
		{"GBP", "MP", "27.3"},
		{"MP", "JPY", "4.7"},
	}
	for _, val := range data {
		w, _ := strconv.ParseFloat(val[2], 64)
		mp[val[0]] = append(mp[val[0]], Edge{A: val[0], B: val[1], Wt: w})
		mp[val[1]] = append(mp[val[1]], Edge{A: val[1], B: val[0], Wt: 1 / w})
	}

	ratio, err := solve(mp, "EUR", "JPY")
	b := 10.0

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(b * ratio)
	}
}
