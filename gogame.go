package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/alexddrexler/gogame/game"
)

func main() {
	fmt.Println("It's GoGAME TIME!")
	game := &game.Game{}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("GoGame")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
	fmt.Println("GoGAME TIME is over!")
}