package main

import (
	"os"
	"math"
	"fmt"
	"strconv"
)

func Reverse(s string) string {
	var reverse string
	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	return reverse
}

func check_polindrom(Num1 int, Num2 int) {
	res := Num1 * Num2
	res1 := strconv.Itoa(res)
	res2 := Reverse(res1)
	if res1 == res2 {
		fmt.Println(res)
		os.Exit(0)
	}
}

func find_next_prime(num float64, pr int) {
	for i := 0; i < pr; i++ {
		_num := num - float64(i)
		if chekPrime(_num, int(pr)) {
			check_polindrom(int(num), int(_num))
		}
	}
}

func chekPrime(num float64, pr int) bool{
	for i := 2; i < int(math.Sqrt(num))+1; i++ {
		if int(num)%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	stdin := os.Args[1:]
	res, _ := strconv.Atoi(stdin[0])
	num := math.Pow10(res) - 1
	pr := math.Pow10(res - 1)
	for i := 0; i < int(pr); i++ {
		_num := num-float64(i)
		if chekPrime(_num, int(pr)){
			find_next_prime(_num, int(pr))
		}
	}
}
