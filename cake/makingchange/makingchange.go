package makingchange

import "fmt"

//MakeChange returns the nlasnlkanf
func MakeChange(denominations []int, amount int) {

	var findSolutions func([]int)

	sum := func(candidate []int) int {
		sum := 0
		for i := 0; i < len(candidate); i++ {
			sum += candidate[i]
		}
		return sum
	}

	findSolutions = func(candidate []int) {
		for i := 0; i < len(denominations); i++ {
			candidate = append(candidate, denominations[i])
			sum := sum(candidate)
			if sum == amount {
				fmt.Println(candidate)
			}
			if sum > amount {
				continue
			}
			findSolutions(candidate)
			candidate = candidate[:len(candidate)-1]
		}
	}

	findSolutions([]int{})

}

//MakeChangeBottomUp yo
func MakeChangeBottomUp(denominations []int, amount int) int {
	solutions := make([]int, amount+1)
	solutions[0] = 1

	for _, coin := range denominations {
		for higherAmount := coin; higherAmount <= amount; higherAmount++ {
			solutions[higherAmount] += solutions[higherAmount-coin]
		}
	}

	fmt.Println(solutions)

	return solutions[amount]
}
