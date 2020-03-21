package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// Create a new type of deck
// Which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "King", "Queen", "Jack", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// function (receiver) identifier(arguments) return
// the use of pointer is not necessary for non
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Shuffle/Randomize an Array
func (d *deck) shuffle() {
	check := [52]int{0}
	mainDeck := *d
	shuffled := deck{}

	for range mainDeck {
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(52)
		for check[randNum] != 0 {
			randNum = rand.Intn(52)
		}
		check[randNum] = 1
		shuffled = append(shuffled, mainDeck[randNum])
	}
	*d = shuffled
}

func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func (d *deck) dealHands() {
	for i := 0; i < 4; i++ {
		hand, remainingCards := d.deal(8)
		*d = remainingCards
		handFilename := fmt.Sprintf("hand%d.txt", i+1)
		hand.saveToFile(handFilename)
	}
}
