package main
import (
 "fmt"
 "math"
)
type geometry interface{
 area() float64
}

type circle struct{
	radius float64
}
type rect struct{
	a,b float64
}
type square struct{
	m float64
}
type trigon struct{
	i,j float64
}
func (c circle) area()float64{
	return math.Pi * c.radius * c.radius
}
func (r rect) area()float64{
	return r.a * r.b
}
func (s square) area()float64{
	return s.m * s.m
}
func (t trigon) area()float64{
	return (t.i * t.j) / 2
}
func print(g geometry){
 fmt.Println(g)
 fmt.Println(g.area())
}
func main(){
	var rad float64
	var aa float64
	var bb float64
	var mm float64
	fmt.Print(" Enter circle's  radius  : ")
	fmt.Scan(&rad)
	fmt.Println()
	fmt.Print(" Enter rect's and trigon's  parametr a : ")
	fmt.Scan(&aa)
	fmt.Println()
	fmt.Print(" Enter rect's  parametr b : ")
	fmt.Scan(&bb)
	fmt.Println()
	fmt.Print(" Enter square's  parametr  : ")
	fmt.Scan(&mm)
	fmt.Println()
	c := circle{radius: rad }
	r := rect{ a: aa, b: bb}
	s := square{ m: aa }
    t := trigon {i: aa, j: bb }
	print(c)
	print(r)
	print(s)
	print(t)

}
