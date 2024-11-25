package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"log"
	//"math/rand/v2"
)

const (
	screenWidth           = 1000
	screenHeight          = 1000
	ballwidth, ballheight = screenWidth / 75, screenHeight / 75
)

var (
	ballx, bally       float64 = screenWidth / 2, screenHeight / 2
	ballvelocityX      float64 = 2
	ballvelocityY      float64 = 2
	paddle1Y, paddle2Y float64 = screenHeight / 2, screenHeight / 2
	paddlewidth        float64 = screenWidth / 120
	paddleheight       float64 = screenHeight / 10
)

type Game struct{}

func (g *Game) Update() error {
	//Checks if the ball is touching the walls
	if ballx >= screenWidth || ballx <= 0 {
		ballvelocityX *= -1
	}
	if bally >= screenHeight || bally <= 0 {
		ballvelocityY *= -1
	}
	//checks if the ball is touching a paddle
	if (ballx <= paddlewidth && bally >= paddle1Y && bally <= paddle1Y+paddleheight) || (ballx >= screenWidth-paddlewidth-ballwidth && bally >= paddle2Y && bally <= paddle2Y+paddleheight) {
		ballvelocityX = -ballvelocityX
	}
	//Paddle movement
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		paddle1Y -= screenHeight / 150
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		paddle1Y += screenHeight / 150
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		paddle2Y -= screenHeight / 150
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		paddle2Y += screenHeight / 150
	}
	ballx += ballvelocityX
	bally += ballvelocityY
	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{255, 255, 255, 255})
	ebitenutil.DrawRect(screen, screenWidth/120, paddle1Y, paddlewidth, paddleheight, color.Black)
	ebitenutil.DrawRect(screen, screenWidth-paddlewidth-(screenWidth/120), paddle2Y, paddlewidth, paddleheight, color.Black)

	ebitenutil.DrawRect(screen, ballx, bally, ballwidth, ballheight, color.Black)
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("The First")

	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
