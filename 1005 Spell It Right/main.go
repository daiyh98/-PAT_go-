package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){
	var a string
	fmt.Scanln(&a)
	r := []rune(a)
	sum := 0
	for _, v := range r {
		n, _ := strconv.Atoi(string(v))
		sum += n
	}
	sSum := strconv.Itoa(sum)
	r2 := []rune(sSum)
	s := make([]string, len(r2))
	for i, v := range r2 {
		switch v {
		case '0':
			s[i] = " zero"
		case '1':
			s[i] = " one"
		case '2':
			s[i] = " two"
		case '3':
			s[i] = " three"
		case '4':
			s[i] = " four"
		case '5':
			s[i] = " five"
		case '6':
			s[i] = " six"
		case '7':
			s[i] = " seven"
		case '8':
			s[i] = " eight"
		case '9':
			s[i] = " nine"
		}
	}
	for i, v := range s {
		if i==0 {
			v = strings.TrimPrefix(v," ")
		}
		fmt.Print(v)
	}
}
