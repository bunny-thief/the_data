package main

import (
	"fmt"

	"github.com/bunny-thief/the_data/palindrome"
)

func main() {
	sator := "sator arepo tenet opera rotas"
	fmt.Printf("%q is a palindrome: %v\n", sator, palindrome.IsPalindrome(sator))

	race := "racecar"
	fmt.Printf("%q is a palindrome: %v\n", race, palindrome.IsPalindrome(race))

	rat := "rattata"
	fmt.Printf("%q is a palindrome: %v\n", rat, palindrome.IsPalindrome(rat))
}
