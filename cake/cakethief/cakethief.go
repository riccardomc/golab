package cakethief

//CakeType is a type of cake!
type CakeType struct {
	weight int
	value  int
}

func MaxDuffleBagValue(cakeTypes []CakeType, capacity int) int {

	valueAtCapacity := make([]int, capacity+1)
	valueAtCapacity[0] = 0

	for currentCapacity := 0; currentCapacity <= capacity; currentCapacity++ {
		currentMaxValue := 0
		for _, cake := range cakeTypes {
			if cake.weight <= currentCapacity {
				maxValueUsingCake := cake.value + valueAtCapacity[currentCapacity-cake.weight]
				currentMaxValue = max(currentMaxValue, maxValueUsingCake)
			}
		}
		valueAtCapacity[currentCapacity] = currentMaxValue
	}

	return valueAtCapacity[capacity]
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
