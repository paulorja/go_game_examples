package main

import (
	"image/color"
	//"time"
  "fmt"

	//"github.com/200sc/go-dist/floatrange"

	"github.com/oakmound/oak"
	//"github.com/oakmound/oak/collision"
	//"github.com/oakmound/oak/entities"
	//"github.com/oakmound/oak/event"
	//"github.com/oakmound/oak/key"
	"github.com/oakmound/oak/render"
	"github.com/oakmound/oak/scene"
	//"github.com/oakmound/oak/timing"
)

var (
  red = color.RGBA{255, 0, 0, 255}
  s   = render.NewColorBox(250, 250, color.RGBA{255, 255, 255, 255})
  r   = s.GetRGBA()
  x   = 0.0
)


func main() {
  oak.Add("game",
      func(prevScene string, inData interface{}) {
        fmt.Printf("hello \n")
        render.Draw(s)
        //render.DrawThickLine(r, 10, 20, 30, 40, red, 1)
        //render.DrawCircle(r, red, 10, 1, 0, 64)
      },
      func() bool {
        x = x + 1
        render.DrawCurve(r, red, 50, 1, 0, x * 0.01, 24, 24)
        return true
      },
      func() (nextScene string, result *scene.Result) {
        return "game", nil
      })
  oak.Init("game")
}
