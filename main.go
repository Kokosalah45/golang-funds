package main

func main() {
	deck := deck{
		{value: "king", suit: "spades"},
		{value: "eleven", suit: "hearts"},
		{value: "three", suit: "diamonds"},
		{value: "four", suit: "clubs"},
		{value: "queen", suit: "hearts"}}
	deck.list()
	deck.shuffle(3)
	deck.list()
	deck.shuffle(3)
	deck.list()
}
