package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []card

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardNames := []string{"Ace", "Two", "Three", "Four"}
	cardValues := []int{1, 2, 3, 4}

	for _, suit := range cardSuits {
		for i, name := range cardNames {
			newCard := card{value: cardValues[i], name: name + " of " + suit}
			cards = append(cards, newCard)
		}
	}

	return cards
}

func (d deck) print() {
	fmt.Println("==================================")
	for _, card := range d {
		fmt.Printf("%+v", card)
	}
	fmt.Println("==================================")

}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	cards := []string{}

	for _, card := range d {
		s := card.toString()
		cards = append(cards, s)
	}

	return strings.Join([]string(cards), ",")
}

func toDeck(s string) deck {
	cards := strings.Split(string(s), ",")
	deck := deck{}

	for _, card := range cards {
		c := toCard(card)
		deck = append(deck, c)
	}

	return deck
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func deckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	s := toDeck(string(bs))
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
