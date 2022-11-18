package main

import "fmt"

func main() {
	list := []int{4, 3, 5, 2, 12}
	fmt.Println(list)
	bubble(&list, len(list))
	fmt.Println(list)
}

func bubble(arr *[]int, n int) {
	if arr != nil && (n == 0 || n == 1) {
		return
	}
	var temp int
	for i := 0; i < n-1; i++ {
		if (*arr)[i] > (*arr)[i+1] {
			temp = (*arr)[i]
			(*arr)[i] = (*arr)[i+1]
			(*arr)[i+1] = temp
		}
	}
	bubble(arr, n-1)
}
