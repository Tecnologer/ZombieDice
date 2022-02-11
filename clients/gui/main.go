// Copyright 2014 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build example
// +build example

package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

const mosaicRatio = 16

type Game struct {
	gophersRenderTarget *ebiten.Image
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Shrink the image once.
	img, _, err := ebitenutil.NewImageFromFile("./images/brain_easy.png")
	if err != nil {
		panic(err)
	}
	screen.DrawImage(img, nil)

	img2, _, err := ebitenutil.NewImageFromFile("./images/brain_easy.png")
	if err != nil {
		panic(err)
	}
	op := &ebiten.DrawImageOptions{}

	screen.DrawImage(img2, op)

	img3, _, err := ebitenutil.NewImageFromFile("./images/brain_easy.png")
	if err != nil {
		panic(err)
	}
	screen.DrawImage(img3, nil)

	// Enlarge the shrunk image.
	// The filter is the nearest filter, so the result will be mosaic.
	// op = &ebiten.DrawImageOptions{}
	// op.GeoM.Scale(mosaicRatio, mosaicRatio)
	// screen.DrawImage(g.gophersRenderTarget, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	g := &Game{}
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Zombie Dice")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
