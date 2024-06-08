package tictacgo

import "errors"

// board represents the grid of the Game.
type board struct {
	grid []byte
	n    int
	size int
}

// newBoard returns a n-by-n game board.
func newBoard(n int) *board {
	// TODO: Set an upper-bound for 'n' so that
	// game doesn't get too complicated.
	size := n * n
	return &board{
		grid: make([]byte, size),
		n:    n,
		size: size,
	}
}

func (b *board) reset() {
	b.grid = make([]byte, b.size)
}

// place will place the mark on the board in the postition
// calculated using the 'row' and 'col' co-ordinates.
func (b *board) place(mark byte, row, col int) error {
	// Position in the 1D representation.
	pos := row*b.n + col
	if pos < 0 || pos >= b.size {
		return errors.New(INVALID_MOVE)
	} else if b.grid[pos] == byteX && b.grid[pos] == byteO {
		return errors.New(CELL_ALREADY_OCCUPIED)
	}
	b.grid[pos] = mark
	return nil
}

// hasWin returns true if there is a winning condition
// in the board false otherwise.
func (b *board) hasWin() bool {
	return b.checkHorizontal() || b.checkVertical() || b.checkDiagonal()
}

// isDraw returns true if the game ends in a draw false otherwise.
func (b *board) isDraw() bool {
	emptyCells := b.size
	for _, mark := range b.grid {
		if mark == byteX || mark == byteO {
			emptyCells--
		}
	}
	return emptyCells == 0
}

// checkHorizontal is a helper function to check for the winning condition.
func (b *board) checkHorizontal() bool {
	i := 0
	lvl := 1
	for i < b.size {
		j := 0
		count := 0
		for i+j < lvl*b.n {
			mark := b.grid[i+j]
			if mark == byteX {
				count++
			} else if mark == byteO {
				count--
			}
			j++
		}
		if count == -b.n || count == b.n {
			return true
		}
		lvl++
		i += j
	}
	return false
}

// checkVertical is a helper function to check for the winning condition.
func (b *board) checkVertical() bool {
	for i := 0; i < b.n; i++ {
		count := 0
		for j := 0; j < b.n; j++ {
			mark := b.grid[i+j*b.n]
			if mark == byteX {
				count++
			} else if mark == byteO {
				count--
			}
		}
		if count == -b.n || count == b.n {
			return true
		}
	}
	return false
}

// checkDiagonal is a helper function to check for the winning condition.
func (b *board) checkDiagonal() bool {
	countMainDiag, countAntiDiag := 0, 0
	for i := range b.n {
		markMainDiag := b.grid[i*b.n+i]
		markAntiDiag := b.grid[i*b.n+(b.n-1-i)]
		if markMainDiag == byteX {
			countMainDiag++
		} else if markMainDiag == byteO {
			countMainDiag--
		}
		if markAntiDiag == byteX {
			countAntiDiag++
		} else if markAntiDiag == byteO {
			countAntiDiag--
		}
		if countMainDiag == -b.n || countMainDiag == b.n {
			return true
		}
		if countAntiDiag == -b.n || countAntiDiag == b.n {
			return true
		}
	}
	return false
}
