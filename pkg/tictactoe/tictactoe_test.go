package tictactoe

import (
	"testing"
)

func TestBoardEval(t *testing.T) {
	b1 := MakeBoard("X-O-X-O-X")
	if b1.EvalBoard() != XWon {
		fail(t, b1)
	}

	b2 := MakeBoard("XXX---O-O")
	if b2.EvalBoard() != XWon {
		fail(t, b2)
	}

	b3 := MakeBoard("X-X---O-O")
	if b3.EvalBoard() != NotEnd {
		fail(t, b3)
	}

	b4 := MakeBoard("XOXOXXOXO")
	if b4.EvalBoard() != Tie {
		fail(t, b4)
	}

	b5 := MakeBoard("XOXXO--OO")
	if b5.EvalBoard() != OWon {
		fail(t, b5)
	}
}

func fail(t *testing.T, b Board) {
	t.Errorf("Board %v does not have the expected evaluation, actual valuation is %v", b.String(), b.EvalBoard())
}
