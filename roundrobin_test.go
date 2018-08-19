package rete_limit_algorithm

import (
	"testing"
	"fmt"
)

func TestRoundRobinAlg_Next(t *testing.T) {

	roundRobinAlg := &RoundRobinAlg{}
	roundRobinAlg.Add("A", 5)
	roundRobinAlg.Add("B", 2)
	roundRobinAlg.Add("C", 3)

	for i:=0; i<10; i++ {
		fmt.Printf("%s ", roundRobinAlg.Next())
	}

	var statistics  = make(map[interface{}]int)
	for i:=0; i<1000; i++ {
		statistics[roundRobinAlg.Next()]++
	}

	if statistics["A"] != 500 && statistics["B"] != 200 && statistics["C"] != 300 {
		t.Error("Round-Robin Algorithm is wrong", statistics)
	}

	roundRobinAlg.Remove("A")
	statistics  = make(map[interface{}]int)
	for i:=0; i<1000; i++ {
		statistics[roundRobinAlg.Next()]++
	}
	if statistics["B"] != 400 && statistics["C"] != 600 {
		t.Error("Round-Robin Algorithm is wrong", statistics)
	}
}