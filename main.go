package main

import "fmt"

func main() {
	// cards:= newDeck()
	// hand,remainingCards :=deal(cards,22)
	// hand.print()
	// remainingCards.print()
	// cards.saveToFile("my_cards")
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	fmt.Println(cards)

}

func newCard() string {
	return "Five of Diamonds"
}
