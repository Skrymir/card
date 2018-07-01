package main

import (
	"os"
	"testing"
)

func TestNewDeck_Count(t *testing.T) {
	d := newDeck()
	exp := 52
	if len(d) != exp {
		t.Errorf("Expected deck length of [%v] but got [%v]", exp, len(d))
	}

	expFirst := "Ace of Spades"
	if d[0] != expFirst {
		t.Errorf("Expected first card [%v] got [%v]", expFirst, d[0])
	}

	expLast := "Ten of Diamonds"
	if d[len(d)-1] != expLast {
		t.Errorf("Expected last card [%v] got [%v]", expLast, d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	df := "_decktesting"
	os.Remove(df)

	d := newDeck()
	d.saveToFile(df)

	expCount := 52
	if len(d) != expCount {
		t.Errorf("Expected loaded from disk deck length of [%v] but got [%v]", expCount, len(d))
	}

	os.Remove(df)

}
