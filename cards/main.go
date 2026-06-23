package main

func main() {
	// var card string = "Ace of Spades"

	cards := newDeck()

	//hand, remainingCards := deal(cards, 5)
	//
	//hand.print()
	//remainingCards.print()

	//fmt.Println(cards.toString())

	//cards.saveToFile("my__cards")

	//cards := newDeckFromFile("my")
	//cards.print()

	cards.shuffle()
	cards.print()
}
