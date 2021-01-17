# Tic Tac Go

Small tic-tac-toe application in [Go](https://golang.org/) to explore some concepts

This is the first thing I wrote in Go after taking a coursera course on it so I'm sure some improvements could be made.

## Installation

1. Install Go: https://golang.org/doc/install
2. Clone the repository into `$GOPATH/src/your/folder`
3. Go to the root directory of the project
4. `go run .\cmd\` will start the program and expose the game at `localhost:8080`

![Game board](/assets/board70.png)

## TODO

- Allow user to choose whether he's X or O
- Algorithm improvements: alfabeta pruning, negamax
- Don't play for a tie when a win is still possible
- Play around with parallel execution
- Write some tests
