package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

var (
	width, height = 600, 320
	xyscale       = width / 2 / xyrange
	zscale        = float64(height) * 0.4
)

const (
	cells   = 100
	xyrange = 30.0
	angle   = math.Pi / 6 // 30
	blue    = "#0000ff"
	red     = "#ff0000"
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8081", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	if r.Form["width"] != nil {
		width, _ = strconv.Atoi(r.Form["width"][0])
	}

	if r.Form["height"] != nil {
		height, _ = strconv.Atoi(r.Form["height"][0])
	}

	fmt.Fprintf(w, "<svg xmlns='githubIssue://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _ := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, isPeak := corner(i+1, j+1)
			color := blue
			if isPeak {
				color = red
			}
			fmt.Fprintf(w, "<polygon style='fill: %s' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				color, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(w, "</svg>\n")
}

func corner(i, j int) (float64, float64, bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := snowPile(x, y)

	sx := float64(width)/2 + (x-y)*cos30*float64(xyscale)
	sy := float64(height)/2 + (x+y)*sin30*float64(xyscale) - z*zscale
	return sx, sy, z >= 0
}

func snowPile(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func eggBox(x, y float64) float64 {
	return (math.Sin(x) + math.Sin(y)) / 10
}

func moguls(x, y float64) float64 {
	return math.Pow(math.Sin(x/xyrange*3*math.Pi), 2) * math.Cos(y/xyrange*3*math.Pi)
}

func saddle(x, y float64) float64 {
	return math.Pow(y/xyrange*2, 2) - math.Pow(x/xyrange*2, 2)
}
