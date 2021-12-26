package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"image"
	"image/color"
)

func main() {
	go func() {
		w := app.NewWindow()

		for e := range w.Events() {
			switch e := e.(type) {
			case system.DestroyEvent:
				fmt.Print(e.Err)
			case system.FrameEvent:
				ops := new(op.Ops)
				draw(ops)
				e.Frame(ops)
			}
		}
	}()
	app.Main()
}

func draw(ops *op.Ops) {
	clip.Rect{Max: image.Pt(100, 100)}.Add(ops)
	paint.ColorOp{Color: color.NRGBA{R: 0x80, A: 0xFF}}.Add(ops)
	paint.PaintOp{}.Add(ops)
}
