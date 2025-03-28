package imageprocessing

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

// readImage function reads an image from the file system and returns the image.
func ReadImage(path string) image.Image {
	inputFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	img, _, err := image.Decode(inputFile)
	if err != nil {
		fmt.Println(path)
		panic(err)
	}
	return img
}

func WriteImage(path string, img image.Image) {
	outputFile, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	err = jpeg.Encode(outputFile, img, nil)
	if err != nil {
		panic(err)
	}
}

// writeImage function converts an image to grayscale and returns the grayscale image.
func GrayScale(img image.Image) image.Image {
	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalPixel := img.At(x, y)
			grayPixel := color.GrayModel.Convert(originalPixel)
			grayImg.Set(x, y, grayPixel)
		}
	}
	return grayImg
}

func ResizeImage(img image.Image) image.Image {
	newWidth := uint(500)
	newHeight := uint(500)
	resizingImage := resize.Resize(newWidth, newHeight, img, resize.Lanczos2)
	return resizingImage
}
