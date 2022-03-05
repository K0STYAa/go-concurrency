// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
package main

import "fmt"

func createSet(x []string) []string {
	set := make(map[string]bool)

	for _, i := range x {
		set[i] = true
	}

	out := make([]string, 0, len(set))
	for i := range set {
		out = append(out, i)
	}

	return out
}

func main() {
	x := []string{"cat", "cat", "dog", "cat", "tree"}

	y := createSet(x)
	fmt.Println(y)
}
