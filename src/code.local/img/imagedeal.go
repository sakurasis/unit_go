package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	// open "test.jpg"
	file, err := os.Open("C:\\Users\\sakura\\Desktop\\unit_go\\src\\code.local\\img\\test.png")
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio

	width := img.Bounds().Dx()  //
	height := img.Bounds().Dy() //

	m := resize.Resize(uint(width)*3, uint(height)*3, img, resize.Lanczos3)
	//m := resize.Resize(1000, 0, img, resize.Lanczos3)

	out, err := os.Create("C:\\Users\\sakura\\Desktop\\unit_go\\src\\code.local\\img\\test1.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
	fmt.Println("------------------")
}
