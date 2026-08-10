package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
		d:=newDeck()
	if len(d)!=16{
	t.Errorf("expected result is 16 but  got %v",len(d))
	}
	if d[0]!="Ace of Spades"{
		t.Errorf("expected value is Ace of Spades but yu got %v ",d[0])
	}
	if d[len(d)-1]!="Four of Clubs"{
		t.Errorf("Expected value is Four of Clubs but ya got %v" ,d[len(d)-1])
	}
}
func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_testingdeck")
	deck:=newDeck()
	deck.saveToFile("_testingdeck")
	loaded:=newDeckFromFile("_testingdeck")
	if len(loaded)!=16 {
		t.Errorf("expected value is 16 but you got %v",len(loaded))
	}
	os.Remove("_testingdeck")
	
}






