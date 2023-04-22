package main

type Cooker interface {
	Grill()
	Fry()
	Saute()
	Pastry()
	Roast()
}

type Saucier interface {
	Saute() string
}
type Griller interface {
	Grill() string
}
type Fryer interface {
	Fry() string
}
type Patisserie interface {
	Bake() string
}
type Roaster interface {
	Roast() string
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



type StaffBase struct {
 Name string
 Gender string
 Age int
 Ethinicty string
}

type staffType1 struct {
 StaffBase
 StaffFryer
 UniquePropertyX int
}

type staffType2 struct {
 StaffBase
 StaffFryer
 StaffGriller
 UniquePropertyY int
}

type staffType3 struct {
 StaffBase
 StaffFryer
 UniquePropertyZ int
}

type StaffFryer struct { }

func (s StaffFryer) Fry() string {
 return "frying..."
}

type StaffGriller struct { }

func (s StaffGriller) Grill() string {
 return "grilling..."
}


func main() {

	fryingStation(staffType1{}) // Works
	fryingStation(staffType2{}) // Works
	fryingStation(staffType3{}) // Works
	grillingStation(staffType2{}) // Works
	fryAndGrillAtTheSameTime(staffType2{})	//Also Works

}
