package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/png"
	"log"
	"os"
	"pyramid/pyramid"
	"runtime"
	"strconv"
	"time"
)

var nFlag = flag.Int("n", 2, "number of pyramid levels")
var gFlag = flag.Int("g", runtime.GOMAXPROCS(0), "number of processing goroutines")
var filenameFlag = flag.String("file", "", "image path")

func main() {
	flag.Parse()
	if *filenameFlag == "" {
		log.Fatalln("Image path is required")
	}

	f, err := os.Open(*filenameFlag)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	intensity, err := GetIntensityByImage(img)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(*nFlag)
	start := time.Now()
	pyramid.BuildPyramid(CreateImageByMatrix, intensity, *nFlag, *gFlag)
	fmt.Println(time.Since(start))
}

func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

func GetIntensityByImage(img image.Image) ([][]uint8, error) {
	size := img.Bounds().Size()
	if !isPowerOfTwo(size.X) || !isPowerOfTwo(size.Y) {
		return nil, fmt.Errorf("the image resolution must be a power of two")
	}
	intensity := make([][]uint8, size.X)
	for i := 0; i < size.X; i++ {
		intensity[i] = make([]uint8, size.Y)
	}

	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			pixel := img.At(x, y)
			originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
			r := float64(originalColor.R)
			g := float64(originalColor.G)
			b := float64(originalColor.B)
			grey := uint8((r + g + b) / 3)
			intensity[x][y] = grey
		}
	}
	return intensity, nil
}

func CreateImageByMatrix(intens [][]uint8, level int) {
	SizeX := len(intens)
	SizeY := len(intens[0])
	rect := image.Rect(0, 0, SizeX, SizeY)
	nextImg := image.NewRGBA(rect)
	fillImage(intens, nextImg)
	CreateImageFile(level, nextImg)
}

func fillImage(intens [][]uint8, img *image.RGBA) {
	for x := 0; x < len(intens); x++ {
		for y := 0; y < len(intens[0]); y++ {
			grey := intens[x][y]
			c := color.Gray{grey}
			img.Set(x, y, c)
		}
	}
}

func CreateImageFile(lvl int, img *image.RGBA) {
	outFile, _ := os.Create("lvl" + strconv.Itoa(lvl) + ".jpeg")
	if err := jpeg.Encode(outFile, img, nil); err != nil {
		outFile.Close()
		log.Fatal(err)
	}
	outFile.Close()
}
