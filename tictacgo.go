package tictacgo

import (
	"fmt"
)

const (
	PlayerX = "X"
	PlayerO = "O"
)

var (
	ByteX byte = 'X'
	ByteO byte = 'O'
)

type Game struct {
	board  *board
	Winner string
	Player byte
}

type GameResult struct {
	IsDraw bool
	Winner string
}

func New(n int) *Game {
	return &Game{
		board:  newBoard(n),
		Winner: "",
		Player: ByteX,
	}
}

func (g *Game) Play(mark string, row, col int) (*GameResult, error) {
	markByte := mark[0]

	if markByte != g.Player {
		return nil, fmt.Errorf("expected player %s to make the move", string(g.Player))
	}

	err := g.board.place(markByte, row, col)
	if err != nil {
		return nil, err
	}
	if g.board.hasWin() {
		var winner string
		if g.Player == ByteX {
			winner = PlayerX
		} else {
			winner = PlayerO
		}
		return &GameResult{
			IsDraw: false,
			Winner: winner,
		}, nil
	} else if g.board.isDraw() {
		return &GameResult{
			IsDraw: true,
			Winner: "",
		}, nil
	}
	g.switchPlayer()
	return &GameResult{
		IsDraw: false,
		Winner: "",
	}, nil
}

func (g *Game) Reset() {
	g.board.reset()
}

func (g *Game) switchPlayer() {
	if g.Player == ByteX {
		g.Player = ByteO
	} else {
		g.Player = ByteX
	}
}
