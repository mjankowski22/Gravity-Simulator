package main

import (
	"fmt"
	"image/color"
	"math"
	"os"
	"strconv"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

const (
	windowWidth  = 800.0
	windowHeight = 600.0
	lineHeight = 300.0
)


var x_result []float64
var v_result []float64
var a_result []float64	
var t_result []float64
var alfa float64

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
	y1 := float64( windowHeight / 2)+100

	x := x1 +3
	y := y1 +3

	
	var draw_result []float64

	for i := 0; i < len(x_result); i++ {
		draw_result = append(draw_result,x_result[i]/x_result[len(x_result)-1]*lineHeight/math.Sin(alfa))
		
	}

	i := 0
	for !win.Closed() {
		win.Clear(color.Black)
		d:=30.0
		
		imd.Clear()

		if(y-d*math.Sin(alfa)>y1-lineHeight){
			x = x1 +3 + draw_result[i]*math.Cos(alfa)
			y = y1 +3 - draw_result[i]*math.Sin(alfa)
			i+=1
		}
		fmt.Println(t_result[i])
		

		

		
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
	plot_results(t_result,x_result,v_result,a_result)
}





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
		var g,m,b,u,height float64
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
		alfa = alfa*math.Pi/180
		height,err=strconv.ParseFloat(input_height.Text,64)
		if err != nil {
			fmt.Println("Erorr: ",err)
		}
		
		x_result,v_result,a_result,t_result=simulation(g,m,u,b,alfa,height)
		
		w.Close()
		
	}

	submit := widget.NewButton("Simulation",handle_submit)

	grid:= container.New(layout.NewGridLayout(2),input_gravity_label,input_mass_label,input_gravity,input_mass,input_friction_label,input_resistance_label,input_friction,input_resistance,input_alfa_label,input_height_label,input_alfa,input_height,submit)
	w.SetContent(grid)

	w.ShowAndRun()
}

func simulation(g,m,u,b,alfa,height float64) ([]float64,[]float64,[]float64,[]float64){
	x := float64(0) 
	x_prev := float64(0) 
	v:=float64(0) 
	v_prev := float64(0) 
	a:=float64(0) 
	a_prev :=float64(0) 
	var x_result []float64
	var v_result []float64
	var a_result []float64
	var t_result []float64

	for i := 0; x*math.Sin(alfa)< height; i++ {
		x = x_prev + h*v_prev+math.Pow(h,2)/2*a_prev
		v = v_prev + h*a_prev
		a = -b/m*math.Pow(v_prev,2)-g*u*math.Cos(alfa)+g*math.Sin(alfa)
		x_result = append(x_result,x)
		v_result = append(v_result,v)
		a_result = append(a_result,a)
		t_result = append(t_result,float64(i)*h)
		x_prev=x
		v_prev = v
		a_prev = a
	}
	return x_result,v_result,a_result,t_result
}


func plot_results(t_result,x_result,v_result,a_result []float64) {

	// Tworzenie nowego wykresu
	p1 := plot.New()
	p2 := plot.New()
	p3 := plot.New()

	// Tworzenie serii danych
	lineData1 := make(plotter.XYs, len(t_result))
	for i := range t_result {
		lineData1[i].X = t_result[i]
		lineData1[i].Y = x_result[i]
	}
	lineData2 := make(plotter.XYs, len(t_result))
	for i := range t_result {
		lineData2[i].X = t_result[i]
		lineData2[i].Y = v_result[i]
	}
	lineData3 := make(plotter.XYs, len(t_result))
	for i := range t_result {
		lineData3[i].X = t_result[i]
		lineData3[i].Y = a_result[i]
	}
	line1, err := plotter.NewLine(lineData1)
	if err != nil {
		fmt.Println(err)
		return
	}
	line2, err := plotter.NewLine(lineData2)
	if err != nil {
		fmt.Println(err)
		return
	}
	line3, err := plotter.NewLine(lineData3)
	if err != nil {
		fmt.Println(err)
		return
	}
	p1.Add(line1)
	p1.X.Label.Text = "Time [s]"
	p1.Y.Label.Text = "Distance [m]"
	p2.Add(line2)
	p2.X.Label.Text = "Time [s]"
	p2.Y.Label.Text = "Velocity [m/s]"
	p3.Add(line3)
	p3.X.Label.Text = "Time [s]"
	p3.Y.Label.Text = "Acceleration [m/s^2]"



	plots := make([][]*plot.Plot, 3)
	plots[0] =make([]*plot.Plot, 1)
	plots[1] = make([]*plot.Plot, 1)
	plots[2] = make([]*plot.Plot, 1)
	plots[0][0] = p1
	plots[1][0] = p2
	plots[2][0] = p3


	img := vgimg.New(vg.Points(500), vg.Points(600))
    dc := draw.New(img)




	t := draw.Tiles{
        Rows: 3,
        Cols: 1,
    }

	canvases := plot.Align(plots, t, dc)
	plots[0][0].Draw(canvases[0][0])
	plots[1][0].Draw(canvases[1][0])
	plots[2][0].Draw(canvases[2][0])

	w, err := os.Create("aligned.png")
    if err != nil {
        panic(err)
    }

    png := vgimg.PngCanvas{Canvas: img}
    if _, err := png.WriteTo(w); err != nil {
        panic(err)
    }
}



func main () {
	parameterWindow()
	pixelgl.Run(run)
}

