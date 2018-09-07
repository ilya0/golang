package main //
import "fmt"

type deck []string // creating a deck class with a string type

func newDeck() deck { // newDeck method that returns a deck string type
	cards := deck{}                                        // create an empty so save the concatinations to
	cardSuits := []string{"Spades", "Diamonds", "Hearts"}  //create the suits iterations
	cardValues := []string{"aces", "Two", "Three", "Four"} //create the values iterations

	for _, suit := range cardSuits { // for loop to go through car  suits

		for _, value := range cardValues { // second loop to go through the values
			cards = append(cards, value+" of  "+suit) // appends to cards
		}
	}
	return cards // return the cards
}

func (d deck) print() { //print loop

	for i, card := range d { // iterates through the card variable
		fmt.Println(i, card) // print the cards in the loop
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]

}
