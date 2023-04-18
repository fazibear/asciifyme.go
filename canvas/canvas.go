package canvas

import (
	"syscall/js"
)

const (
	CanvasWidth  = 80
	CanvasHeight = 40
)

var (
	ctx js.Value
)

func init() {
	ctx = js.Global().Get("document").Call("createElement", "canvas").Call("getContext", "2d")
}

func DrawImage(video js.Value) {
	ctx.Call("drawImage", video, 0, 0, CanvasWidth, CanvasHeight)
}

func GetImageData() []uint8 {
	data := ctx.Call("getImageData", 0, 0, CanvasWidth, CanvasHeight).Get("data")

	lenght := data.Get("length").Int()

	goData := make([]uint8, lenght)

	js.CopyBytesToGo(goData, data)

	return goData
}
