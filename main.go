package main

import (
	"os"
	"math"
	"fmt"
	"strconv"
	"time"
)

func Reverse(s string) string {
	var reverse string
	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	return reverse
}

func check_polindrom(Num1 int, Num2 int, start time.Time) {
	res := Num1 * Num2
	res1 := strconv.Itoa(res)
	res2 := Reverse(res1)
	if res1 == res2 {
		fmt.Print("Number 1 = ", Num1, "\n")
		fmt.Print("Number 2 = ", Num2, "\n")
		fmt.Println("Result = ", res, "\n")
		fmt.Println(time.Now().Sub(start))
		os.Exit(0)
	}
}

func find_next_prime(num int, start time.Time, res int) {
	for i := 0; i < num; i++ {
		_num := num - i
		if (len(strconv.Itoa(_num)) < res){
			return
		}
		if chekPrime(_num) {
			check_polindrom(int(num), int(_num), start)
		}
	}
}

func chekPrime(num int) bool {
	for i := 2; i < int(math.Sqrt(float64(num)))+1; i++ {
		if int(num)%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	start := time.Now()
	stdin := os.Args[1:]
	res, _ := strconv.Atoi(stdin[0])
	num := math.Pow10(res) - 1
	pr := math.Pow10(res - 1)
	for i := 0; i < int(pr); i++ {
		_num := int(num) - i
		if chekPrime(int(num) - i) {
			find_next_prime(_num, start, res)
		}
	}
}
