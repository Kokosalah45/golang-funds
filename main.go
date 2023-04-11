package main

import "fmt"

func main() {
	dck := generateDeck()
	dck.shuffle(100)
	hand1 := dck.deal(7)
	fmt.Println(len(dck))
	hand2 := dck.deal(7)
	fmt.Println(len(dck))
	hand3 := dck.deal(7)
	fmt.Println(hand1)
	fmt.Println(hand2)
	fmt.Println(hand3)
	fmt.Println(len(dck))
	dck.saveToFile("./test.txt")
	loadedck := readFromFile("./test.txt")
	fmt.Printf("len original : %d / len loaded : %d", len(dck) , len(loadedck))


}
