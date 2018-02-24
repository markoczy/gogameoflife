package game

func NewGame(width, height int, seeder Seeder) (Game, error) {
	ret := game{ width: width, height: height}
	data, err := seeder.Seed(&ret)
	if err != nil { return nil, err }
	ret.data = data
	return &ret, nil
}
type Game interface {
	Width() int
	Height() int
	Tick() error
	Render() ([][]bool, error)
}

func (g *game) Width() int {
	return g.width
}

func (g *game) Height() int {
	return g.height
}


func (g *game) Render() ([][]bool, error) {
	return g.data, nil
}


func (g *game) Tick() error {
	data := make([][]bool, g.height)
	for iRow, row := range g.data {
		data[iRow] = make([]bool, g.width)
		for iCell, _ := range row {
			neighbours := g.countNeighbours(iCell, iRow)
			if !g.data[iRow][iCell] {
				if neighbours == 3 {
					data[iRow][iCell] = true // live
				}
			} else {
				if neighbours < 2 || neighbours > 3 {
					data[iRow][iCell] = false // die
				} else {
					data[iRow][iCell] = true
				}
			}
		}
	}
	g.data = data
	return nil
}

type game struct {
	data   [][]bool
	width  int
	height int
}

func (g *game) countNeighbours(x, y int) int {
	ret := 0
	for iX := x - 1; iX < x+2; iX++ {
		for iY := y - 1; iY < y+2; iY++ {
			if !(iX == x && iY == y) && g.getCellAt(iX, iY) {
				ret += 1
			}
		}
	}
	return ret
}

func (g *game) getCellAt(x, y int) bool {
	return g.data[mod(y, g.height)][mod(x, g.width)]
}

func mod(val, mod int) int {
	if val >= 0 {
		return val % mod
	} else {
		for val < 0 {
			val += mod
		}
		return val
	}
}