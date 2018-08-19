package rete_limit_algorithm

import (
	"testing"
	"fmt"
)

func TestSmoothRoundRobinAlg_Next(t *testing.T) {

	smoothRoundRobinAlg := &SmoothRoundRobinAlg{}
	smoothRoundRobinAlg.Add("A", 5)
	smoothRoundRobinAlg.Add("B", 2)
	smoothRoundRobinAlg.Add("C", 3)

	for i:=0; i<7; i++ {
		fmt.Printf("%s ", smoothRoundRobinAlg.Next())
	}

	var statistics  = make(map[interface{}]int)
	for i:=0; i<1000; i++ {
		statistics[smoothRoundRobinAlg.Next()]++
	}

	if statistics["A"] != 500 && statistics["B"] != 200 && statistics["C"] != 300 {
		t.Error("Round-Robin Algorithm is wrong", statistics)
	}

	smoothRoundRobinAlg.Remove("A")
	statistics  = make(map[interface{}]int)
	for i:=0; i<1000; i++ {
		statistics[smoothRoundRobinAlg.Next()]++
	}
	if statistics["B"] != 400 && statistics["C"] != 600 {
		t.Error("Round-Robin Algorithm is wrong", statistics)
	}
}
