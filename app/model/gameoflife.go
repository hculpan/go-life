package model

import (
	"math/rand"
	"time"

	"github.com/hculpan/go-sdl-lib/component"
)

type GameBoard [][]byte

type GameOfLife struct {
	component.BaseGame

	BoardWidth  int
	BoardHeight int

	activeBoard *GameBoard
	targetBoard *GameBoard

	cycle int
}

var Game *GameOfLife

func NewGameOfLife(gameWidth, gameHeight int32, livingRatio float32) *GameOfLife {
	rand.Seed(time.Now().UnixNano())

	result := &GameOfLife{
		BoardWidth:  int(gameWidth / 4),
		BoardHeight: int(gameHeight / 4),
	}

	result.activeBoard = result.initalizeBoard()
	result.targetBoard = result.initalizeBoard()

	for x := 0; x < result.BoardWidth; x++ {
		for y := 0; y < result.BoardHeight; y++ {
			if livingRatio > rand.Float32() {
				result.SetTargetBoardState(x, y, 1)
			}
		}
	}
	result.SwitchBoards()

	result.Initialize(gameWidth, gameHeight)

	Game = result

	return result
}

func (g GameOfLife) initalizeBoard() *GameBoard {
	var result GameBoard = make([][]byte, g.BoardWidth)
	for x := 0; x < int(g.BoardWidth); x++ {
		result[x] = make([]byte, g.BoardHeight)
		for y := 0; y < int(g.BoardHeight); y++ {
			result[x][y] = 0
		}
	}

	return &result
}

func (g *GameOfLife) SwitchBoards() {
	g.activeBoard, g.targetBoard = g.targetBoard, g.activeBoard
}

func (g GameOfLife) GetCurrentBoardState(x, y int) byte {
	if x < 0 || x >= g.BoardWidth || y < 0 || y >= g.BoardHeight {
		return 0
	}
	return (*g.activeBoard)[x][y]
}

func (g *GameOfLife) SetTargetBoardState(x, y int, newState byte) {
	(*g.targetBoard)[x][y] = newState
}

func (g GameOfLife) countNeighbors(x, y int) byte {
	var result byte = 0

	if g.GetCurrentBoardState(x-1, y-1) > 0 {
		result++
	}
	if g.GetCurrentBoardState(x, y-1) > 0 {
		result++
	}
	if g.GetCurrentBoardState(x+1, y-1) > 0 {
		result++
	}
	if g.GetCurrentBoardState(x-1, y) > 0 {
		result++
	}
	if g.GetCurrentBoardState(x+1, y) > 0 {
		result++
	}
	if g.GetCurrentBoardState(x-1, y+1) > 0 {
		result++
	}
	if g.GetCurrentBoardState(x, y+1) > 0 {
		result++
	}
	if g.GetCurrentBoardState(x+1, y+1) > 0 {
		result++
	}

	return result
}

func (g *GameOfLife) Update() error {
	g.cycle++

	for x := 0; x < int(g.BoardWidth); x++ {
		for y := 0; y < int(g.BoardHeight); y++ {
			n := g.countNeighbors(x, y)
			switch {
			case n < 2 || n > 3:
				g.SetTargetBoardState(x, y, 0)
			case n == 2:
				if g.GetCurrentBoardState(x, y) > 0 {
					g.SetTargetBoardState(x, y, 1)
				} else {
					g.SetTargetBoardState(x, y, 0)
				}
			case n == 3:
				g.SetTargetBoardState(x, y, 1)
			}
		}
	}

	g.SwitchBoards()

	return nil
}
