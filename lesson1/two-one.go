package main

import "fmt"

func main(){
	var x float64
	var y float64
fmt.Println("Введите число в сантиметрах : ")
fmt.Scanf("%2.f",&x)
y = x / 100

fmt.Println(" число в метрах:" , y)

}