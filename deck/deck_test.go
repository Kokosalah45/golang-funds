package deck

import "testing"



func TestGenerateDeck(t *testing.T) {
	dck := generateDeck()
	if(len(dck) != 52){
		t.Errorf("Expected 52 got %d" , len(dck))
	}
}
func TestHandSize(t *testing.T) {
	dck := generateDeck()
	handSize := 7
	hand := dck.deal(7);

	if(len(hand) != handSize){
		t.Errorf("Expected %d got %d" ,handSize  , len(hand))
	}
}

func TestRemainingDeckSize(t *testing.T){
	dck := generateDeck()
	handSize := 7
	hand := dck.deal(7);

	if(len(dck) != 52 - len(hand)){
		t.Errorf("Expected %d got %d" , 52 - handSize  , len(dck))
	}
}

func TestShuffleDeck(t *testing.T){
	dck1 := generateDeck()
	dck2 := generateDeck();
	dck2.shuffle(100);
	differences := 0;

	for i , card := range dck1{
		if(card != dck2[i]){
			differences += 1
		}
	}

	if(differences == 0){
		t.Errorf("their is no differences between the two decks")
	}
	if(differences <= 10 ){
		t.Logf("their is not enough differences between the two decks should find better algorithm : %d" , differences)
		return 
	}
	t.Logf("differences : %d" , differences)

}