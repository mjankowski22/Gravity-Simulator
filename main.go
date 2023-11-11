package main

import (
	"fmt"
	"image/color"
	"math"
	"os"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

const (
	windowWidth  = 800.0
	windowHeight = 600.0
	lineHeight = 300.0
)


var x_result []float64
var y_result []float64
var vx_result []float64
var vy_result []float64
var ax_result []float64	
var ay_result []float64	
var t_result []float64
var gravity_inputs []Gravity_Input
var gravity_sources []Gravity_Source
var height float64



func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Animacja Równi Pochylonej",
		Bounds: pixel.R(0, 0, windowWidth, windowHeight),
		VSync:  true,
	}

	alfa:= math.Pi/6
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	
	x1 := 0.0+5
	y1 := lineHeight+5

	

	
	var draw_result_x []float64
	var draw_result_y []float64
	
	meters_to_pixels := lineHeight/height

	for i := 0; i < len(x_result); i++ {
		draw_result_x = append(draw_result_x,x_result[i]*meters_to_pixels+5)
		draw_result_y = append(draw_result_y,y_result[i]*meters_to_pixels+5)
		
	}
	

	i := 0
	for !win.Closed() {
		win.Clear(color.Black)

		d:=30.0
		
		imd.Clear()
		var x float64
		var y float64
		if(y_result[i]>0){
			x =  draw_result_x[i] + 3
			y =  draw_result_y[i] +3
		
			i+=1
		}
		
		for _,gravity_src := range gravity_sources{
			imd.Color = color.RGBA{0,255,0,0}
			imd.Push(pixel.V(gravity_src.x*meters_to_pixels,gravity_src.y*meters_to_pixels))
			imd.Circle(20,5)
		}
		
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
		
		
		
	}
	
	plot_results(t_result,x_result,y_result,ax_result)
}





var h = math.Pow(10,-2)


type Gravity_Input struct{
	input_gravity *widget.Entry
	input_x *widget.Entry
	input_y *widget.Entry
}

type Gravity_Source struct{
	mass float64
	x float64
	y float64
}

func parameterWindow () {

	app := app.New()
	w := app.NewWindow("Choose parameters")
	
	var grid *fyne.Container
	
	

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

	input_velocity :=widget.NewEntry()
	input_velocity_label := widget.NewLabel("Velocity[m/s^2]")
	input_velocity.Text = "2"

	handle_submit:= func(){
		var g,m,b,u,alfa,v float64
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
		v,err=strconv.ParseFloat(input_velocity.Text,64)
		if err != nil {
			fmt.Println("Erorr: ",err)
		}

		for _,gravity_input := range gravity_inputs{
			mass,err := strconv.ParseFloat(gravity_input.input_gravity.Text,64)
			if err != nil {
				fmt.Println("Erorr: ",err)
			}
			x,err := strconv.ParseFloat(gravity_input.input_x.Text,64)
			if err != nil {
				fmt.Println("Erorr: ",err)
			}
			y,err := strconv.ParseFloat(gravity_input.input_y.Text,64)
			if err != nil {
				fmt.Println("Erorr: ",err)
			}
			gravity_sources = append(gravity_sources, Gravity_Source{
				mass: mass,
				x: x,
				y: y,
			})
		}
		
		x_result,y_result,t_result=simulation(g,m,u,b,alfa,v)
		
		w.Close()
		app.Quit()
	}
	handle_add := func(){
		
		input_number := strconv.Itoa(len(gravity_sources)+1)
		input_gravity :=widget.NewEntry()
		input_gravity_label := widget.NewLabel(input_number+" Mass [kg]")
		input_x :=widget.NewEntry()
		input_x_label := widget.NewLabel(input_number+" X[m]")
		input_y :=widget.NewEntry()
		input_y_label := widget.NewLabel(input_number+" Y[m]")

		input := Gravity_Input{
			input_gravity:input_gravity,
			input_x:input_x,
			input_y:input_y,
		}
		
		gravity_inputs = append(gravity_inputs,input)
		
		grid.Add(input_gravity_label)
		grid.Add(input_x_label)
		grid.Add(input_gravity)
		grid.Add(input_x)
		grid.Add(input_y_label)
		grid.Add(widget.NewLabel(""))
		grid.Add(input_y)
		grid.Add(widget.NewLabel(""))

		

	}
	add_gravity_src := widget.NewButton("Add gravity source",handle_add)
	submit := widget.NewButton("Simulation",handle_submit)
	

	grid = container.New(layout.NewGridLayout(2),input_gravity_label,input_mass_label,input_gravity,input_mass,input_friction_label,input_resistance_label,input_friction,input_resistance,input_alfa_label,input_height_label,input_alfa,input_height,input_velocity_label,widget.NewLabel(""),input_velocity,widget.NewLabel(""),submit,add_gravity_src)
	

	w.SetContent(grid)

	w.ShowAndRun()
}

