package main

import (
  "fmt"
  //"time"

  "path/filepath"
	"image/color"

	"github.com/oakmound/oak"
	"github.com/oakmound/oak/render"
	//"github.com/oakmound/oak/render/mod"
	"github.com/oakmound/oak/dlog"
	"github.com/oakmound/oak/scene"
	//"github.com/oakmound/oak/mouse"
)

var (
  red = color.RGBA{255, 0, 0, 255}
  s   = render.NewColorBox(650, 450, color.RGBA{255, 255, 255, 255})
  r   = s.GetRGBA()
  x   = float32(0)
)

func  main() {
  oak.Add("game",
    func(prevScene string, inData interface{}) {
      fmt.Printf("hello\n")
      render.Draw(s)

      spr, err := render.LoadSprite(filepath.Join("images"), "bow.png")
      if err != nil {
        dlog.Error(err)
      }
      //render.Draw(spr.Modify(mod.Rotate(x)))
      render.Draw(spr)
      //time.Sleep(100 * time.Millisecond)
    },
    func() bool {
      return true
    },
    func() (nextScene string, result *scene.Result) {
      return "game", nil
    })
  oak.Init("game")
}
