package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/jpeg"
	"log"
	"net/http"
	"strconv"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
)

func main() {

	router := httprouter.New()

	n := negroni.Classic()
	n.UseHandler(router)

	router.GET("/", Color)

	n.Run(":3000")
}

func Color(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	x := 600
	y := 600

	tangle := image.Rect(0, 0, x, y)
	m := image.NewRGBA(tangle)

	m.Set(5, 5, color.RGBA{255, 0, 0, 255})

	palt := palette.WebSafe
	total := len(palt)

	fmt.Println(total)

	var img image.Image = m
	pix(&img)
	writeImage(w, &img)
}

func writeImage(w http.ResponseWriter, img *image.Image) {

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}

func pix(img *image.Image) {
	buf := new(bytes.Buffer)
	err := jpeg.Encode(buf, *img, nil)
	fmt.Println(err)
	arr := buf.Bytes()
	fmt.Println(arr)
}
