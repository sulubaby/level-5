package piscine

func ConcatAlternate(a, b []int) []int {
	if len(a) == 0 && len(b) == 0 {
        return nil
    }
    
    if len(a) < len(b) {
		a, b = b, a
	}
	r := []int{}
	for i := 0; i < len(a) || i < len(b); i++ {
		if i < len(a) {
			r = append(r, a[i])
		}
		if i < len(b) {
			r = append(r, b[i])
		}
	}
	return r
}
