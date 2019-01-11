package main

import (
	"bytes"
	"image"
	"image/png"
)

// func main() {
// 	// fmt.Println(GenerateQRCode("555-555"))
// 	// var val byte = 0xFa
// 	// fmt.Println(val)
// 	// fmt.Println(GenerateQRCode("555-2368"))
// 	// result := GenerateQRCode("555-2368")
// 	// buffer := bytes.NewBuffer(result)
// 	// _, err := png.Decode(buffer)

// 	// fmt.Println(len(GenerateQRCode("")))

// 	fmt.Printf("hello")

// }
func GenerateQRCode(code string) []byte {
	// return []byte{0xFF}
	img := image.NewRGBA(image.Rect(0, 0, 8, 8))
	buffer := new(bytes.Buffer)
	_ = png.Encode(buffer, img)
	return buffer.Bytes()
}
