package asciifyier

import (
	"asciifyme/canvas"
)

const (
	Chars       = "   .,:;i1tfLCG08@"
	CharsLength = 16
)

func Asciify(data []uint8) string {
	output := ""

	for y := 0; y < canvas.CanvasHeight; y++ {
		for x := 0; x < canvas.CanvasWidth; x++ {
			offset := (y*canvas.CanvasWidth + x) * 4

			red := data[offset]
			green := data[offset+1]
			blue := data[offset+2]
			//alpha := data[offset+3]

			brightness := (0.3*float64(red) + 0.59*float64(green) + 0.11*float64(blue)) / 255.0

			char_index := CharsLength - int(brightness*CharsLength)

			output += string(Chars[char_index])
		}
		output += "\n"
	}

	return output
}
