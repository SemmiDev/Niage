package main

import "fmt"

func main(){
	fmt.Println(no7() == no8())
}

func no7() bool {
	x := []int{2,3,4}
	y := []int{3,4,5}

	// z = 2x + y
	var result, expected []int

	for i, v := range x {
		result = append(result, 2 * v + y[i])
	}

	temp := []int{7,10,13}
	expected = append(expected, temp...)

	for i := range result {
		if (result[i] == expected[i]) != true {
			return false
		}
	}
	return true
}

func no8() bool {
	x := []int{2,3,4}
	y := []int{3,4,5}

	// dot product = 6 + 12 + 20 = 38
	var result int
	for i := range x {
		result += x[i] * y[i]
	}

	return result == 38
}
