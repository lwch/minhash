package minhash

import (
	"math/rand"
	"sort"
	"sync"
)

type MinHash struct {
	dict       *dict
	hashSize   int
	m          sync.Mutex
	vocab      map[Shingle]struct{}
	cacheVocab []Shingle
	hashFuncs  [][]int
}

type Hash []int

func New(hashSize int) *MinHash {
	return &MinHash{
		dict:     newDict(),
		hashSize: hashSize,
		vocab:    make(map[Shingle]struct{}),
	}
}

func (h *MinHash) Add(str string) ShingleSet {
	if h.hashFuncs != nil {
		panic("hash functions already initialized")
	}
	ret := make(ShingleSet)
	prev := -1
	for i, ch := range str {
		if i == 0 {
			prev = h.dict.Add(ch)
			continue
		}
		var sh Shingle
		next := h.dict.Add(ch)
		sh[0] = prev
		sh[1] = next
		ret[sh]++
		prev = next
		h.m.Lock()
		h.vocab[sh] = struct{}{}
		h.m.Unlock()
	}
	return ret
}

func (h *MinHash) Hash(ss ShingleSet) Hash {
	if h.hashFuncs == nil {
		h.buildHashFuncs()
	}
	vocabs := h.vocabs()
	ret := make(Hash, h.hashSize)
	for i, hf := range h.hashFuncs {
		ret[i] = h.hash(ss, hf, vocabs)
	}
	return ret
}

func (h *MinHash) vocabs() []Shingle {
	if h.cacheVocab != nil {
		return h.cacheVocab
	}
	h.m.Lock()
	defer h.m.Unlock()
	ret := make([]Shingle, len(h.vocab))
	i := 0
	for sh := range h.vocab {
		ret[i] = sh
		i++
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i][0] < ret[j][0] || (ret[i][0] == ret[j][0] && ret[i][1] < ret[j][1])
	})
	h.cacheVocab = ret
	return ret
}

func (h *MinHash) hash(ss ShingleSet, hf []int, vocabs []Shingle) int {
	for _, fn := range hf {
		s := vocabs[fn]
		if _, ok := ss[s]; ok {
			return fn
		}
	}
	panic("hash function not found")
}

func (h *MinHash) buildHashFuncs() {
	size := len(h.vocab)
	tmp := make([]int, size)
	for i := range tmp {
		tmp[i] = i
	}
	shuffle := func() []int {
		rand.Shuffle(size, func(i, j int) {
			tmp[i], tmp[j] = tmp[j], tmp[i]
		})
		return append([]int(nil), tmp...)
	}
	h.hashFuncs = make([][]int, h.hashSize)
	for i := range h.hashFuncs {
		h.hashFuncs[i] = shuffle()
	}
}
