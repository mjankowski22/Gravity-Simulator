package main

import (
	"image/color"
	"math"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

const (
	windowWidth  = 800
	windowHeight = 600
	lineHeight = 200
	lineAngle    = math.Pi / 6 // Kąt nachylenia w radianach
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

	for !win.Closed() {
		win.Clear(color.Black)

		imd.Clear()

		
		

		
		// Narysuj równię pochyloną
		imd.Color = color.White
		imd.Push(pixel.V(x1, y1-lineHeight), pixel.V(x1+lineHeight/math.Tan(alfa), y1-lineHeight))
		imd.Push(pixel.V(x1, y1), pixel.V(x1+lineHeight/math.Tan(alfa), y1-lineHeight))
		imd.Push(pixel.V(x1,y1),pixel.V(x1,y1-lineHeight))

		imd.Line(5)

		imd.Color = color.RGBA{255,0,0,0}
		d:=40.0
		


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
		x += 2*math.Cos(alfa)
		y -= 2*math.Sin(alfa)
		time.Sleep(time.Second / 60) // Ograniczenie liczby klatek na sekundę
	}
}


const (
	b=float64(1)
	m=float64(1)
	g=float64(9.81)
	u=float64(0.5)
	alfa = math.Pi/6
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
		x_prev=x
		v_prev = v
		a_prev = a
	}

	pixelgl.Run(run)
}

