package main

type Cooker interface {
	Grill()
	Fry()
	Saute()
	Pastry()
	Roast()
}

type Saucier interface {
	Saute() bool
}
type Griller interface {
	Grill() bool
}
type Fryer interface {
	Fry() bool
}
type Patisserie interface {
	Bake() bool
}
type Roaster interface {
	Roast() bool
}

type FryerGriller interface {
	Griller
	Fryer
}

func fryingStation(f Fryer) {
	f.Fry()
}

func grillingStation(g Griller) {
	g.Grill()
}

func fryAndGrillAtTheSameTime(fg FryerGriller) {
	//concept of channels is not covered here however imagine that they are
	// happening in parallel
	fryingStation(fg)
	grillingStation(fg)
	// will work as fg is both a Fryer type and Griller Type
}

func main() {

}