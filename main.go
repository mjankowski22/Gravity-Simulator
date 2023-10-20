package main

import (
	"fmt"
	"math"
)

const (
	b=float64(1)
	m=float64(1)
	g=float64(9.81)
	u=float64(0.5)
	alfa = math.Pi/4
)


func main () {
	x := float64(0) 
	x_prev := float64(0) 
	v:=float64(0) 
	v_prev := float64(0) 
	a:=float64(0) 
	a_prev :=float64(0) 

	h := math.Pow(10,-3)
	

	for i := 0; float64(i)*h < float64(10); i++ {
		x = x_prev + h*v_prev+math.Pow(h,2)/2*a_prev
		v = v_prev + h*a_prev
		a = -b/m*math.Pow(v_prev,2)-g*u*math.Cos(alfa)+g*math.Sin(alfa)
		fmt.Println(x)
		x_prev=x
		v_prev = v
		a_prev = a
	}
}

