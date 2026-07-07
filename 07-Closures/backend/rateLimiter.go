package main

func rateLimiter(limit int) func() string {
	count := 0

	return func() string {
		if count < limit {
			count++
			return "Allowed"
		}
		return "Deined"
	}
}
