package main

import (
	"fmt"
)

func main() {
	var s1 string
	fmt.Println(s1)

	var i, j, k int
	fmt.Println(i, j, k)
	var b, f, s = true, 2.3, "four"
	fmt.Println(b, f, s)

	l := 100
	var boiling float64 = 100
	var names []string
	var err error
	fmt.Println(l, boiling, names, err)

	m, n := 0, 1
	fmt.Println(m, n)

	m, n = n, m
	fmt.Println(m, n)

	x := 1
	p := &x
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)

	var c, d int
	fmt.Println(&c == &c, &c == &d, &c == nil)

	// var q1 = f()
	// fmt.Println(f() == f())

	pNew := new(int)
	fmt.Println(*pNew)
	*pNew = 2
	fmt.Println(*pNew)

	medals := []string{"gold", "silver", "bronze"}
	fmt.Println(medals)
}

func f() *int {
	v := 1
	return &v
}
