package main
import "fmt"
func main(){
var n int64
var s int64
var m int64
var ch int64
fmt.Println( " vvedite n ") 
fmt.Scan( &n)
ch = n / 3600
m = (n % 3600)/60
s = (n % 60) % 60
fmt.Println(ch,"chas",m,"min",s,"sec")
}