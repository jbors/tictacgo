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

	board = board.PlayRandomMove()

	var result jsonState

	result.State = board.String()
	result.Result = board.EvalBoard().String()

	//This just gives back the state but we should add a move first
	encoder := json.NewEncoder(writer)
	encoder.Encode(result)
}
