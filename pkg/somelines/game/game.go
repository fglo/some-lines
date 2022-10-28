package game

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/fglo/some-lines/pkg/somelines/board"
	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	Terminated = errors.New("terminated")
)

// Game encapsulates game logic
type Game struct {
	// input      *Input
	board  *board.Board
	pixels []byte

	screenWidth  int
	screenHeight int

	quitIsPressed    bool
	forwardIsPressed bool
	reverseIsPressed bool
	debugIsToggled   bool

	paused bool

	focalLength int

	counter int
}

// New generates a new Game object.
func New() *Game {
	g := new(Game)
	g.screenWidth = screenWidth
	g.screenHeight = screenHeight

	g.board = board.New(g.screenWidth, g.screenHeight)

	g.focalLength = 80

	g.counter = 0

	ebiten.SetWindowSize(g.getWindowSize())
	ebiten.SetWindowTitle("Particles' Rules of Attraction")

	return g
}

func (g *Game) getWindowSize() (int, int) {
	var factor float32 = 3
	return int(float32(g.screenWidth) * factor), int(float32(g.screenHeight) * factor)
}

// Layout implements ebiten.Game's Layout.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.screenWidth, g.screenHeight
}

// Update updates the current game state.
func (g *Game) Update() error {
	g.checkShadeButton()
	g.checkPauseButton()
	g.checkForwardButton()
	g.checkReverseButton()
	g.checkDebugButton()
	g.checkPlusButton()
	g.checkMinusButton()
	if err := g.board.Update(); err != nil {
		return err
	}
	return g.checkQuitButton()
}

func (g *Game) checkQuitButton() error {
	if !g.quitIsPressed && inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		g.quitIsPressed = true
	}
	if g.quitIsPressed && inpututil.IsKeyJustReleased(ebiten.KeyQ) {
		g.quitIsPressed = false
		return Terminated
	}
	return nil
}

func (g *Game) checkShadeButton() {
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.board.ToggleShading()
	}
}

func (g *Game) checkPauseButton() {
	if inpututil.IsKeyJustPressed(ebiten.KeyP) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.paused = !g.paused
	}
}

func (g *Game) checkForwardButton() {
	if !g.forwardIsPressed && (inpututil.IsKeyJustPressed(ebiten.KeyF) || inpututil.IsKeyJustPressed(ebiten.KeyArrowRight)) {
		g.forwardIsPressed = true
	}
	if g.forwardIsPressed && (inpututil.IsKeyJustReleased(ebiten.KeyF) || inpututil.IsKeyJustReleased(ebiten.KeyArrowRight)) {
		g.forwardIsPressed = false
	}
}

func (g *Game) checkReverseButton() {
	if !g.reverseIsPressed && (inpututil.IsKeyJustPressed(ebiten.KeyB) || inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft)) {
		g.reverseIsPressed = true
	}
	if g.reverseIsPressed && (inpututil.IsKeyJustReleased(ebiten.KeyB) || inpututil.IsKeyJustReleased(ebiten.KeyArrowLeft)) {
		g.reverseIsPressed = false
	}
}

func (g *Game) checkDebugButton() {
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		g.debugIsToggled = !g.debugIsToggled
	}
}

func (g *Game) checkPlusButton() {
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		g.focalLength += 5
		if g.focalLength > 200 {
			g.focalLength = 200
		}
		fmt.Println(g.focalLength)
	}
}

func (g *Game) checkMinusButton() {
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		g.focalLength -= 5
		if g.focalLength < 40 {
			g.focalLength = 40
		}
		fmt.Println(g.focalLength)
	}
}

// Draw draws the current game to the given screen.
func (g *Game) Draw(screen *ebiten.Image) {
	if g.pixels == nil {
		g.pixels = make([]byte, screenWidth*screenHeight*4)
	}
	if !g.paused || g.forwardIsPressed {
		g.counter++
	} else if g.reverseIsPressed {
		g.counter--
	}
	g.clearPixels()
	// g.board.Draw(g.pixels, g.counter, g.focalLength)
	g.board.DrawScene(g.pixels, g.counter, g.focalLength)
	// g.board.DrawTeapot(g.pixels, g.counter, g.focalLength)
	screen.WritePixels(g.pixels)
}

func (g *Game) clearPixels() {
	for i := range g.pixels {
		g.pixels[i] = 0
	}
}
