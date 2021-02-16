package main

import "fmt"

func main() {
	oddeven()
	cards := newDeck()
	hand, deal := cards.deal(4)
	hand.print()
	fmt.Println("After shuffle")
	hand.shuffle()
	hand.saveToFile("my_hand")
	deal.saveToFile("my_deal")
	readDeckFromFile("my_hand").print()
}
