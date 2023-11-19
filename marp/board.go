package main

import "fmt"

type Board struct {
	grid [rows][cols]string
}

func NewBoard() *Board {
	var b Board
	for i := range b.grid {
		for j := range b.grid[i] {
			b.grid[i][j] = " "
		}
	}
	return &b
}

func (b *Board) Reset() {
	for i := range b.grid {
		for j := range b.grid[i] {
			b.grid[i][j] = " "
		}
	}
}

func (b *Board) DrawPiece(row, col int, piece string) {
	b.grid[row][col] = piece
}

func printBoard(b *Board) {
	for _, row := range b.grid {
		for _, cell := range row {
			fmt.Printf("| %s ", cell)
		}
		fmt.Println("|")
	}
}
