package stolendrone

//FindUnique find a unique ID in the ID list
func FindUnique(ids []int) int {
	idMap := make(map[int]bool, 0)

	for _, id := range ids {
		_, ok := idMap[id]
		if ok {
			delete(idMap, id)
		} else {
			idMap[id] = true
		}
	}

	for k := range idMap {
		return k
	}

	return -1
}

//FindUniqueXOR find a unique ID in the ID list
func FindUniqueXOR(ids []int) int {

	unique := 0

	for _, id := range ids {
		unique = unique ^ id
	}

	return unique
}
