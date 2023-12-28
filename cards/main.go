package main

func main() {
	cards := newDeck()
	left, right := deal(cards, 10)

	left.saveToFile("left.txt")
	right.saveToFile("right.txt")

	cards = newDeck()
	cards.shuffle()
	cards.print()
}
