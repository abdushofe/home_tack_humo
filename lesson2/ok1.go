package main 
import  (
	"fmt"

)
func print(a *int, b *int , c *string ) {
	fmt.Println("---------------------------------------------")
	fmt.Println(" Opertor : ", *c)
	fmt.Println(" Number : ", *a)
	fmt.Println(" Sum :", *b)
	fmt.Println(" OperationStatus : Successful")
	fmt.Println("---------------------------------------------")

}
func main(){
fmt.Println("Welcome to out pay terminal ")
fmt.Println(" Choose your operator : ")

x := make(map[string]int)
x["Megafon"] = 1
x["Beeline"] = 2
x["Babilon"] = 3
x["Tcell"] = 4
for i,value := range x {
	fmt.Println(value, i)
	}
	fmt.Print(" Choose your operator : ")
	var choose  int
	var phone  int
	var sum int 
	fmt.Scan(&choose)
	if choose == x["Megafon"] {
		fmt.Print(" Enter your number (Megafon) : " )
		fmt.Scan(&phone)
		if ( phone / 10000000 == 90 || phone / 10000000 == 88  || phone / 10000000 == 55){
			fmt.Print(" Enter sum:")
			fmt.Scan(&sum)
			l := "Megafon"
			if sum > 0 {
			print(&phone,&sum, &l)
			} else {
				fmt.Println(" // Please, enter fight sum > 0 ")
			}

	} else {
		fmt.Println(" Please, enter right operator number ")

	}}
	
	}