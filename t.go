package main

import "fmt"

func main() {
	v := []int{1, 2, 3}
	s := v[:2]
	fmt.Println(s)
	s = append(s, 10)
	fmt.Println(v)

	v = []int{1, 2, 3}
	s = v[:2:2]
	fmt.Println(s)
	s = append(s, 10)
	fmt.Println(v)

}
