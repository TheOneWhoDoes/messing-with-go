package main

import (
	"fmt"
)

func increment(x *int) {
	*x++
}

type workers interface {
	allSingle() bool
}

type alii struct {
	son    string
	single bool
}

func (x alii) allSingle() bool {
	return x.single
}

type hassan struct {
	daughter string
	single   bool
}

func (x hassan) allSingle() bool {
	return x.single
}

func allSingle(w ...workers) {
	var y bool
	for _, val := range w {
		y = val.allSingle()
		fmt.Println(y)
	}
}

func (x *alii) function() {
	x.single = false
}

func arrfunc(x []int) {
	x[0] = 0
}

func main() {
	var ali string = "sa"
	g := []rune(ali)
	g[1] = 'L'
	l := string(g)
	c := "ma"
	fmt.Println("I'm coming", ali, l, len(ali), c[0:1], ali[0:2])

	if j := ali + ali[0:1]; j < ali[0:1] {
		fmt.Println("baby")
	}

	i := 2
	switch i {
	case 1, 2, 4:
		fmt.Println("baby")
	}

	switch {
	case i == 2:
		fmt.Println("baby")
	default:
		fmt.Println("baby2")
	}

	for j := 0; j < len(ali); j++ {
		fmt.Println(string(ali[j]))
	}

	arr := []string{"ali", "mammad", "shahin"}
	for _, val := range arr {
		fmt.Println(val)
	}

	fmt.Println(i)
	increment(&i)
	fmt.Println(i)

	farshid := alii{single: true}
	mammad := hassan{single: false}
	// ptfarshid := &farshid
	fmt.Println(farshid.son)
	// fmt.Println(ptfarshid.single)
	// farshid.function()
	// fmt.Println(farshid.single)
	allSingle(farshid, mammad)

	var arrr [2]string
	arrr[0] = "ali"
	arrr[1] = "mammad"
	fmt.Println(arrr)

	var x []int
	x = append(x, 3)

	arrfunc(x)
	fmt.Println(x)

	y := []int{0, 2, 4}
	fmt.Println(cap(y))

	z := make([]int, 3)
	fmt.Println(z)

	o := y[0:2]

	n := make([]int, 0, 3)
	fmt.Println(n, n[0:2], o[0:3])

}
