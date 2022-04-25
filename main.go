package main

import "log"

func main() {
	myString := "Green"

	log.Println("myString is set to", myString)
	changeUSingPointer(&myString)
	log.Panicln("after func call myString is set to", myString)
}

func changeUSingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
}
