package main

import "fmt"

func main() {
	cards := newDeck()
	deck := newDeck()
	hand := newDeck()
	hand, deck = cards.deal(3)

	fmt.Println("Hand")
	hand.print()

	fmt.Println("Deck")
	deck.print()

	hand, deck = dealTheDumbWay(cards, 4)

	fmt.Println("Hand")
	hand.print()

	fmt.Println("Deck")
	deck.print()

	fmt.Println(cards.toString())

	deck.saveToFile("deck.txt")

	d2 := newDeckFromFile("deck.txt")

	fmt.Println("Deck 2:")
	d2.print()
	d2.shuffle()
	fmt.Println("Deck 2  shuf:")
	d2.print()

}
