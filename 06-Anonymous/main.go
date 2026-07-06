// Print Hello
// Sum
// Square
// Even/Odd
// Maximum
// Calculate
// Greeting
// Filter
// Apply Operation
// Custom Sort

package main

import "fmt"

func main() {
	func() string {
		return "Hello"
	}()

	sum := func(a int, b int) int {
		return a + b
	}

	fmt.Println(sum(5, 6))

	square := func(n int) int {
		return n * n
	}
	fmt.Println(square(8))

	EvenOdd := func(n int) string {
		if n%2 == 0 {
			return "Even"
		}
		return "Odd"
	}

	fmt.Println(EvenOdd(9))

	maximum := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	fmt.Println(maximum(9, 10))

	calculator := func(a int, b int, op func(int, int) int) int {

		return op(a, b)

	}

	fmt.Println(calculator(5, 6, sum))

	greet := func(name string) string {
		return "hello " + name + " !"
	}

	fmt.Println(greet("varun"))

	filter := func(nums []int, condition func(int) bool) []int {

		var res []int

		for _, num := range nums {
			if condition(num) {
				res = append(res, num)
			}
		}
		return res

	}
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	even := filter(nums, func(n int) bool {
		return n%2 == 0
	})

	fmt.Println(even)

	odd := filter(nums, func(i int) bool {
		return i%2 != 0
	})
	fmt.Println(odd)

	apply := func(nums []int, op func(int) int) []int {

		var res []int

		for _, i := range nums {
			res = append(res, op(i))
		}
		return res
	}
	squarenums := apply(nums, square)
	fmt.Println(squarenums)

}
