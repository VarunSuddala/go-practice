package utils

func IsValidEmail(email string) bool {
	// Basic email validation logic
	if len(email) < 5 || len(email) > 50 {
		return false
	}
	return true
}
