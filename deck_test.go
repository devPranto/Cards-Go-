package main

import (
	"testing"
)

func TestDeck(t *testing.T){
	d:= newDeck()
	if len(d)!= 52 {
		t.Error("Something is wrong ",len(d))
	}
}