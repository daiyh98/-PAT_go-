package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b int
	fmt.Scanln(&a, &b)
	c := a + b
	s := strconv.Itoa(c)
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		fmt.Print(string(r[i]))
		if string(r[i]) == "-" {
			continue
		}
		if (i+1)%3 == len(r)%3 && i != len(r)-1 {
			fmt.Print(",")
		}
	}
}

//if c >= -999 && c <= 999 {
//	fmt.Println(strconv.Itoa(c))
//}else {
//	s := strconv.Itoa(c)
//	r := []rune(s)
//	s = strconv.Itoa(c/1000) + "," + string(r[len(s)-3:len(s)])
//	fmt.Println(s)
//}
