package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type edge struct {
	id string
	p  float64
}

func buildMarkovChange(src string) map[string][]edge {
	chars := strings.Split(src, "")
	chain := make(map[string]map[string]int)
	for i := 1; i < len(chars); i++ {
		edges := chain[chars[i - 1]]
		if edges == nil {
			edges = make(map[string]int)
			chain[chars[i - 1]] = edges
		}
		edges[chars[i]] = edges[chars[i]] + 1
	}
	result := make(map[string][]edge)
	for k, v := range chain {
		sum := 0
		for _, ct := range v {
			sum += ct
		}
		for lt, ct := range v {
			result[k] = append(result[k], edge{
				id: lt,
				p: float64(ct) / float64(sum),
			})
		}
	}
	return result
}

func main() {
	//could be incomplete
	fmt.Println("read from file")

	byteArray, err := ioutil.ReadFile("markov/src/main/sampletext");
	check(err)

	line := string(byteArray)
	markovMap := buildMarkovChange(line)

	for char, val := range markovMap {
		for _, edge := range val {
			fmt.Printf("Char: %s -> Char: %s -> p=%g\n", char, edge.id, edge.p)
		}
	}
}