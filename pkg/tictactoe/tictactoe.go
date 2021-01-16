// Package tictactoe defines the concepts we need to represent a game
// and defines some algorithms to make the computer play
// Adapted from https://github.com/shurcooL/tictactoe
package tictactoe

import (
	"fmt"
	"math/rand"
	"strings"
)

// State of a board cell.
type State uint8

// States of a board cell.
const (
	F State = iota // Free cell.
	X              // Cell with an X mark.
	O              // Cell with an O mark.
)

// Player X or O player
type Player uint8

// Players
const (
	XPlayer Player = iota // Human plays the Xs
	OPlayer               // Computer plays the Os
)

//Board representation
type Board struct {
	// Cells is a 3x3 matrix in row major order.
	// Cells[3*r + c] is the cell in the r'th row and c'th column.
	Cells [9]State
	//Should this have something to indicate whos turn it is??
	//NextMove Player
}

// Condition of the board configuration.
type Condition uint8

// Conditions of the board configuration.
const (
	NotEnd Condition = iota
	XWon
	OWon
	Tie
)

// EvalBoard evaluates the current board state to see whether this is
// NotEnd: Game still in play
// XWon: Player X won
// YWon: Player Y won
// Draw: No winner, no more moves to play
func (b Board) EvalBoard() Condition {
	var (
		x = (b.Cells[0] == X && b.Cells[1] == X && b.Cells[2] == X) || // Check all rows.
			(b.Cells[3] == X && b.Cells[4] == X && b.Cells[5] == X) ||
			(b.Cells[6] == X && b.Cells[7] == X && b.Cells[8] == X) ||

			(b.Cells[0] == X && b.Cells[3] == X && b.Cells[6] == X) || // Check all columns.
			(b.Cells[1] == X && b.Cells[4] == X && b.Cells[7] == X) ||
			(b.Cells[2] == X && b.Cells[5] == X && b.Cells[8] == X) ||

			(b.Cells[0] == X && b.Cells[4] == X && b.Cells[8] == X) || // Check all diagonals.
			(b.Cells[2] == X && b.Cells[4] == X && b.Cells[6] == X)

		o = (b.Cells[0] == O && b.Cells[1] == O && b.Cells[2] == O) || // Check all rows.
			(b.Cells[3] == O && b.Cells[4] == O && b.Cells[5] == O) ||
			(b.Cells[6] == O && b.Cells[7] == O && b.Cells[8] == O) ||

			(b.Cells[0] == O && b.Cells[3] == O && b.Cells[6] == O) || // Check all columns.
			(b.Cells[1] == O && b.Cells[4] == O && b.Cells[7] == O) ||
			(b.Cells[2] == O && b.Cells[5] == O && b.Cells[8] == O) ||

			(b.Cells[0] == O && b.Cells[4] == O && b.Cells[8] == O) || // Check all diagonals.
			(b.Cells[2] == O && b.Cells[4] == O && b.Cells[6] == O)

		freeCellsLeft = b.Cells[0] == F || b.Cells[1] == F || b.Cells[2] == F ||
			b.Cells[3] == F || b.Cells[4] == F || b.Cells[5] == F ||
			b.Cells[6] == F || b.Cells[7] == F || b.Cells[8] == F
	)

	switch {
	case x && !o:
		return XWon
	case o && !x:
		return OWon
	case !freeCellsLeft:
		return Tie
	default:
		return NotEnd
	}
}

func (p Player) opponent() Player {
	switch p {
	case XPlayer:
		return OPlayer
	case OPlayer:
		return XPlayer
	default:
		panic("Unknown player")
	}
}

//PlayMove writes to the board stae on the selected position
func (b Board) PlayMove(move int, p Player) Board {
	switch p {
	case XPlayer:
		b.Cells[move] = X
	case OPlayer:
		b.Cells[move] = O
	}
	return b
}

//PlayRandomMove lets the computer play a random move
func (b Board) PlayRandomMove() Board {
	moves := b.generatePossibleMoves()
	index := rand.Intn(len(moves))
	b.Cells[moves[index]] = O
	return b
}

//TODO Both Negamax and alphabeta pruning could improve on this

//PlayMiniMaxMove lets the computer play a move
func (b Board) PlayMiniMaxMove() Board {
	moves := b.generatePossibleMoves()
	topMoveValue := 2
	var selectedMove int
	for _, move := range moves {
		moveVal := miniMax(b.PlayMove(move, OPlayer), XPlayer, true)
		// fmt.Printf("MoveVal: %v\n", moveVal)
		if moveVal < topMoveValue {
			topMoveValue = moveVal
			selectedMove = move
		}
	}

	b.Cells[selectedMove] = O
	return b
}

func miniMax(b Board, p Player, maximizingPlayer bool) int {
	switch b.EvalBoard() {
	case XWon:
		return 1
	case OWon:
		return -1
	case Tie:
		return 0
	case NotEnd:
		if maximizingPlayer {
			value := -2
			moves := b.generatePossibleMoves()
			for _, m := range moves {
				minimaxVal := miniMax(b.PlayMove(m, p), p.opponent(), false)
				value = Max(value, minimaxVal)
			}
			return value
		} else {
			value := 2
			moves := b.generatePossibleMoves()
			for _, m := range moves {
				minimaxVal := miniMax(b.PlayMove(m, p), p.opponent(), true)
				value = Min(value, minimaxVal)
			}
			return value
		}
	default:
		panic("Impossible board")
	}
}

// Max I have absolutely no idea why this would not just be part of the standard library
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Min I have absolutely no idea why this would not just be part of the standard library
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func (b Board) generatePossibleMoves() []int {
	var free []int
	for i, c := range b.Cells {
		if c == F {
			free = append(free, i)
		}
	}
	return free
}

func (b Board) String() string {
	var sb strings.Builder
	for _, c := range b.Cells {
		switch c {
		case X:
			sb.WriteRune('X')
		case O:
			sb.WriteRune('O')
		case F:
			sb.WriteRune('-')
		}
	}
	return sb.String()
}

func (c Condition) String() string {
	var sb strings.Builder
	switch c {
	case XWon:
		sb.WriteString("XWon")
	case OWon:
		sb.WriteString("OWon")
	case Tie:
		sb.WriteString("Tie")
	default:
		sb.WriteString("NotEnd")
	}
	return sb.String()
}

//MakeBoard reads a string representation into a board
func MakeBoard(s string) Board {
	var b Board
	for i, c := range s {
		switch c {
		case 'X':
			b.Cells[i] = X
		case 'O':
			b.Cells[i] = O
		case '-':
			b.Cells[i] = F
		default:
			fmt.Printf("Illegal character %v in board representation string", c)
		}
	}
	return b
}