// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"math"
	"net/http"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	http.HandleFunc("/display", sin)
	err := http.ListenAndServe(":9100", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func sin(w http.ResponseWriter, q *http.Request) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 300   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := nframes; i > 0; i-- {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for loop := int(phase*100); loop > 0; loop -= 1 {
			r := float64(50)
			offset_x := r*(2*math.Cos(float64(loop))-math.Cos(2*float64(loop))) + 200;
			offset_y := r*(2*math.Sin(float64(loop))-math.Sin(2*float64(loop))) + 200;
			img.SetColorIndex(int(offset_x)+size/2, int(offset_y)+size/2, blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(w, &anim) //浏览器显示
}