func simulation(g,m,u,b,alfa,v float64) ([]float64,[]float64,[]float64){
	x_prev := 0.0
	y_prev := height
	vx_prev := v*math.Cos(alfa)
	vy_prev:= -v*math.Sin(alfa)
	Fx_sources := 0.0
	Fy_sources := 0.0
	for _,gravity_source := range gravity_sources{
		r := math.Sqrt(math.Pow(gravity_source.x-x_prev,2)+math.Pow(gravity_source.y-y_prev,2))
		F := 6.67430*math.Pow(10,-11)*m*gravity_source.mass/math.Pow(r,2)
		fmt.Println(F/m)
		Fx_sources += F*(gravity_source.x-x_prev)/r 
		Fy_sources += F*(gravity_source.y-y_prev)/r 
	}
	Fx := -b*(math.Pow(vx_prev,2)+math.Pow(vy_prev,2))*math.Cos(alfa) + Fx_sources
	Fy := b*(math.Pow(vx_prev,2)+math.Pow(vy_prev,2))*math.Sin(alfa) - m*g +Fy_sources
	diff := Fy*math.Cos(alfa)+Fx*math.Sin(alfa)
	if diff<0{
		Fx += -diff*math.Sin(alfa)+diff*u*math.Cos(alfa) 
		Fy += -diff*math.Cos(alfa) - diff*u*math.Sin(alfa)
	}
	ax_prev := Fx/m
	ay_prev := Fy/m

	x_result = append(x_result, x_prev)
	y_result = append(y_result, y_prev)
	vx_result = append(vx_result, vx_prev)
	vy_result = append(vy_result, vy_prev)
	ax_result = append(ax_result, ax_prev)
	ay_result = append(ay_result, ay_prev)

	for i:=1;y_prev>0;i++{
		x := x_prev + h*vx_prev + math.Pow(h,2)/2*ax_prev
		y := y_prev + h*vy_prev + math.Pow(h,2)/2*ay_prev
		vx := vx_prev + h*ax_prev
		vy := vy_prev + h*ay_prev
		Fx_sources = 0.0
		Fy_sources = 0.0
		for _,gravity_source := range gravity_sources{
			r := math.Sqrt(math.Pow(gravity_source.x-x,2)+math.Pow(gravity_source.y-y,2))
			F := 6.67430*math.Pow(10,-11)*m*gravity_source.mass/math.Pow(r,2)
			Fx_sources += F*(gravity_source.x-x)/r 
			Fy_sources += F*(gravity_source.y-y)/r 
		}
		Fx = -b*(math.Pow(vx,2)+math.Pow(vy,2))*math.Cos(alfa) + Fx_sources
		Fy = b*(math.Pow(vx,2)+math.Pow(vy,2))*math.Sin(alfa) - m*g +Fy_sources
		if y<=(-math.Tan(alfa)*x+height+0.01) && y>=(-math.Tan(alfa)*x+height-0.01){
			diff := Fy*math.Cos(alfa)+Fx*math.Sin(alfa)
			if diff<0{
				Fx = Fx -diff*math.Sin(alfa)+diff*u*math.Cos(alfa) 
				Fy = Fy -diff*math.Cos(alfa) - diff*u*math.Sin(alfa)
			}
		}
		
		
		// fmt.Println(float64(i)*h)

		ax := Fx/m
		ay := Fy/m
		
		
		x_result = append(x_result, x)
		y_result = append(y_result, y)
		vx_result = append(vx_result, vx)
		vy_result = append(vy_result, vy)
		ax_result = append(ax_result, ax)
		ay_result = append(ay_result, ay)
		t_result = append(t_result, float64(i)*h)
		x_prev = x 
		y_prev = y 
		vx_prev = vx 
		vy_prev = vy 
		ax_prev = ax 
		ay_prev =ay
	}
	return x_result,y_result,t_result

	// x := float64(0) 
	// x_prev := float64(0) 
	// v:=float64(0) 
	// v_prev := float64(0) 
	// a:=float64(0) 
	// a_prev :=float64(0) 
	// var x_result []float64
	// var v_result []float64
	// var a_result []float64
	// var t_result []float64

	// for i := 0; x*math.Sin(alfa)< height; i++ {
	// 	x = x_prev + h*v_prev+math.Pow(h,2)/2*a_prev
	// 	v = v_prev + h*a_prev
	// 	a = -b/m*math.Pow(v_prev,2)-g*u*math.Cos(alfa)+g*math.Sin(alfa)
	// 	x_result = append(x_result,x)
	// 	v_result = append(v_result,v)
	// 	a_result = append(a_result,a)
	// 	t_result = append(t_result,float64(i)*h)
	// 	x_prev=x
	// 	v_prev = v
	// 	a_prev = a
	// }
	// return x_result,v_result,a_result,t_result
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
		lineData3[i].X = x_result[i]
		lineData3[i].Y = y_result[i]
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

	w, err := os.Create("plot.png")
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

