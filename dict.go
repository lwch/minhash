package minhash

import "sync"

type dict struct {
	sync.RWMutex
	rune2id map[rune]int
	id2rune map[int]rune
}

func newDict() *dict {
	var d dict
	d.rune2id = make(map[rune]int)
	d.id2rune = make(map[int]rune)
	return &d
}

func (d *dict) Rune(id int) (rune, bool) {
	d.RLock()
	defer d.RUnlock()
	r, ok := d.id2rune[id]
	return r, ok
}

func (d *dict) ID(r rune) (int, bool) {
	d.RLock()
	defer d.RUnlock()
	id, ok := d.rune2id[r]
	return id, ok
}

func (d *dict) Add(r rune) int {
	d.Lock()
	defer d.Unlock()
	id, ok := d.rune2id[r]
	if ok {
		return id
	}
	id = len(d.rune2id)
	d.rune2id[r] = id
	d.id2rune[id] = r
	return id
}
