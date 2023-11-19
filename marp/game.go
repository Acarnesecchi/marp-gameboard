package main

type Game struct {
	players            []string
	currentPlayerIndex int
}

func NewGame(playerNames []string) (*Game, *Board) {
	return &Game{
		players:            playerNames,
		currentPlayerIndex: 0,
	}, NewBoard()
}

func (g *Game) NextPlayer() string {
	player := g.players[g.currentPlayerIndex]
	g.currentPlayerIndex = (g.currentPlayerIndex + 1) % len(g.players)
	return player
}
