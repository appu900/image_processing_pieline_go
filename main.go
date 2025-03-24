package main

import (
	"fmt"
	"image"
	imageProcessing "image-processing-pipeline/image-processing"
	"strings"
)

type JOB struct {
	InputPath  string
	Image      image.Image
	OutputPath string
}

func LoadImage(paths []string) <-chan JOB {
	out := make(chan JOB)
	go func() {
		for _, path := range paths {
			job := JOB{InputPath: path, OutputPath: strings.Replace(path, "images/", "images/output/", 1)}
			job.Image = imageProcessing.ReadImage(path)
			out <- job
		}
		close(out)
	}()
	return out
}

func resize(input <-chan JOB) <-chan JOB {
	out := make(chan JOB)
	go func() {
		for job := range input {
			job.Image = imageProcessing.ResizeImage(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func convertToGrayScale(input <-chan JOB) <-chan JOB {
	out := make(chan JOB)
	go func() {
		for job := range input {
			job.Image = imageProcessing.GrayScale(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func saveImage(input <-chan JOB) <-chan bool {
	out := make(chan bool)
	go func() {
		for job := range input {
			imageProcessing.WriteImage(job.OutputPath, job.Image)
			out <- true
		}
		close(out)
	}()
	return out
}

func main() {
	imagePaths := []string{"images/image1.jpg",
		"images/image2.jpg",
	}

	channaelone := LoadImage(imagePaths)
	channelTwo := resize(channaelone)
	channelThree := convertToGrayScale(channelTwo)
	writeresults := saveImage(channelThree)
	for success := range writeresults {
		if success {
			fmt.Println("Image processed successfully")
		} else {
			fmt.Println("Error processing image")
		}
	}
}
