package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck' which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Daimonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}

	return cards
}

// d deck is a receiver similar to class instantiation arguments
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// hS: handSize
func deal(d deck, hS int) (deck, deck) {
	return d[:hS], d[hS:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bS, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bS), ","))
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
