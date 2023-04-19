package deck

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

func getCardSuits() [4]string {
	return [4]string{"Diamonds" , "Spades" , "Hearts" , "Clubs"}
}

func getCardValues() [4]string {
	return [4]string{"Diamonds" , "Spades" , "Hearts" , "Clubs"}
}


type card struct {
	Suit  string `json:"suit"`
	Value string `json:"value"`
}

type cardList []card


type deck struct{
	cards cardList
}

func NewDeck() *deck {
	dck := &deck{}
	suits := getCardSuits()
	values := getCardValues()
	for _, suit := range suits{
		for _, value := range values{
			(*dck).cards = append((*dck).cards, card{suit , value})
		}
	}
	return dck
}



func (dck deck) List() {
	for _, card := range dck.cards {
		fmt.Println(card.Suit + " of " + card.Value)
	
	}
	fmt.Println("---------------");
}

func (dck deck) Count() {
	fmt.Println(len(dck.cards))
}

func (dck deck) Shuffle(rounds int) {
	
	dckLen := len(dck.cards)
	i := 1
	for i <= rounds {
		randIndxs := getRandomNumbers(dckLen , false);
		dck.Swap(randIndxs[0] , randIndxs[1])
		i = i + 1;
	}
}
func (dck *deck) Deal(handSize int) []card  {
	hand := (*dck).cards[:handSize]
	// here i need to send by reference as i want to change the value of the deck instance itself 
	dck.cards = (*dck).cards[handSize:]
	return hand
} 
func (dck deck) Swap(i int , j int ){
	dck.cards[i] , dck.cards[j] = dck.cards[j] , dck.cards[i]
}





// --------------------------------------------------

func Parse(stringifiedDeck []byte) deck{
	var dck deck ;
	err := json.Unmarshal(stringifiedDeck,&dck)
	if (err != nil){
		panic(err)
	}
	return dck
}
func SaveToFile(filename string , bs []byte ){
	err := os.WriteFile(filename , bs , 0644)
	if (err != nil){
		panic(err)
	}
}
func ReadFromFile(filename string) []byte {
	stringifiedText , err := os.ReadFile(filename)
	if (err != nil){
		panic(err)
	}
	return stringifiedText
}

func getRandomNumbers(upperBound int , isEqualAllowed bool ) [2]int {
	randNums := [2]int{rand.Intn(upperBound) , rand.Intn(upperBound)}
	if (isEqualAllowed ||randNums[0] != randNums[1] ){
		return randNums
	}	
	
	return getRandomNumbers(upperBound ,  isEqualAllowed)	
}	


func stringify(v any) []byte {
	bs , err := json.Marshal(v)
	if (err != nil){
		panic(err)
	}
	return bs
}
// --------------------------------------------------
