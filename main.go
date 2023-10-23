package main

import (
	"fmt"
	"image/color"
	"math"
	"strconv"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

const (
	windowWidth  = 800.0
	windowHeight = 600.0
	lineHeight = 200.0// Kąt nachylenia w radianach
)


func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Animacja Równi Pochylonej",
		Bounds: pixel.R(0, 0, windowWidth, windowHeight),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	
	x1 := float64( windowWidth / 2)-300
	y1 := float64( windowHeight / 2)

	x := x1 +3
	y := y1 +3

	i := 0
	fmt.Println(x_result)
	for !win.Closed() {
		win.Clear(color.Black)
		d:=30.0
		
		imd.Clear()

		if(y-d*math.Sin(alfa)>y1-lineHeight){
			x = x1 +3 + x_result[i]*math.Cos(alfa)
			y = y1 +3 - x_result[i]*math.Sin(alfa)
			i+=1
		}
		

		

		
		// Narysuj równię pochyloną
		imd.Color = color.White
		imd.Push(pixel.V(x1, y1-lineHeight), pixel.V(x1+lineHeight/math.Tan(alfa), y1-lineHeight))
		imd.Push(pixel.V(x1, y1), pixel.V(x1+lineHeight/math.Tan(alfa), y1-lineHeight))
		imd.Push(pixel.V(x1,y1),pixel.V(x1,y1-lineHeight))

		imd.Line(5)

		imd.Color = color.RGBA{255,0,0,0}
		

		x2 := x + d * math.Cos(alfa)
		y2 := y - d * math.Sin(alfa)

		x3 := x + d * math.Sin(alfa)
		y3 := y + d * math.Cos(alfa)

		x4 := x + d * (math.Cos(alfa) + math.Sin(alfa))
		y4 := y - d * (math.Sin(alfa) - math.Cos(alfa))
		
		imd.Push(pixel.V(x,y),pixel.V(x2,y2))
		imd.Push(pixel.V(x2,y2),pixel.V(x4,y4))
		imd.Push(pixel.V(x4,y4),pixel.V(x3,y3))
		imd.Push(pixel.V(x3,y3),pixel.V(x,y))
		

		imd.Line(5)
		
		
		

		imd.Draw(win)

		win.Update()
		
		time.Sleep(time.Second/100) 
		fmt.Println(float64(i)/100)
		
	}
}

var x_result []float64
var v_result []float64
var a_result []float64

const (
	b=float64(0)
	m=float64(1)
	g=float64(20)
	u=float64(0)
	alfa = math.Pi/6
)


var h = math.Pow(10,-2)


func parameterWindow () {

	
	app := app.New()
	w := app.NewWindow("Hello World")
	
	
	input_gravity := widget.NewEntry()
	input_gravity_label :=widget.NewLabel("Gravity acceleration [m/s^2]")
	input_gravity.Text = "9.81"

	input_mass := widget.NewEntry()
	input_mass_label :=widget.NewLabel("Mass [kg]")
	input_mass.Text = "1"

	input_friction := widget.NewEntry()
	input_friction_label :=widget.NewLabel("Friction parameter")
	input_friction.Text ="0.1"

	input_resistance := widget.NewEntry()
	input_resistance_label :=widget.NewLabel("Resistance paramater [jakis parametr]")
	input_resistance.Text ="0.1"

	input_alfa := widget.NewEntry()
	input_alfa_label := widget.NewLabel("Angle [deg]")
	input_alfa.Text = "30"

	input_height :=widget.NewEntry()
	input_height_label := widget.NewLabel("Height [m]")
	input_height.Text = "2"

	handle_submit:= func(){
		var g,m,b,u,alfa float64
		var err error
		g,err=strconv.ParseFloat(input_gravity.Text,64)
		if err != nil {
			fmt.Println("Erorr: ",err)
		}
		m,err=strconv.ParseFloat(input_mass.Text,64)
		if err != nil {
			fmt.Println("Erorr: ",err)
		}
		u,err=strconv.ParseFloat(input_friction.Text,64)
		if err != nil {
			fmt.Println("Erorr: ",err)
		}
		b,err=strconv.ParseFloat(input_resistance.Text,64)
		if err != nil {
			fmt.Println("Erorr: ",err)
		}
		alfa,err=strconv.ParseFloat(input_alfa.Text,64)
		if err != nil {
			fmt.Println("Erorr: ",err)
		}
		
		x_result,v_result,a_result=simulation(g,m,u,b,alfa*math.Pi/360)
		w.Close()
		
	}

	submit := widget.NewButton("Simulation",handle_submit)

	grid:= container.New(layout.NewGridLayout(2),input_gravity_label,input_mass_label,input_gravity,input_mass,input_friction_label,input_resistance_label,input_friction,input_resistance,input_alfa_label,input_height_label,input_alfa,input_height,submit)
	w.SetContent(grid)

	w.ShowAndRun()
}

func simulation(g,m,u,b,alfa float64) ([]float64,[]float64,[]float64){
	fmt.Println(g,m,u,b,alfa)
	x := float64(0) 
	x_prev := float64(0) 
	v:=float64(0) 
	v_prev := float64(0) 
	a:=float64(0) 
	a_prev :=float64(0) 
	var x_result []float64
	var v_result []float64
	var a_result []float64

	for i := 0; x*10*math.Sin(alfa)< lineHeight; i++ {
		x = x_prev + h*v_prev+math.Pow(h,2)/2*a_prev
		v = v_prev + h*a_prev
		a = -b/m*math.Pow(v_prev,2)-g*u*math.Cos(alfa)+g*math.Sin(alfa)
		x_result = append(x_result,x*20)
		v_result = append(v_result,v)
		a_result = append(a_result,a)
		x_prev=x
		v_prev = v
		a_prev = a
	}
	return x_result,v_result,a_result
}


func main () {
	

	parameterWindow()
	pixelgl.Run(run)
}

