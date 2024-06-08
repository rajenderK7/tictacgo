package tictacgo

import "testing"

type (
	move struct {
		mark byte
		row  int
		col  int
	}
)

var (
	winMoves = []move{
		{byteX, 0, 0}, {byteO, 0, 1}, {byteX, 0, 2},
		{byteX, 1, 0}, {byteX, 1, 1}, {byteO, 1, 2},
		{byteX, 2, 0}, {byteO, 2, 1},
	}

	drawMoves = []move{
		{byteX, 0, 0}, {byteO, 0, 1}, {byteX, 0, 2},
		{byteX, 1, 0}, {byteX, 1, 1}, {byteO, 1, 2},
		{byteO, 2, 0}, {byteX, 2, 1}, {byteO, 2, 2},
	}
)

func TestPlace(t *testing.T) {
	b := newBoard(3)
	err := b.place(byteX, 0, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if b.grid[0] != byteX {
		t.Errorf("Expected %v at position 0, got %v instead", byteX, b.grid[0])

	}
	err = b.place(byteX, -1, 0)
	if err == nil {
		t.Errorf("Expected error \"%v\", got nil", INVALID_MOVE)
	}
	err = b.place(byteO, 0, 0)
	if err == nil {
		t.Errorf("Expected error: \"%v\", got nil", CELL_ALREADY_OCCUPIED)
	}
}

func TestHasWin(t *testing.T) {
	b := newBoard(3)
	for _, move := range winMoves {
		b.place(move.mark, move.row, move.col)
	}
	got := b.hasWin()
	want := true
	if got != want {
		t.Errorf("Expected %v, got %v", got, want)
	}
}

func TestIsDraw(t *testing.T) {
	b := newBoard(3)
	for _, move := range drawMoves {
		b.place(move.mark, move.row, move.col)
	}
	got := b.isDraw()
	want := true
	if got != want {
		t.Errorf("Expected %v, got %v", got, want)
	}
}

func TestCheckHorizontal(t *testing.T) {
	b := newBoard(4)
	b.place(byteO, 0, 0)
	b.place(byteO, 0, 1)
	b.place(byteO, 0, 2)
	b.place(byteO, 0, 3)
	got := b.checkHorizontal()
	want := true
	if got != want {
		t.Errorf("Expected %v, got %v", got, want)
	}
	b.reset()
	b.place(byteO, 0, 0)
	b.place(byteO, 0, 1)
	b.place(byteX, 0, 2)
	b.place(byteO, 0, 3)
	got = b.checkHorizontal()
	want = false
	if got != want {
		t.Errorf("Expected %v, got %v", got, want)
	}
}

func TestCheckVertical(t *testing.T) {
	b := newBoard(4)
	b.place(byteO, 0, 0)
	b.place(byteO, 1, 0)
	b.place(byteO, 2, 0)
	b.place(byteO, 3, 0)
	got := b.checkVertical()
	want := true
	if got != want {
		t.Errorf("Expected %v, got %v", got, want)
	}
	b.reset()
	b.place(byteO, 0, 0)
	b.place(byteO, 1, 0)
	b.place(byteX, 2, 0)
	b.place(byteO, 3, 0)
	got = b.checkVertical()
	want = false
	if got != want {
		t.Errorf("Expected %v, got %v", got, want)
	}
}

func TestCheckDiagonal(t *testing.T) {
	b := newBoard(3)
	b.place(byteO, 0, 0)
	b.place(byteO, 1, 1)
	b.place(byteO, 2, 2)
	got := b.checkDiagonal()
	want := true
	if got != want {
		t.Errorf("Expected %v, got %v", got, want)
	}
	b.reset()
	b.place(byteX, 0, 2)
	b.place(byteX, 1, 1)
	b.place(byteX, 2, 0)
	got = b.checkDiagonal()
	want = true
	if got != want {
		t.Errorf("Expected %v, got %v", got, want)
	}
	b.reset()
	b.place(byteX, 0, 2)
	b.place(byteO, 1, 1)
	b.place(byteX, 2, 0)
	got = b.checkDiagonal()
	want = false
	if got != want {
		t.Errorf("Expected %v, got %v", got, want)
	}
}
