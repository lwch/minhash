package minhash

import (
	"fmt"
	"testing"
)

func TestMinHash(t *testing.T) {
	a := "flying fish flew by the space station"
	b := "we will not allow you to bring your pet armadillo along"
	c := "he figured a few sticks of dynamite were easier than a fishing pole to catch fish"
	h := New(32)
	sa := h.Add(a)
	sb := h.Add(b)
	sc := h.Add(c)
	ha := h.Hash(sa)
	hb := h.Hash(sb)
	hc := h.Hash(sc)
	fmt.Println(ha)
	fmt.Println(hb)
	fmt.Println(hc)
	fmt.Println(jaccard(sa, sb), Jaccard(ha, hb))
	fmt.Println(jaccard(sa, sc), Jaccard(ha, hc))
}

func jaccard(a, b ShingleSet) float64 {
	return float64(len(a.Intersection(b))) / float64(len(a.Union(b)))
}
