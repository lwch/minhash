package minhash

func Jaccard(a, b Hash) float64 {
	return float64(len(intersection(a, b))) / float64(len(union(a, b)))
}

func intersection(a, b []int) []int {
	ret := make([]int, 0)
	for _, v := range a {
		if contains(b, v) {
			ret = append(ret, v)
		}
	}
	return ret
}

func union(a, b []int) []int {
	ret := make([]int, len(a))
	copy(ret, a)
	for _, v := range b {
		if !contains(ret, v) {
			ret = append(ret, v)
		}
	}
	return ret
}

func contains(a []int, v int) bool {
	for _, x := range a {
		if x == v {
			return true
		}
	}
	return false
}
