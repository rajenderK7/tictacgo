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
		{ByteX, 0, 0}, {ByteO, 0, 1}, {ByteX, 0, 2},
		{ByteX, 1, 0}, {ByteX, 1, 1}, {ByteO, 1, 2},
		{ByteX, 2, 0}, {ByteO, 2, 1},
	}

	drawMoves = []move{
		{ByteX, 0, 0}, {ByteO, 0, 1}, {ByteX, 0, 2},
		{ByteX, 1, 0}, {ByteX, 1, 1}, {ByteO, 1, 2},
		{ByteO, 2, 0}, {ByteX, 2, 1}, {ByteO, 2, 2},
	}
)

func TestPlace(t *testing.T) {
	b := newBoard(3)
	err := b.place(ByteX, 0, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if b.grid[0] != ByteX {
		t.Errorf("Expected %v at position 0, got %v instead", ByteX, b.grid[0])

	}
	err = b.place(ByteX, -1, 0)
	if err == nil {
		t.Errorf("Expected error \"%v\", got nil", INVALID_MOVE)
	}
	err = b.place(ByteO, 0, 0)
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
	b.place(ByteO, 0, 0)
	b.place(ByteO, 0, 1)
	b.place(ByteO, 0, 2)
	b.place(ByteO, 0, 3)
	got := b.checkHorizontal()
	want := true
	if got != want {
		t.Errorf("Expected %v, got %v", got, want)
	}
	b.reset()
	b.place(ByteO, 0, 0)
	b.place(ByteO, 0, 1)
	b.place(ByteX, 0, 2)
	b.place(ByteO, 0, 3)
	got = b.checkHorizontal()
	want = false
	if got != want {
		t.Errorf("Expected %v, got %v", got, want)
	}
}

func TestCheckVertical(t *testing.T) {
	b := newBoard(4)
	b.place(ByteO, 0, 0)
	b.place(ByteO, 1, 0)
	b.place(ByteO, 2, 0)
	b.place(ByteO, 3, 0)
	got := b.checkVertical()
	want := true
	if got != want {
		t.Errorf("Expected %v, got %v", got, want)
	}
	b.reset()
	b.place(ByteO, 0, 0)
	b.place(ByteO, 1, 0)
	b.place(ByteX, 2, 0)
	b.place(ByteO, 3, 0)
	got = b.checkVertical()
	want = false
	if got != want {
		t.Errorf("Expected %v, got %v", got, want)
	}
}

func TestCheckDiagonal(t *testing.T) {
	b := newBoard(3)
	b.place(ByteO, 0, 0)
	b.place(ByteO, 1, 1)
	b.place(ByteO, 2, 2)
	got := b.checkDiagonal()
	want := true
	if got != want {
		t.Errorf("Expected %v, got %v", got, want)
	}
	b.reset()
	b.place(ByteX, 0, 2)
	b.place(ByteX, 1, 1)
	b.place(ByteX, 2, 0)
	got = b.checkDiagonal()
	want = true
	if got != want {
		t.Errorf("Expected %v, got %v", got, want)
	}
	b.reset()
	b.place(ByteX, 0, 2)
	b.place(ByteO, 1, 1)
	b.place(ByteX, 2, 0)
	got = b.checkDiagonal()
	want = false
	if got != want {
		t.Errorf("Expected %v, got %v", got, want)
	}
}
