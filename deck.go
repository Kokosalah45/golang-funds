package main

import (
	"fmt"
	"math/rand"
)

type card struct {
	suit  string
	value string
}	

type deck []card

func (dck deck) list() {

	for _, card := range dck {
		fmt.Println(card.value , "of" , card.suit)
	
	}
	fmt.Println("---------------");
}

func getRandomNumbers(upperBound int) (int,int) {
	x , y := rand.Intn(upperBound) , rand.Intn(upperBound)
	return x,y
}


func (dck *deck) shuffle(rounds int) {
	dckLen := len(*dck)
	temp := *dck;
	i := 1
	for i <= rounds {
		var randIndx1 , randIndx2 int 
		// needs optimizting if either bounds do something while in middle do different thing
	    for {
			randIndx1 , randIndx2 =  getRandomNumbers(dckLen)
			if (randIndx1 != randIndx2){
				break
			}
		}
		temp[randIndx1] , temp[randIndx2] = temp[randIndx2] , temp[randIndx1]
		i = i + 1;
	}
	
	
}