package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a

	a = append(a, 4, 5, 6)
	multiple(a)

	fmt.Printf("%v\n", a) //
	fmt.Printf("%v\n", b) //
}

func multiple(a []int) {
	a = append(a, 10)
	for i, _ := range a {
		a[i] = a[i] * 2
	}
}
