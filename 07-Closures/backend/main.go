// Now implement these using closures:

// 1.
// shoppingCart()

// Return closures

// AddItem()

// RemoveItem()

// Total()

// Items()
// 2.
// rateLimiter(5)

// Output

// Allowed
// Allowed
// Allowed
// Allowed
// Allowed
// Blocked
// Blocked
// 3
// urlShortener()

// Output

// url1

// url2

// url3
package main

import (
	"fmt"
)

func main() {
	addToCart, removeFromCart, totalItems, reviewCart := shoppingCart()
	fmt.Println(addToCart("Apple", 5))
	fmt.Println(addToCart("Banana", 3))
	fmt.Println(addToCart("Orange", 2))
	fmt.Println(reviewCart())
	fmt.Println(totalItems())
	fmt.Println(removeFromCart("Banana", 1))
	fmt.Println(reviewCart())
	fmt.Println(totalItems())
	rl := rateLimiter(5)
	fmt.Println(rl())

	fmt.Println(rl())
	fmt.Println(rl())
	fmt.Println(rl())
	fmt.Println(rl())
	fmt.Println(rl())
	fmt.Println(rl())
	fmt.Println(rl())

	shortener, reslove := urlShortener()

	u1 := shortener("varunsuddala.me")
	fmt.Println(u1)
	fmt.Println(reslove(u1))

}
