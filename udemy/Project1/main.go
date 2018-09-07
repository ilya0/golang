package main

func main() { // run the main program
	cards := newDeck() // cards is a the program itter

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}
