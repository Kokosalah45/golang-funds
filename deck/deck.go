package deck

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

var suits = [4]string{"Diamonds" , "Spades" , "Hearts" , "Clubs"}
var values = [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten" , "King", "Jack", "Queen"}	


func GenerateDeck() deck {
	dck := deck{}
	for _, suit := range suits{
		for _, value := range values{
			dck = append(dck, card{suit , value})
		}
	}
	return dck
}

func (dck deck) List() {
	for _, card := range dck {
		fmt.Println(card.Suit + " of " + card.Value)
	
	}
	fmt.Println("---------------");
}

func getRandomNumbers(upperBound int , isEqualAllowed bool ) [2]int {
	randNums := [2]int{rand.Intn(upperBound) , rand.Intn(upperBound)}
	if (isEqualAllowed ||randNums[0] != randNums[1] ){
		return randNums
	}
	
	return getRandomNumbers(upperBound ,  isEqualAllowed)	
}
func (dck deck) Shuffle(rounds int) {
	dckLen := len(dck)
	i := 1
	for i <= rounds {
		randIndxs := getRandomNumbers(dckLen , false);
		dck[randIndxs[0]] , dck[randIndxs[1]] = dck[randIndxs[1]] , dck[randIndxs[0]]
		i = i + 1;
	}
}

func (dck *deck) Deal(handSize int) deck {
	hand := (*dck)[:handSize]
	// here i need to send by reference as i want to change the value of the deck instance itself 
	*dck = (*dck)[handSize:]
	return hand
} 


func (dck deck) stringify() []byte {
	jsonDeck , err := json.Marshal(dck)
	if (err != nil){
		panic(err)
	}
	return jsonDeck
}

func Parse(stringifiedDeck []byte) deck{
	var dck deck ;
	err := json.Unmarshal(stringifiedDeck,&dck)
	if (err != nil){
		panic(err)
	}
	return dck
}
func (dck deck) SaveToFile(filename string){
	stringDeck := dck.stringify()
	err := os.WriteFile(filename , stringDeck , 0644)
	if (err != nil){
		panic(err)
	}
}
func ReadFromFile(filename string) deck {

	stringifiedText , err := os.ReadFile(filename)
	if (err != nil){
		panic(err)
	}
	return Parse(stringifiedText)
}