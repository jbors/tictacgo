package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	ttt "github.com/jbors/tictacgo/pkg/tictactoe"
)

type jsonState struct {
	State  string
	Result string
	Player string
}

func main() {
	rand.Seed(12)

	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.HandleFunc("/api/", handleRestRequest)
	http.ListenAndServe(":8080", nil)
}

func handleRestRequest(writer http.ResponseWriter, request *http.Request) {
	//Get payload
	decoder := json.NewDecoder(request.Body)
	var state jsonState
	err := decoder.Decode(&state)
	if err != nil {
		fmt.Printf("error processing %v", err)
	}

	board := ttt.MakeBoard(state.State)

	//Only play if the game is not over
	if board.EvalBoard() == ttt.NotEnd {
		switch state.Player {
		case "parallel":
			board = board.PlayParallelMinimaxMove()
		case "minimax":
			board = board.PlayMiniMaxMove()
		default:
			board = board.PlayRandomMove()
		}
	}

	var result jsonState

	result.State = board.String()
	result.Result = board.EvalBoard().String()

	//Give back the new state
	encoder := json.NewEncoder(writer)
	encoder.Encode(result)
}
