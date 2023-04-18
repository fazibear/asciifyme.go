package webcam

import (
	"fmt"
	"syscall/js"
)

var (
	navigator js.Value
	video     js.Value
)

func init() {
	navigator = js.Global().Get("navigator")
	video = js.Global().Get("document").Call("createElement", "video")
}

func Setup() js.Value {
	user_media_params := map[string]interface{}{
		"video": true,
	}

	navigator.Call("getUserMedia", user_media_params, js.FuncOf(stream), js.FuncOf(err))

	return video
}

func err(this js.Value, args []js.Value) interface{} {
	fmt.Println("err")
	return nil
}

func stream(this js.Value, args []js.Value) interface{} {
	video.Set("srcObject", args[0])
	video.Call("addEventListener", "canplaythrough", js.FuncOf(canPlay))
	return nil
}

func canPlay(this js.Value, args []js.Value) interface{} {
	video.Call("play")
	return nil
}
