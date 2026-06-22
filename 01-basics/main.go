package main

import "fmt"

func main() {
	a := 10
	b := 5

	fmt.Printf(" a  : = %v , b  : = %v \n", a, b)

	a = a + b
	b = a - b
	a = a - b
	fmt.Println("after swaping")
	fmt.Printf(" a  : = %v , b  : = %v \n ", a, b)
	a, b = b, a
	fmt.Println("after swaping")
	fmt.Printf(" a  : = %v , b  : = %v ", a, b)
	x := 2
	y := 3
	z := 1
	fib := 5
	num := 123456
	if evenOdd(x) {
		fmt.Printf("%v is even number \n", x)
	} else {
		fmt.Printf("%v is odd number \n ", x)
	}

	fmt.Printf("largest number in %v ,%v ,%v  is : = %v \n ", x, y, z, largestNumber(x, y, z))

	fmt.Printf("factorial of %v is %v \n", y, factorial(y))

	fmt.Printf("Fibonacc of %v is : %v \n", fib, fibonacc(fib))
	fmt.Printf("reverse a number of %v is : %v \n", num, reverseNumber(num))

	fmt.Printf("palindrom of %v is : %v \n", num, palindrom(num))
	fmt.Printf("number of digits in %v is : %v \n", num, countDigits(num))
	fmt.Printf("sum of digits in %v is : %v \n", num, sumOfDigits(num))

}

func evenOdd(x int) bool {
	return x%2 == 0
}

func largestNumber(x, y, z int) int {
	largest := x
	if y > largest {
		largest = y
	}
	if z > largest {
		largest = z
	}
	return largest
}

func factorial(x int) int {
	fact := 1

	for i := 1; i <= x; i++ {
		fact *= i
	}

	return fact

}
func fibonacc(x int) int {
	// a,b:=0,1
	// for i:=0 ; i<x ;i++{
	// 	a,b=b,a+b
	// }
	// return b

	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}

	return fibonacc(x-1) + fibonacc(x-2)

}

func reverseNumber(num int) int {
	reverse := 0
	rem := 0
	for num > 0 {
		rem = num % 10
		reverse = reverse*10 + rem
		num = num / 10
	}

	return reverse

}
func palindrom(num int) bool {
	return num == reverseNumber(num)
}

func countDigits(num int) int {
	if num == 0 {
		return 1
	}
	count := 0
	for num > 0 {
		num = num / 10
		count++
	}

	return count
}

func sumOfDigits(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num = num / 10
	}

	return sum
}
