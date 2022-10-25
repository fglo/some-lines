package main

import (
	"log"

	"github.com/fglo/some-lines/pkg/somelines/game"
	ebiten "github.com/hajimehoshi/ebiten/v2"
)

func run() {
	g := game.New()

	if err := ebiten.RunGame(g); err != nil {
		if err == game.Terminated {
			return
		}
		log.Fatal(err)
	}
}

func main() {
	run()
}
