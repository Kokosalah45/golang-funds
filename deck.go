package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

type card struct {
	Suit  string `json:"suit"`
	Value string `json:"value"`
}	

type deck []card



func generateDeck() deck {
	dck  := deck{}
	suits := [4]string{"Diamonds" , "Spades" , "Hearts" , "Clubs"}
 	values := [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten" , "King", "Jack", "Queen"}	
	for _, suit := range suits{
		for _, value := range values{
			card := card{suit , value}
			dck = append(dck, card)
		}
	}
	return dck
}

func (dck deck) list() {
	for _, card := range dck {
		fmt.Println(card.Suit + " of " + card.Value)
	
	}
	fmt.Println("---------------");
}

func getRandomNumbers(upperBound int) [2]int {
	randNums := [2]int{rand.Intn(upperBound) , rand.Intn(upperBound)}
	return randNums
}

func (dck *deck) deal(handSize int) deck {
	hand := (*dck)[:handSize]
	*dck = (*dck)[handSize:]
	return hand
} 



func (dck *deck) shuffle(rounds int) {
	dckLen := len(*dck)
	temp := *dck;
	i := 1
	for i <= rounds {
		var randIndxs [2]int
		// needs optimizting if either bounds do something while in middle do different thing
	    for {
			randIndxs =  getRandomNumbers(dckLen)
			if (randIndxs[0] != randIndxs[1]){
				break
			}
		}
		temp[randIndxs[0]] , temp[randIndxs[1]] = temp[randIndxs[1]] , temp[randIndxs[0]]
		i = i + 1;
	}
}

func (dck deck) saveToFile(filename string){
	stringDeck := dck.stringify()
	err := os.WriteFile(filename , stringDeck , 0644)
	if (err != nil){
		panic(err)
	}
}
func readFromFile(filename string) deck {
	
	stringifiedText , err := os.ReadFile(filename)
	
	if (err != nil){
		panic(err)
	}
	return parse(stringifiedText)
}



func (dck deck) stringify() []byte {
	jsonDeck , err := json.Marshal(dck)
	if (err != nil){
		panic(err)
	}
	return jsonDeck
}

func parse(stringifiedDeck []byte) deck{
	var dck deck ;
	err := json.Unmarshal(stringifiedDeck,&dck)
	 if (err != nil){
		panic(err)
	}
	fmt.Println(dck)
	return dck
}