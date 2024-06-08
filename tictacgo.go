package tictacgo

const (
	playerX = "X"
	playerO = "O"
)

var (
	byteX byte = 'X'
	byteO byte = 'O'
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
		Player: byteX,
	}
}

func (g *Game) Play(mark string, row, col int) (*GameResult, error) {
	markByte := mark[0]
	err := g.board.place(markByte, row, col)
	if err != nil {
		return nil, err
	}
	if g.board.hasWin() {
		var winner string
		if g.Player == byteX {
			winner = playerX
		} else {
			winner = playerO
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

func (g *Game) switchPlayer() {
	if g.Player == byteX {
		g.Player = byteO
	} else {
		g.Player = byteX
	}
}
