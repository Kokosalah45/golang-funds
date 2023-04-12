package pointer

import "fmt"
 
func Run() {
 name := "bill"
 namePointer := &name
 // name has an address x
 // i get store the value of that address by a pointer y
 // the pointer itself has an address z
 fmt.Printf("name value : %s /  name address : %d / name pointer address : %d \n",  name , namePointer , &namePointer)
 printPointer(namePointer) 
}
 
func printPointer(pointer *string) {
 // name has an address x
 // i get store the value of that address by a pointer y
 // the pointer itself has an address z' WAIT WHY ? z != z' because i copy the value of address stored in the pointer i create a new pointer 
 // and store the value of that address in it simple as that  
 fmt.Printf("name value : %s / name address : %d / name pointer address : %d \n" , *pointer , pointer , &pointer)
}

