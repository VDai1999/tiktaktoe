# Tic-Tac-Toe Game
This is a simple implementation of the Tic-Tac-Toe game in Go. The game allows two players to play against each other on a nxn grid with n being chosen by players.

## Features
* Player vs. Computer: Players will play the game against the computers.
* Players can choose whether they want to start first.
* Valid Move Checking: Only valid moves are accepted during gameplay.
* Win Condition Checking: The game checks for a win condition after each move.
* Tie Condition Checking: The game checks for a tie condition when the board is full with no winner.
* Interactive User Interface: The game provides a user-friendly command-line interface for playing.

## How to Run
1. Make sure you have Go installed on your system. If not, you can download and install it from the official Go website: https://golang.org/
2. Clone or download the Tic-Tac-Toe game repository to your local machine.
3. Open a terminal or command prompt and navigate to the project's root directory.
4. Run the following command to start the game:
```
go run game.go 
```

## How to Play
1. The game will prompt players to make choices regarding the board size, then will display an empty nxn grid.
2. The game will next prompt players to choose the symbol they wish to use, and whether they want to take the first turn.
3. When prompted about players' move location, enter the row and column numbers of the desired cell to place the symbol (x or o) with the format like `0 0` for the top-left cell.
4. Players and the computer will take turns entering their moves.
5. The game will validate players' moves and display the updated board.
6. Play continues until either player or the computer wins, or the game ends in a tie.

Enjoy playing Tic-Tac-Toe!

## Acknowledgements
The game logic and structure are inspired by the classic Tic-Tac-Toe game.
