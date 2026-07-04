// Create a map
// Add key-value pairs
// Update a value
// Delete a key
// Check if a key exists
// Print all entries
// Frequency count (integers)
// Frequency count (characters)
// Frequency count (words)
// First non-repeating element
// Most frequent element
// Remove duplicates
// Inventory system
// Two Sum

package main

import "fmt"

func main() {
	// map1 := make(map[string]int)
	// fmt.Println("Initial map:", map1)

	// map1["varun"] = 35
	// map1["john"] = 25
	// map1["alice"] = 30
	// fmt.Println("After adding key-value pairs:", map1)

	// map1["varun"] = 36
	// fmt.Println("After updating a value:", map1)
	// delete(map1, "john")
	// fmt.Println("After deleting a key:", map1)

	// val, ok := map1["varun"]
	// if ok {
	// 	fmt.Println("Key exists with value:", val)
	// } else {
	// 	fmt.Println("Key does not exist")
	// }

	// fmt.Println("All entries in the map:")
	// for key, val := range map1 {
	// 	fmt.Println(key, ":", val)
	// }

	// nums := []int{1, 2, 2, 3, 4, 4, 5}
	// fmt.Println(frequencyCount(nums))
	// char := []string{"a", "b", "a", "c", "b", "d"}
	// fmt.Println(frequencyChar(char))
	// words := []string{"apple", "banana", "apple", "orange", "banana", "grape"}
	// fmt.Println(frequencyWords(words))
	// fmt.Println(firstNonRepeatingElement(nums))
	// fmt.Println(Mostfrequentelement(nums))
	// fmt.Println(RemoveDuplicates(nums))
	// nums2 := []int{2, 7, 11, 15}
	// target := 9
	// fmt.Println(twoSum(nums2, target))
	inventory := inventory()
	fmt.Println("Initial Inventory:", inventory)
	addProduct(inventory, "P004", Product{Name: "Smartwatch", Price: 199.99, Quantity: 15})
	fmt.Println("After adding a product:", inventory)
	removeProduct(inventory, "P002")
	fmt.Println("After removing a product:", inventory)
	updateProductQuantity(inventory, "P001", 20)
	fmt.Println("After updating product quantity:", inventory)
	product, exists := retrieveProduct(inventory, "P003")
	if exists {
		fmt.Println("Retrieved product:", product)
	}
	incrementProductQuantity(inventory, "P001", 5)
	fmt.Println("After incrementing product quantity:", inventory)
	dropProductQuantity(inventory, "P001", 10)
	fmt.Println("After dropping product quantity:", inventory)

	printInventory(inventory)

}

func frequencyCount(nums []int) map[int]int {
	freqMap := make(map[int]int)

	for _, num := range nums {
		freqMap[num]++
	}
	return freqMap
}
func frequencyChar(chars []string) map[string]int {
	freqMap := make(map[string]int)

	for _, char := range chars {
		freqMap[char]++
	}
	return freqMap
}
func frequencyWords(words []string) map[string]int {
	freqMap := make(map[string]int)

	for _, word := range words {
		freqMap[word]++
	}
	return freqMap
}

func firstNonRepeatingElement(nums []int) int {

	if len(nums) == 0 {
		return -1
	}
	freqmap := frequencyCount(nums)
	for _, num := range nums {
		if freqmap[num] == 1 {
			return num
		}
	}
	return -1
}
func Mostfrequentelement(nums []int) []int {
	res := []int{}
	freqmap := frequencyCount(nums)
	maxFreq := 0
	for _, freq := range freqmap {
		if freq > maxFreq {
			maxFreq = freq
		}

	}
	for num, freq := range freqmap {
		if freq == maxFreq {
			res = append(res, num)
		}
	}
	return res
}
func RemoveDuplicates(nums []int) []int {
	// freqmap := frequencyCount(nums)
	// res := []int{}

	// for num := range freqmap {
	// 	res = append(res, num)
	// }
	// return res
	freqmap := make(map[int]bool)
	res := []int{}

	for _, num := range nums {
		if !freqmap[num] {
			freqmap[num] = true
			res = append(res, num)
		}

	}
	return res
}

func twoSum(nums []int, target int) []int {
	freqmap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if index, ok := freqmap[complement]; ok {
			return []int{index, i}
		}
		freqmap[num] = i
	}

	return []int{}

}
