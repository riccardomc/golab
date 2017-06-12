package arraysstrings

import "fmt"

// Rotate a matrix M 90 degrees
func Rotate(M [][]int) {

	N := len(M)

	rotatePixels := func(i, j int) {
		nextJ := N - 1 - i
		nextI := j
		for k := 0; k < 3; k++ {
			fmt.Println(i, j, nextI, nextJ)
			temp := M[i][j]
			M[i][j] = M[nextI][nextJ]
			M[nextI][nextJ] = temp
			tempI := nextI
			nextI = nextJ
			nextJ = N - 1 - tempI
		}
	}

	for i := 0; i < N/2; i++ {
		for j := i; j < N-1-i; j++ {
			fmt.Printf("Rotate(%d,%d)\n", i, j)
			rotatePixels(i, j)
		}
	}
}
