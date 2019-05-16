package main 
import "fmt"
func main(){
var a float64
var b float64
var a1 float64
var s float64
var s1 float64

fmt.Print("vvedite a =  ")
fmt.Scan( &a)
fmt.Println("")
fmt.Print("vvedite b =  ")
fmt.Scan( &b)
fmt.Println("")
fmt.Print("vvedite parmetr kvadrata a1 =  ")
fmt.Scan( &a1)
fmt.Println("")
s = a * b
s1 = s / (a1 * a1)
fmt.Println(s1)


} 