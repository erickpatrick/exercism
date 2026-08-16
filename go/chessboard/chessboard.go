package chessboard

type File []bool

type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	occupied := 0

	if len(cb) == 0 {
		return 0
	}

	if file < "A" || file > "H" {
		return 0
	}

	for _, value := range cb[file] {
		if value {
			occupied++
		}
	}

	return occupied
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	occupied := 0

	if rank < 1 || rank > 8 {
		return 0
	}

	for _, file := range cb {
		if file[rank-1] {
			occupied++
		}
	}

	return occupied
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	return len(cb) * len(cb["A"])
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	occupied := 0

	for file := 'A'; file <= 'Z'; file++ {
		occupied += CountInFile(cb, string(file))
	}

	return occupied
}
