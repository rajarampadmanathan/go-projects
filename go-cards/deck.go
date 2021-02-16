package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	numbers := []string{"ace", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "jack", "king", "queen"}
	varieties := []string{"heart", "diamond", "spade", "clubs"}
	for _, variety := range varieties {
		for _, number := range numbers {
			cards = append(cards, number+" of "+variety)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func readDeckFromFile(filename string) deck {
	deckBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	deckString := string(deckBytes)
	return deck(strings.Split(deckString, ","))
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) deal(hand int) (deck, deck) {
	return d[:hand], d[hand:]
}

func (d deck) shuffle() {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	for i := range d {
		randomNum := r.Intn(len(d) - 1)
		d[i], d[randomNum] = d[randomNum], d[i]
	}
}
