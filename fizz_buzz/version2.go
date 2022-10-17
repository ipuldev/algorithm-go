package main

import (
	"fmt"
)

func version2() []string {
	//slice variable to storing log activity
	var fizzBuzz []string
	var count3 int
	var count5 int
	for i := 1; i <= 100; i++ {
		count3++
		count5++
		if count3 == 3 && count5 == 5 {
			fizzBuzz = append(fizzBuzz, "FizzBuzz")
			count3 = 0
			count5 = 0
			continue
		}

		if count3 == 3 {
			fizzBuzz = append(fizzBuzz, "Fizz")
			count3 = 0
			continue
		}
		if count5 == 5 {
			fizzBuzz = append(fizzBuzz, "Buzz")
			count5 = 0
			continue
		}
		fizzBuzz = append(fizzBuzz, fmt.Sprint(i))
	}
	return fizzBuzz
}
