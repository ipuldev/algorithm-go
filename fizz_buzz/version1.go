package main

import "fmt"

func version1() []string {
	//slice variable to storing log activity
	var fizzBuzz []string
	for i := 1; i <= 100; i++ {

		if i%3 == 0 && i%5 == 0 {
			fizzBuzz = append(fizzBuzz, "FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fizzBuzz = append(fizzBuzz, "Fizz")
			continue
		}
		if i%5 == 0 {
			fizzBuzz = append(fizzBuzz, "Buzz")
			continue
		}
		fizzBuzz = append(fizzBuzz, fmt.Sprint(i))
	}
	return fizzBuzz
}
