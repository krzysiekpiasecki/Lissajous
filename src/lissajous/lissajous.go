package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

var pallete = []color.Color{color.White, color.Black}

const (
	whiteIndex    = 0
	blackIndex    = 1
	cyclesDefault = 5
	sizeDefualt   = 100
	framesDefault = 64
	delayDefault  = 8
	resDefault    = 0.001
)

type LissajousConfig struct {
	Cycles int
	Res    float64
	Size   int
	Frames int
	Delay  int
}

func main() {
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			printLissajous(w, parseConfig(r.URL.Query()))
		})
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func printLissajous(out io.Writer, cfg LissajousConfig) {
	cycles := cfg.Cycles
	res := cfg.Res
	size := cfg.Size
	nframes := cfg.Frames
	delay := cfg.Delay

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallete)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func parseConfig(q url.Values) LissajousConfig {
	cycles, err := strconv.Atoi(q.Get("cycles"))
	if err != nil {
		cycles = cyclesDefault
	}
	size, err := strconv.Atoi(q.Get("size"))
	if err != nil {
		size = sizeDefualt
	}
	frames, err := strconv.Atoi(q.Get("frames"))
	if err != nil {
		frames = framesDefault
	}
	delay, err := strconv.Atoi(q.Get("delay"))
	if err != nil {
		delay = delayDefault
	}
	res, err := strconv.ParseFloat(q.Get("res"), 64)
	if err != nil {
		res = resDefault
	}

	return LissajousConfig{Cycles: cycles, Size: size, Frames: frames, Delay: delay, Res: res}
}
