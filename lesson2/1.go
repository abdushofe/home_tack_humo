package main

import "fmt"

func main() {
	
	var x1 int
	var x2 int 
	x1  = 3 
	x2 = 5 
	fmt.Printf (" x1 = %d, x2 = %d \n", x1, x2)
	fmt.Printf( &x1 , &x2 )
	fmt.Printf( "swapping ...")
	swap(&x1,&x2)
	fmt.Printf( x1 , x2 )
}
func swap(a *int, b *int){
var c = int
c = *a 
*a = *b
*b = c
}


