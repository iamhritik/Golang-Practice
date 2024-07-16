package main

import "testing"

func TestNewDeck(t *testing.T) { // in the output it shows only one test and it's based on the no of test function u created
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Exptected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Exptected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}
