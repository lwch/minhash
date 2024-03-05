package minhash

type shingle [2]int

func (s shingle) String(dict *dict) string {
	var str string
	for i := 0; i < 2; i++ {
		r, ok := dict.Rune(s[i])
		if !ok {
			str += "?"
		} else {
			str += string(r)
		}
	}
	return str
}

type ShingleSet map[shingle]int

func (ss ShingleSet) Intersection(other ShingleSet) ShingleSet {
	ret := make(ShingleSet)
	for k, v := range ss {
		if v2, ok := other[k]; ok {
			ret[k] = min(v, v2)
		}
	}
	return ret
}

func (ss ShingleSet) Union(other ShingleSet) ShingleSet {
	ret := make(ShingleSet)
	for k, v := range ss {
		ret[k] = v
	}
	for k, v := range other {
		if v2, ok := ret[k]; ok {
			ret[k] = max(v, v2)
		} else {
			ret[k] = v
		}
	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
