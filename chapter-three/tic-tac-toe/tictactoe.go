package tictactoe

type game struct {
	dimension uint
	row       []*boardPart
	col       []*boardPart
	diagonals [2]*boardPart
}

type boardPart struct {
	player        bool
	occupiedCount int
	denied        bool
}

func New(dimension uint) *game {
	return &game{
		dimension: dimension,
		row:       make([]*boardPart, dimension),
		col:       make([]*boardPart, dimension),
		diagonals: [2]*boardPart{},
	}
}

func (g *game) Move(row, col uint, player bool) bool {
	if g.row[row] == nil {
		g.row[row] = &boardPart{player: player, occupiedCount: 1}
	} else if !g.row[row].denied {
		if g.row[row].player != player {
			g.row[row].denied = true
		} else {
			g.row[row].occupiedCount++
			if g.row[row].occupiedCount == int(g.dimension) {
				return true
			}
		}
	}

	if g.col[col] == nil {
		g.col[col] = &boardPart{player: player, occupiedCount: 1}
	} else if !g.col[col].denied {
		if g.col[col].player != player {
			g.col[col].denied = true
		} else {
			g.col[col].occupiedCount++
			if g.col[col].occupiedCount == int(g.dimension) {
				return true
			}
		}
	}

	if row == col {
		if g.diagonals[0] == nil {
			g.diagonals[0] = &boardPart{player: player, occupiedCount: 1}
		} else if !g.diagonals[0].denied {
			if g.diagonals[0].player != player {
				g.diagonals[0].denied = true
			} else {
				g.diagonals[0].occupiedCount++
				if g.diagonals[0].occupiedCount == int(g.dimension) {
					return true
				}
			}
		}
	}

	if row == g.dimension-1-col {
		if g.diagonals[1] == nil {
			g.diagonals[1] = &boardPart{player: player, occupiedCount: 1}
		} else if !g.diagonals[1].denied {
			if g.diagonals[1].player != player {
				g.diagonals[1].denied = true
			} else {
				g.diagonals[1].occupiedCount++
				if g.diagonals[1].occupiedCount == int(g.dimension) {
					return true
				}
			}
		}
	}

	return false
}

// TODO how could I refactor the above?
func setMove(part *boardPart, player bool) bool {
	return false
}
