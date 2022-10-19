package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a

	//multiple(a)

	fmt.Printf("%v\n", a) //
	fmt.Printf("%v\n", b) //
}

func multiple(a []int) {
	for i, _ := range a {
		a[i] = a[i] * 2
	}
}
