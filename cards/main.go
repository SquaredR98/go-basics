package main

import "fmt"

func main() {
	cards := []string{newCard()}
	cards = append(cards, "Ace of Hearts")

	// Looping slices
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

// Add return type of the function otherwise void error will be thrown
func newCard() string {
	return "Five of Diamonds"
}
