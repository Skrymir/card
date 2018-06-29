package main

import "fmt"

func main() {
	cards := newDeck()

	hand, deck := cards.deal(3)

	fmt.Println("Hand")
	hand.print()

	fmt.Println("Deck")
	deck.print()

	hand, deck = dealTheDumbWay(cards, 4)

	fmt.Println("Hand")
	hand.print()

	fmt.Println("Deck")
	deck.print()

	fmt.Print(cards.toString())

	deck.saveToFile("deck.txt")

}
