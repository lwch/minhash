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

type shingleSet map[shingle]int

func (ss shingleSet) Intersection(other shingleSet) shingleSet {
	ret := make(shingleSet)
	for k, v := range ss {
		if v2, ok := other[k]; ok {
			ret[k] = min(v, v2)
		}
	}
	return ret
}

func (ss shingleSet) Union(other shingleSet) shingleSet {
	ret := make(shingleSet)
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
