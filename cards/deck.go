package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
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

// Shuffle/Randomize an Array
func (d deck) shuffle() {
	check := [52]int{0}
	for i := range d {
		rand.Seed(time.Now().UnixNano())
		newPosition := rand.Intn(len(d) - 1)
		for check[newPosition] != 0 {
			newPosition = rand.Intn(52)
		}
		check[newPosition] = 1
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func (d *deck) dealHands() {
	for i := 0; i < 4; i++ {
		hand, remainingCards := d.deal(8)
		*d = remainingCards
		handFilename := fmt.Sprintf("hand%d.txt", i+1)
		hand.saveToFile(handFilename)
	}
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	stringDeck := strings.Split(string(byteSlice), ", ")
	return deck(stringDeck)
}
