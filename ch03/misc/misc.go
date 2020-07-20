package main

import (
	"fmt"
	"math"
)

var x uint8 = 1<<1 | 1<<5
var y uint8 = 1<<1 | 1<<2

func main() {
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)

	fmt.Printf("%08b\n", x^y)
	fmt.Printf("%08b\n", x&^y)

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}

	fmt.Printf("%08b\n", x<<1)
	fmt.Printf("%08b\n", x>>1)

	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i])
	}

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d	ex  =  %8.3f\n", x, math.Exp(float64(x)))
	}

	var a complex128 = complex(3, 4)
	var b complex128 = complex(1, 2)
	fmt.Println(a * b)
	fmt.Println(real(a * b))
	fmt.Println(imag(a * b))

	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7]) // 104 and 119
	fmt.Println(s[:5])

}
