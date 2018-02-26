package grid

// Grid binary matrix with height and width
type Grid struct {
	Data          [][]bool
	Width, Height int
}

// CreateEmpty creates an empty Grid with desired width and height
func CreateEmpty(width, height int) *Grid {
	data := make([][]bool, height)
	for iRow := range data {
		data[iRow] = make([]bool, width)
	}
	return &Grid{Data: data, Width: width, Height: height}
}

// Copy creates a copy with all values
func (g *Grid) Copy() *Grid {
	data := make([][]bool, g.Height)
	for iRow, row := range g.Data {
		data[iRow] = make([]bool, g.Width)
		for iCol, cell := range row {
			data[iRow][iCol] = cell
		}
	}
	return &Grid{
		Data:   data,
		Width:  g.Width,
		Height: g.Height}
}

// GetCellAt retreives cell value at position
func (g *Grid) GetCellAt(x, y int) bool {
	col, row := getPos(x, y, g.Width, g.Height)
	return g.Data[row][col]
}

// CountNeighboursAt counts neighbours at position
func (g *Grid) CountNeighboursAt(x, y int) int {
	ret := 0
	for iX := x - 1; iX < x+2; iX++ {
		for iY := y - 1; iY < y+2; iY++ {
			if !(iX == x && iY == y) && g.GetCellAt(iX, iY) {
				ret++
			}
		}
	}
	return ret
}

// RotateFlip rotate and/or flip Prefab
func (g *Grid) RotateFlip(rotation int, flipX, flipY bool) {
	// _                        _
	//  |  =>  _|  =>  |_  =>  |
	//
	//  0      1       2       3
	//
	rot := rotation % 4
	height := g.Height
	width := g.Width
	if rot%2 == 1 {
		height = g.Width
		width = g.Height
	}
	data := make([][]bool, height)
	for iRow := range data {
		data[iRow] = make([]bool, width)
		var row, col int
		for iCol := range data[iRow] {
			switch rot {
			case 0:
				row = iRow
				col = iCol
				if flipY {
					row = height - iRow - 1
				}
				if flipX {
					col = width - iCol - 1
				}
			case 1:
				row = g.Height - iCol - 1
				col = iRow
				if flipX {
					row = iCol
				}
				if flipY {
					col = g.Width - iRow - 1
				}
			case 2:
				row = g.Height - iRow - 1
				col = g.Width - iCol - 1
				if flipY {
					row = iRow
				}
				if flipX {
					col = iCol
				}
			case 3:
				row = iCol
				col = g.Width - iRow - 1
				if flipX {
					row = g.Height - iCol - 1
				}
				if flipY {
					col = iRow
				}
			}
			data[iRow][iCol] = g.Data[row][col]
		}
	}
	g.Data = data
}

// InsertGridAt insert another grid at position
func (g *Grid) InsertGridAt(grid *Grid, x, y int) {
	for iRow, row := range grid.Data {
		for iCol := range row {
			col, row := getPos(iCol+x, iRow+y, g.Width, g.Height)
			if grid.Data[iRow][iCol] {
				g.Data[row][col] = true
			}
		}
	}
}

func (g *Grid) String() string {
	ret := ""
	for _, row := range g.Data {
		for _, cell := range row {
			if cell {
				ret += "+"
			} else {
				ret += "-"
			}
		}
		ret += "\n"
	}
	return ret
}

func getPos(x, y, width, height int) (int, int) {
	return mod(x, width), mod(y, height)
}

func mod(val, mod int) int {
	if val >= 0 {
		return val % mod
	}
	for val < 0 {
		val += mod
	}
	return val
}
