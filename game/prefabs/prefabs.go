package prefabs

import "github.com/markoczy/gogameoflife/game/grid"

var prefabs = initPrefabs()

func initPrefabs() map[string]*grid.Grid {
	ret := make(map[string]*grid.Grid)
	ret["glider"] = &grid.Grid{
		Data: [][]bool{
			[]bool{false, true, false},
			[]bool{false, false, true},
			[]bool{true, true, true}},
		Width:  3,
		Height: 3}

	ret["test"] = &grid.Grid{
		Data: [][]bool{
			[]bool{false, true, true, true},
			[]bool{false, false, true, false},
			[]bool{false, false, false, true}},
		Width:  4,
		Height: 3}

	return ret
}

// Get retreives Prefab by name
func Get(name string, rotation int, flipX, flipY bool) *grid.Grid {
	ret := prefabs[name].Copy()
	if rotation != 0 || flipX || flipY {
		ret.RotateFlip(rotation, flipX, flipY)
	}
	return ret
}
