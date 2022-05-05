package main

import "log"

func main() {
	type user struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []user
	users = append(users, user{"John", "Smith", "john@smith.com", 30})
	users = append(users, user{"Mary", "Jones", "mary@jones.com", 20})
	users = append(users, user{"Sally", "Brown", "sally@brown.com", 45})
	users = append(users, user{"Alex", "Anderson", "alex@anderson.com", 17})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}
