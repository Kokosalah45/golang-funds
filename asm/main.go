package main

import (
	"fmt"
)

func main() {
	
	p := fmt.Println

	/// Arrays -> Sent by value 

	var a [3]int
	b := [3]int{}
	c := [...]int{0, 0, 0}
	// ALL ARE THE SAME


	x := [3]int{1,2,3}
	y := [3]int{}

	y = x // now all content of x copied into y however x and y are not the same they are two different arrays

	// z := [2]int{}
	// z = x // type mismatch


	p(a, b, c , x ,y)



	/// Slices -> Sent by reference 

	//---------
		var e []int  // nil
		f := []int{} // empty with length of 0 and capacity of 0
	//---------

	//make()
	//---
		// NOTE make() the capacity takes the same value as the length unless it's specified and the capacity has to be 
		//more than or equal to the specified length

		//---------
			g := make([]int , 0) // empty with length of 0 and capacity of 0
			h := make([]int , 0 , 0) // empty with length of 0 and capacity of 0
		//---------

		//---------
			i := make([]int, 5)     // non empty initialized with length of 5 zeros and capacity of 5 same as -> make([]int, 5 , 5)
			j := make([]int, 5, 10) // non empty initialized with length of 5 zeros and capacity of 10
		//---------
		p(e, f, g , h ,i , j)
	//---


	//append()

	//---------
		//---------
			e = append(e, 3) // i can append to nil slice thats fine
			f = append(f, 3)  
			g = append(g, 3)  
			h = append(h, 3)  
			// everytime you append to anyone of those a new memory block is created and the elements are copied to the new block
		//---------

		//---------
			i = append(i, 3) 
			// allocates a new place in the memory block as the length exceeded the capacity and copies all the data in the old memory block onto the new one
			j = append(j, 3) 
			// a new block in memory doesn't get allocated as the capacity of h is 10, and the length after appending is 6 -> length 6 > capacity 10
		//---------
		p(e, f, g , h ,i , j)
	//---------

	//copy()

	//---------
		// copy is only available for slices 
		// copy copies from the src to the destination slice elements equal to the LENGTH of the src slice
		// so if a src slice has 4 elements and the destination slice has a LENGTH if only 3 it gets the first 3 elements only 
		// if it's more or equal its take all of the elements 
		// it overwrites the content of the destination slice

		// cases 
		k := make([]int , 3) // length and capacity are equal
		l := make([]int , 3 , 5)	// length and capacity are not equal
		src := []int{1,2,3,4,5,6}
		copy(k , src)
		copy(l,src)
		p(l,k) // same output it doesn't matter regarding the capacity
		// again the length of the src doesn't matter what matters is the length of destination 
		// less than the src ? take what's in the src until the max length is reached
		// more than or equal to the src ? then take all what's in the src 

	//---------

	/// Maps -> Sent by reference
	// maps can't be compared to one another same as slices 
	// i can compare them to nil of course like slices however  

	var m map[string]int // nil (we can only read from it and we can never write so it's not useful at all  we never declare a map like this)
	n := make(map[string]int) // not nill just empty (can read and write)  (we declare it with make() )

	mm := m["x"] // despite it being a nil map i can read and it returns the default value 0 as the declared value type is int 
	nn := n["y"] // returns the default value 0 as the declared value type is int

	m["x"] = 1 // panic !
	n["y"] = 1 // the key is created and assigned the value 1 
	 
	if val , ok :=  m["x"]; ok {
		fmt.Println(val)
	}
	// another way to index maps 
	p(m,n , mm , nn )





	


}
