package main 
import  (
	"fmt"
	"time"
)
func print(a *int, b *int , c *string ) {
	
	fmt.Println("---------------------------------------------")
	fmt.Println(" Opertor : ", *c)
	fmt.Println(" Number : ", *a)
	fmt.Println(" Sum :", *b)
	dt := time.Now()
    fmt.Println(" Paid Date :" ,dt.Format("01/02/2006 15:04:05"))
	fmt.Println(" OperationStatus : Successful")
	fmt.Println("---------------------------------------------")
}
// Функция print для вывода 
func main(){
fmt.Println("Welcome to out pay terminal ")
fmt.Println(" Choose your operator : ")

x := make(map[string]int)
x["Megafon"] = 1
x["Beeline"] = 2
x["Babilon"] = 3
x["Tcell"] = 4
for i,value := range x {
	fmt.Println(value, i)}
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
				fmt.Println(" // Please, enter fight sum > 0 ")}
	} else {
		fmt.Println(" Please, enter right operator number ")}}
	// Генератор отчета для компания Megafon
	if choose == x["Beeline"] {
		fmt.Print(" Enter your number (Beeline) : " )
		fmt.Scan(&phone)
		if ( phone / 10000000 == 91 ){
			fmt.Print(" Enter sum:")
			fmt.Scan(&sum)
				l := "Beeline"
			if sum > 0 {
			print(&phone,&sum, &l)
			} else {
				fmt.Println(" // Please, enter fight sum > 0 ")}
	} else {
		fmt.Println(" Please, enter right operator number ")}}
	// Генератор отчета для компания Beeline
	if choose == x["Tcell"] {
		fmt.Print(" Enter your number (Tcell) : " )
		fmt.Scan(&phone)
		if ( phone / 10000000 == 92 || phone / 10000000 == 93 || phone / 10000000 == 50  ){
			fmt.Print(" Enter sum:")
			fmt.Scan(&sum)
				l := "Tcell"
			if sum > 0 {
			print(&phone,&sum, &l)
			} else {
				fmt.Println(" // Please, enter fight sum > 0 ")}
	} else {
		fmt.Println(" Please, enter right operator number ")}}
///// Генератор отчета для компания Tcell
if choose == x["Babilon"] {
	fmt.Print(" Enter your number (Babilon) : " )
	fmt.Scan(&phone)
	if ( phone / 10000000 == 91 || phone / 10000000 == 98  ){
		fmt.Print(" Enter sum:")
		fmt.Scan(&sum)
		l := "Babilon"
			if sum > 0 {
			print(&phone,&sum, &l)
		} else {
			fmt.Println(" // Please, enter fight sum > 0 ")}
} else {
	fmt.Println(" Please, enter right operator number ")
// // Генератор отчета для компания Babilon
}}}
