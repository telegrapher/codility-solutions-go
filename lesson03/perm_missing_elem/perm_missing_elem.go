package perm_missing_elem

func Solution(permutation []int) int {
	hashMap := make(map[int]int)
	for i := 0; i < len(permutation); i++ {
		hashMap[permutation[i]] = 1
	}
	for i := 1; i <= len(permutation)+1; i++ {
		_, ok := hashMap[i]
		if !ok {
			return i
		}
	}
	return 0
}
