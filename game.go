package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func resetBoard(brd [][]string) {
	for i := 0; i < len(brd); i++ {
		for j := 0; j < len(brd[i]); j++ {
			brd[i][j] = " "
		}
	}
}

func showBoard(brd [][]string) {
	numRow := len(brd)
	numCol := len(brd[0])

	fmt.Println()

	// First write the column header
	fmt.Print("       ")
	for i := 0; i < numCol; i++ {
		// fmt.Print(strconv.Itoa(i) + "   ")
		fmt.Print(strconv.Itoa(i) + "      ")
	}
	fmt.Printf("\n    ")

	for i := 0; i < numCol; i++ {
		fmt.Printf(strings.Repeat("-", 6))

		if i != numCol-1 {
			fmt.Printf("+")
		}
	}

	fmt.Println() // blank line after the header

	// The write the table
	for i := 0; i < numRow; i++ {
		fmt.Print(strconv.Itoa(i))
		fmt.Printf(strings.Repeat(" ", 2))
		for j := 0; j < numCol; j++ {
			fmt.Print("|")
			if brd[i][j] == " " {
				fmt.Printf(strings.Repeat(" ", 6))
			} else {
				fmt.Printf(strings.Repeat(" ", 3))
				fmt.Print(brd[i][j])
				fmt.Printf(strings.Repeat(" ", 2))
			}
		}

		fmt.Print("|")
		fmt.Printf(strings.Repeat(" ", 6))

		fmt.Printf("\n    ")

		if i != (numRow - 1) {
			// separator line
			for j := 0; j < numCol; j++ {
				if j != 0 {
					fmt.Print("+")
				}

				fmt.Printf(strings.Repeat("-", 6))
			}
			fmt.Println()
		}
	}

	for i := 0; i < numCol; i++ {
		fmt.Printf(strings.Repeat("-", 6))

		if i != numCol-1 {
			fmt.Printf("+")
		}
	}

	fmt.Println()
}

func userPlay(brd [][]string, usym string) {
	fmt.Print("\nEnter the row and column indices: ")
	var rowIndex, colIndex int
	fmt.Scan(&rowIndex, &colIndex)

	for brd[rowIndex][colIndex] != " " {
		fmt.Print("\n!! The cell is already taken. \n Enter the row and column indices: ")
		fmt.Scan(&rowIndex, &colIndex)
	}

	brd[rowIndex][colIndex] = usym
}

func compPlay(brd [][]string, csym string) {
	// Find the first empty cell and put a tic there
	rows := len(brd)    // Number of rows
	cols := len(brd[0]) // Number of columns

	// Create a random number generator
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < rows*cols; i++ {
		// Generate a random integer between 0 and rows/cols
		rand_row := rand.Intn(rows)
		rand_col := rand.Intn(cols)
		fmt.Print(rand_row)
		fmt.Print(rand_row)

		if brd[rand_row][rand_col] == " " { // empty cell
			brd[rand_row][rand_col] = csym
			return
		}
	}
}

func isGameWon(brd [][]string, turn int, usym string, csym string) bool {
	var sym string

	if turn == 0 {
		sym = usym
	} else {
		sym = csym
	}

	win := false
	rows := len(brd)
	cols := len(brd[0])

	// Check win by a row
	i := 0
	j := 0
	for i < rows && !win {
		for j < cols {
			if brd[i][j] != sym {
				break
			}
			j++
		}
		if j == len(brd[0]) {
			win = true
		}
		i++
	}

	// Check win by a column
	i = 0
	j = 0
	for j < cols && !win {
		for i < rows {
			if brd[i][j] != sym {
				break
			}
			i++
		}
		if i == len(brd) {
			win = true
		}
		j++
	}

	// Check win by a diagonal (1)
	i = 0
	j = 0
	if !win {
		for i < rows {
			if brd[i][i] != sym {
				break
			}
			i++
		}
		if i == len(brd) {
			win = true
		}
	}

	// Check win by a diagonal (2)
	i = 0
	j = 0
	if !win {
		for i < rows {
			if brd[i][rows-1-i] != sym {
				break
			}
			i++
		}
		if i == len(brd) {
			win = true
		}
	}

	// Finally return win
	return win
}

func createBoard(size int) [][]string {
	// Create a board to play tiktaktoe
	board := make([][]string, size)
	for i := 0; i < size; i++ {
		board[i] = make([]string, size)
	}

	return board
}

func main() {
	// First, welcome message and display the board.
	fmt.Printf("===== WELCOME TO THE TIC-TAC-TOE GAME!! =====\n\n")
	fmt.Printf("Which size of board do you like to play? ")
	var size int
	fmt.Scanln(&size)

	// Create a board
	board := createBoard(size)

	// Initialize the board with " " in all cells
	resetBoard(board)

	// Show the inital board
	showBoard(board)

	// Then ask the user which symbol (x or o) he/she wants to play.
	fmt.Printf("Which symbol do you want to play, \"x\" or \"o\"? ")
	var userSymbol, compSymbol string
	fmt.Scanln(&userSymbol)
	userSymbol = strings.ToLower(userSymbol)
	if userSymbol == "x" {
		compSymbol = "o"
	} else {
		compSymbol = "x"
	}

	// Also ask whether or not the user wants to go first
	fmt.Println()
	fmt.Print("Do you want to go first (y/n)? ")
	var ans string
	fmt.Scanln(&ans)
	ans = strings.ToLower(ans)[:1]

	turn := 0                  // 0 -- the user, 1 -- the computer
	remainCount := size * size // empty the cell call

	// THE VERY FIRST MOVE
	if ans == "y" {
		turn = 0
		userPlay(board, userSymbol) // user puts his/her first tic
	} else {
		turn = 1
		compPlay(board, compSymbol) // computer puts its first tic
	}

	// Show the board, and decrement the count of remaining cells
	showBoard(board)
	remainCount--

	// Play the game until either one wins.
	done := false
	winner := -1 // 0 -- the user, 1 -- the computer, -1 -- draw

	for !done && remainCount > 0 {
		// If there is a winner at this time, set the winner and the done flag to true.
		done = isGameWon(board, turn, userSymbol, compSymbol) // Did the turn won?

		if done {
			winner = turn // the one who made the last move won the game
		} else {
			// No winner yet.  Find the next turn and play.
			turn = (turn + 1) % 2

			if turn == 0 {
				userPlay(board, userSymbol)
			} else {
				compPlay(board, compSymbol)
			}

			// Show the board after one tic, and decrement the rem count.
			showBoard(board)
			remainCount--
		}
	}

	// Winner is found.  Declare the winner.
	if winner == 0 {
		fmt.Println("\n** YOU WON.  CONGRATULATIONS!! **")
	} else if winner == 1 {
		fmt.Println("\n** YOU LOST..  Maybe next time :) **")
	} else {
		fmt.Println("\n** DRAW... **")
	}
}
