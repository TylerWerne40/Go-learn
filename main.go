package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func main() {
	{
		var str = "Hello Go"
		fmt.Println(str)

		i := 1
		for i <= 10 {
			fmt.Println(i)
			i = i + 1
		}
		if i%2 == 0 {
			// even
		} else {
			// odd
		}
		for i := 1; i <= 10; i++ {
			// iterate 10 times
		}

		switch i {
		case 10:
			fmt.Println("i is 10")
		default:
			fmt.Println("i is not 10")
		}
	}
	var y [5]int
	y[4] = 100
	fmt.Println(y)
	{
		z := make([]float64, 10)
		z2 := append(z, 1)
		fmt.Println(z2)
	}
	/*	{
		var a map[string]int
		a["key"] = 10
		fmt.Println(a)
		b := make(map[int]int)
		b[1] = 10
		delete(b, 1)
	}*/
	{
		xfl := []float64{98, 93, 77, 82, 83}
		fmt.Println(average(xfl))
		fmt.Println(add(1, 2, 3))

		sub := func(a int, b int) int {
			return a - b
		}
		fmt.Println(sub(1, 2))
	}
	defer func() {
		defer func() {
			str := recover()
			fmt.Println(str)
		}()
		panic("PANIC")

	}()
	{
		c := new(Circle)

		c.r = 5
		fmt.Println(c.area())

	}
}

func (c *Circle) area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
