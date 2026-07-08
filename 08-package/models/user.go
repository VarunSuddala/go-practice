package models

type User struct {
	Name     string
	Email    string
	Password string
}

func (u User) Display() string {
	return "Name: " + u.Name + "\nEmail: " + u.Email
}
