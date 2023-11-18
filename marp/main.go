package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	START   = "<!-- Connect 4 Board Start -->"
	END     = "<!-- Connect 4 Board End -->"
	PLAYER1 = "🟠"
	PLAYER2 = "🔵"
	rows    = 6
	cols    = 7
)

func main() {
	file, err := os.Open("gameboard.md")
	if err != nil {
		fmt.Println("Could not locate board!")
	}
	defer file.Close()

	boardFound := false
	var board [6][7]string
	//boardRow := 0
	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		if line == "|   |   |   |   |   |   |   |" || strings.Contains(line, "---") {
			continue // Skip the header and divider lines
		}

		if strings.Contains(line, "<!-- Connect 4 Board Start -->") {
			boardFound = true
			board = startGame()
			lines = append(lines, generateBoardLines(board)...)
			printBoard(board)
			continue
		}

		if boardFound && strings.Contains(line, "<!-- Connect 4 Board End -->") {
			boardFound = false
			lines = append(lines, line)
			break
		}

		if boardFound {
			// cells := strings.Split(line, "|")
			// for i, cell := range cells {
			// 	if i >= cols || boardRow > rows {
			// 		break
			// 	}
			// 	fmt.Printf("cell: %s\n", cell)
			// 	cell = strings.TrimSpace(cell)

			// 	switch cell {
			// 	case PLAYER1:
			// 		board[boardRow][i] = "Y"
			// 		fmt.Println("a")
			// 	case PLAYER2:
			// 		board[boardRow][i] = "B"
			// 		fmt.Println(PLAYER2)
			// 	default:
			// 		board[boardRow][i] = ""
			// 	}
			// }
			// boardRow++
			continue
		}
	}
	err = writeToFile("gameboard.md", lines)
	if err != nil {
		fmt.Println("Error writing to file: ", err)
	}
}

func startGame() [rows][cols]string {
	var board [rows][cols]string
	boardInit := make(chan bool)
	go func() {
		for i := range board {
			for j := range board[i] {
				board[i][j] = ""
			}
		}
		boardInit <- true
	}()

	<-boardInit
	close(boardInit)
	fmt.Println("Board initialized")
	return board
}

func printBoard(board [rows][cols]string) {
	for _, row := range board {
		for _, cell := range row {
			fmt.Printf("| %s ", cell)
		}
		fmt.Println("|")
	}
}

func generateBoardLines(board [6][7]string) []string {
	var boardLines []string
	for _, row := range board {
		rowLine := "|"
		for _, cell := range row {
			if cell == "Y" {
				rowLine += " 🟠 |"
			} else if cell == "B" {
				rowLine += " 🔵 |"
			} else {
				rowLine += "   |"
			}
		}
		boardLines = append(boardLines, rowLine)
	}
	return boardLines
}

func writeToFile(filename string, lines []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}
