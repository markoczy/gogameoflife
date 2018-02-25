package game

var prefabs = initPrefabs()

func initPrefabs() map[string]Grid {
	ret := make(map[string]Grid)
	ret["glider"] = Grid{
		Data: [][]bool{
			[]bool{false, true, false},
			[]bool{false, false, true},
			[]bool{true, true, true}},
		Width:  3,
		Height: 3}

	ret["test"] = Grid{
		Data: [][]bool{
			[]bool{false, true, true, true},
			[]bool{false, false, true, false},
			[]bool{false, false, false, true}},
		Width:  4,
		Height: 3}

	return ret
}

// GetPrefab ...
func GetPrefab(name string, rotation int, flipX, flipY bool) Grid {
	return prefabs[name].RotateFlip(rotation, flipX, flipY)
}
