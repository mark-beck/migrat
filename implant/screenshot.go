package main

import (
	"bytes"
	"image"
	"image/png"
	"log"

	"github.com/vova616/screenshot"
)

func take_screenshot() ([]byte, error) {
	img, err := screenshot.CaptureScreen()
	if err != nil {
		log.Println("error taking screenshot: ", err.Error())
		return nil, err
	}
	myImg := image.Image(img)

	var image_bytes bytes.Buffer

	err = png.Encode(&image_bytes, myImg)
	if err != nil {
		log.Println("error encoding: ", err.Error())
		return nil, err
	}
	return image_bytes.Bytes(), nil
}
