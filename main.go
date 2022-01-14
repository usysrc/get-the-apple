package main

import (
	"cart/w4"
)

var smileySprite = [8]byte{
	0b11000011,
	0b10000001,
	0b00100100,
	0b00100100,
	0b00000000,
	0b00100100,
	0b10011001,
	0b11000011,
}
var appleSprite = [8]byte{
	0b11000010,
	0b10000001,
	0b00000000,
	0b00000000,
	0b00000000,
	0b00000000,
	0b10000001,
	0b11000011,
}

func drawApple(x, y int) {
	w4.Blit(&appleSprite[0], x, y, 8, 8, w4.BLIT_1BPP)
}

func drawPlayer(x, y int) {
	w4.Blit(&smileySprite[0], x, y, 8, 8, w4.BLIT_1BPP)
}

var player = struct {
	x int
	y int
}{
	x: 1,
	y: 1,
}

var apple = struct {
	x int
	y int
}{
	x: 0,
	y: 0,
}

func randomApple() {
	apple.x = (apple.x + 123456789) % 20
	apple.y = (apple.x + 123456789) % 20
}

var previousGamepad byte

type Game struct {
	CurrentState func()
}

var game Game

func (g *Game) update() {
	g.CurrentState()
}

func (g *Game) won() {
	w4.Text("You got the apple!", 10, 10)
	w4.Text("Button 1 to\n continue.", 40, 120)

	var gamepad = *w4.GAMEPAD1
	var pressedThisFrame = gamepad & (gamepad ^ previousGamepad)
	previousGamepad = gamepad

	if pressedThisFrame&w4.BUTTON_1 != 0 {
		randomApple()
		g.CurrentState = g.gameplay
	}
}

func (g *Game) checkPlayer() {
	if player.x == apple.x && player.y == apple.y {
		g.CurrentState = g.won
	}
}

func (g *Game) gameplay() {

	*w4.DRAW_COLORS = 2

	var gamepad = *w4.GAMEPAD1
	var pressedThisFrame = gamepad & (gamepad ^ previousGamepad)
	previousGamepad = gamepad

	if pressedThisFrame&w4.BUTTON_RIGHT != 0 {
		player.x++
		g.checkPlayer()
	}
	if pressedThisFrame&w4.BUTTON_LEFT != 0 {
		player.x--
		g.checkPlayer()

	}
	if pressedThisFrame&w4.BUTTON_UP != 0 {
		player.y--
		g.checkPlayer()

	}
	if pressedThisFrame&w4.BUTTON_DOWN != 0 {
		player.y++
		g.checkPlayer()
	}
	drawApple(apple.x*8, apple.y*8)
	drawPlayer(player.x*8, player.y*8)
}

//go:export start
func start() {
	game = Game{}
	randomApple()
	game.CurrentState = game.gameplay
}

//go:export update
func update() {
	game.update()
}
