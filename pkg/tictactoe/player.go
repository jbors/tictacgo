// Package tictactoe defines some algorithms to make the computer play
package tictactoe

import (
	"fmt"
	"math/rand"
)

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

//PlayRandomMove lets the computer play a random move
func (b Board) PlayRandomMove() Board {
	moves := b.generatePossibleMoves()
	index := rand.Intn(len(moves))
	b.Cells[moves[index]] = O
	return b
}

//TODO Both Negamax and alphabeta pruning could improve on this
//Also, theoretical draws with a winning move possible are better than theoretical draws without

//PlayMiniMaxMove lets the computer play a move
func (b Board) PlayMiniMaxMove() Board {
	moves := b.generatePossibleMoves()
	topMoveValue := 2
	var selectedMoves []int
	for _, move := range moves {
		moveVal := miniMax(b.PlayMove(move, OPlayer), XPlayer, true)
		if moveVal < topMoveValue {
			topMoveValue = moveVal
			selectedMoves = nil
			selectedMoves = append(selectedMoves, move)
		} else if moveVal == topMoveValue {
			selectedMoves = append(selectedMoves, move)
		}
	}
	fmt.Printf("Selected moves: %v", selectedMoves)

	//Pick a random candidate from the 'best' moves so we do not play
	//the same move every time
	index := rand.Intn(len(selectedMoves))
	b.Cells[selectedMoves[index]] = O
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
