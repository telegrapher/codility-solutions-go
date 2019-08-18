package frogjmp

func Solution(X int, Y int, Z int) int {
	distance := (Y - X)
	jumps := distance / Z
	if distance == 0 {
		jumps = 0
	} else if jumps == 0 && distance > 0 {
		jumps = 1
	} else if (distance % jumps) != 0 {
		jumps++
	}
	return jumps
}
