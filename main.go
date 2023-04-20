package main

import (
	"asciifyme/asciifyier"
	"asciifyme/canvas"
	"asciifyme/webcam"
	"syscall/js"
)

var (
	camera js.Value
	window js.Value
	pre    js.Value
)

func init() {
	camera = webcam.Setup()
	window = js.Global().Get("window")
	pre = js.Global().Get("document").Call("getElementById", "pre")
}

func loop(this js.Value, args []js.Value) interface{} {
	window.Call("requestAnimationFrame", js.FuncOf(loop))
	canvas.DrawImage(camera)
	imageData := canvas.GetImageData()
	output := asciifyier.Asciify(imageData)
	pre.Set("innerHTML", output)
	return nil
}

func main() {
	loop(js.ValueOf(nil), make([]js.Value, 0))

	select {}
}
