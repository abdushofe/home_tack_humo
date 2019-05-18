package main

import (
	"fmt"
	"time")
func print(a string){
	        dt := time.Now()
	        fmt.Println( " ___________________________________________________ ")
			fmt.Println( "              Your check  ")
			fmt.Println( " ___________________________________________________ ")
			fmt.Println( " Cost :    ", a )
            fmt.Println(" Paid Date :" ,dt.Format("01/02/2006 15:04:05"))
	        fmt.Println(" OperationStatus : Successful")
			fmt.Println( " ___________________________________________________ ")
}
func main() {
	var choose int
	fmt.Println(" <<<  Welcome to Pizza's restaurant >>>")
	fmt.Println( "________________________________________")
	fmt.Println( " <<< List type of Pizza : >>> ")
	list_manu := make(map [string]int)
	list_manu["Chicago Pizza - 25$ "] = 1
	list_manu["Greek Pizza - 50.5$ "] = 2
	list_manu["Colifornia Pizza - 45$ "] = 3
	list_manu["New York-Style Pizza - 30.4$ "] = 4
	list_manu["Neapolitan Pizza - 38.45$ "] = 5
	for i,value := range list_manu {
		fmt.Println(value, i)
		}
		fmt.Print( " Can you choose your favorite Pizza : ")
		fmt.Scan(&choose)
		for choose !=1 && choose !=2 && choose !=3 && choose !=4 && choose !=5 {
			for i,value := range list_manu {
				fmt.Println(value, i)
				}
				fmt.Print( " Sorry you can't choose anather type of Pizza \n Only choose in this list  <<Please>> !!! ")
				fmt.Scan(&choose)
		}
		if choose == 1 {print("Chicago Pizza - 25$ ")}
		if choose == 2 {print("Greek Pizza - 50.5$ ")}
		if choose == 3 {print("Colifornia Pizza - 45$ ")}
		if choose == 4 {print("New York-Style Pizza - 30.4$ ")}
		if choose == 5 {print("Neapolitan Pizza - 38.45$ ")}
		


		
		

		







}