package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 6, 8, 9}

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])
	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	s = append(s, 110)
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	// s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	meioDoArrayEsquerda := s[:(len(s))/2]
	fmt.Println("Meio do array (arredondado para baixo):", meioDoArrayEsquerda)

	meioDoArrayDireita := s[(len(s))/2:]
	fmt.Println("Meio do array (arredondado para cima):", meioDoArrayDireita)

	fmt.Println(s)
	d := append(s, 10)
	fmt.Println(d)
	println(len(s))
	println(cap(s))

}

// s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
// write a binary search implementation
// func binarySearch(arr []int, target int) int {
// 	left, right := 0, len(arr)-1

// 	for left <= right {
// 		mid := left + (right-left)/2

// 		if arr[mid] == target {
// 			return mid
// 		} else if arr[mid] < target {
// 			left = mid + 1
// 		} else {
// 			right = mid - 1
// 		}
// 	}

// 	return -1
// }
