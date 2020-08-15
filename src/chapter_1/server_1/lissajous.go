// Генерируем анимированный GIF из случайных фигур Лиссажу

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"time"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{54, 207, 31, 1},
	color.RGBA{207, 31, 69, 1},
	color.RGBA{5, 27, 230, 1},
	color.RGBA{207, 5, 230, 1},
	color.RGBA{226, 230, 5, 1},
	color.RGBA{5, 226, 230, 1},
}

const (
	whiteIndex = 0 //First color of palette
	blackIndex = 1 //Next color of palette
)

func lissajous(out io.Writer, params map[string]int) {
	const (
		res = 0.001 // Угловое разрешение
	)
	cycles := params["cycles"]   // Количество полных колебаний x
	size := params["size"]       // Канва изображения охватывает [size..+size]
	nframes := params["nframes"] //Количество кадров анимации
	delay := params["delay"]     //Задержка между кажрами (единица - 10мс)
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0 // Относительная частота колебаний y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Разность фаз
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		colorIndex := rand.Intn(len(palette))
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), uint8(colorIndex))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // Игнорируем ошибки
}
