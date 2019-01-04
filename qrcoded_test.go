package main

import (
	"bytes"
	"image/png"
	"testing"
)

func TestGenerateQrCode(t *testing.T) {
	// result := GenerateQRCode("555-2368")

	// if result == nil {
	// 	t.Errorf("nil")
	// }
	// if len(result) == 0 {
	// 	t.Errorf("has no data")
	// }
	result := GenerateQRCode("555-2368")
	buf := bytes.NewBuffer(result)
	_, err := png.Decode(buf)
	if err != nil {
		t.Errorf("not a png: %s", err)
	}

}
