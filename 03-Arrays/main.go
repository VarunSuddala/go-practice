// Sum of elements
// Maximum element
// Minimum element
// Average
// Count even and odd
// Linear search
// Count occurrences
// Reverse array
// Check if sorted
// Second largest
// Rotate left by one
// Rotate right by one
// Merge two arrays
// Remove duplicates
// Find the missing number

package main

import (
	"fmt"
)

func main() {
	fmt.Println("")
	arr := [6]int{1, 2, 3, 4, 5}
	sum := sumOfElements(arr[:])

	arr1 := []int{5, 2, 8, 1, 3}
	maxi, mini := maxMiniElement(arr1[:])
	fmt.Println(sum)
	fmt.Println(maxi, mini)

	avg := Average(arr1[:])
	fmt.Println(avg)
	arr2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(CountEvenOdd(arr2[:]))
	fmt.Println(LinearSearch(arr2[:], 5))
	fmt.Println(CountOccurrences(arr2[:], 5))
	reversed := ReverseArray(arr2[:])
	fmt.Println(reversed)
	fmt.Println(checkIfSorted(arr2[:]))

	fmt.Println(secondLargest(arr2[:]))

	merged := mergeArrays([]int{1, 3, 5}, []int{2, 4, 6})
	fmt.Println(merged)

	rotatedLeft := rotate([]int{1, 2, 3, 4, 5}, "left")
	fmt.Println(rotatedLeft)
	rotatedRight := rotate([]int{1, 2, 3, 4, 5}, "right")
	fmt.Println(rotatedRight)

	removedDuplicates := removeDuplicates([]int{1, 2, 2, 3, 4, 4, 5})
	fmt.Println(removedDuplicates)

	arr3 := []int{1, 2, 3, 4, 6}
	max2, _ := maxMiniElement(arr3[:])
	missingNumber := findMissingNumber(arr3[:], max2)
	fmt.Println(missingNumber)

}

func sumOfElements(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}

	return sum

}

func maxMiniElement(arr []int) (int, int) {

	if len(arr) <= 0 {
		return 0, 0
	}

	maxi := arr[0]
	mini := arr[0]

	for _, val := range arr {
		if val > maxi {
			maxi = val
		}
		if val < mini {
			mini = val
		}
	}

	return maxi, mini
}

func Average(arr []int) float64 {
	len := len(arr)
	if len <= 0 {
		return 0
	}
	sum := sumOfElements(arr)
	return float64(sum) / float64(len)
}

func CountEvenOdd(arr []int) (int, int) {
	evenCount := 0
	oddCount := 0

	for _, val := range arr {
		if val%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}

	}
	return evenCount, oddCount
}

func LinearSearch(arr []int, target int) int {
	for i, val := range arr {
		if val == target {
			return i
		}
	}
	return -1

}

func CountOccurrences(arr []int, target int) int {
	count := 0
	for _, val := range arr {
		if val == target {
			count++
		}

	}
	return count
}

func ReverseArray(arr []int) []int {
	reversed := make([]int, len(arr))
	for i, val := range arr {
		reversed[len(arr)-1-i] = val
	}
	return reversed
}

func checkIfSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func secondLargest(arr []int) int {
	if len(arr) < 2 {
		return -1 // Not enough elements for second largest
	}

	maxi := arr[0]
	secondMax := -1

	for _, val := range arr {
		if val > maxi {
			secondMax = maxi
			maxi = val
		}
		if val > secondMax && val < maxi {
			secondMax = val
		}

	}

	return secondMax
}

func rotate(arr []int, direction string) []int {
	if len(arr) == 0 {
		return arr
	}

	if direction == "left" {
		first := arr[0]
		for i := 0; i < len(arr)-1; i++ {
			arr[i] = arr[i+1]
		}
		arr[len(arr)-1] = first
	}

	if direction == "right" {
		last := arr[len(arr)-1]
		for i := len(arr) - 1; i > 0; i-- {
			arr[i] = arr[i-1]
		}
		arr[0] = last
	}

	return arr
}

func mergeArrays(arr1 []int, arr2 []int) []int {
	merged := make([]int, 0, len(arr1)+len(arr2))

	i := 0
	j := 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			merged[i+j] = arr1[i]
		} else {
			merged[i+j] = arr2[j]
		}

	}
	merged = append(merged, arr1[i:]...)
	merged = append(merged, arr2[j:]...)
	return merged
}

func removeDuplicates(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	prev := arr[0]
	result := []int{prev}

	for i := 1; i < len(arr); i++ {
		if arr[i] != prev {
			result = append(result, arr[i])
			prev = arr[i]
		}

	}
	return result
}

func findMissingNumber(arr []int, n int) int {
	expectedSum := n * (n + 1) / 2
	actualSum := sumOfElements(arr)
	return expectedSum - actualSum
}
