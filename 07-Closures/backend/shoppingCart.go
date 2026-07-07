package main

func shoppingCart() (func(string, int) map[string]int, func(string, int) map[string]int, func() int, func() map[string]int) {
	cart := make(map[string]int)

	addToCart := func(item string, quantity int) map[string]int {
		cart[item] += quantity
		return cart
	}
	removeFromCart := func(item string, quantity int) map[string]int {
		if cart[item] >= quantity {
			cart[item] -= quantity
		}
		return cart
	}

	totalItems := func() int {
		total := 0
		for _, quantity := range cart {
			total += quantity
		}
		return total
	}
	reviewCart := func() map[string]int {
		return cart
	}
	return addToCart, removeFromCart, totalItems, reviewCart
}
