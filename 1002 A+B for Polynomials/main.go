package main

import (
	"fmt"
)


func main () {
	var na, nb int //项数
	c := make([]float64, 1001)
	var t int
	var num float64
	fmt.Scan(&na)
	for i := 0; i < na; i++ {
		fmt.Scan(&t, &num)
		c[t] += num
	}
	fmt.Scan(&nb)
	for i := 0; i < nb; i++ {
		fmt.Scan(&t, &num)
		c[t] += num
	}
	cnt := 0
	for i, _ := range c {
		if c[i] != 0 {
			cnt++
		}
	}
	fmt.Print(cnt)
	for i := 1000; i >= 0; i-- {
		if c[i] != 0{
			fmt.Printf(" %d %.1f",i,c[i])
		}
	}
}