package random

import "math/rand"

type Rand struct {
	*rand.Rand
}

func New(src rand.Source) *Rand {
	return &Rand{rand.New(src)}
}

func (m *Rand) Range(min, max int) int {
	return m.Intn(max-min) + min
}

func (m *Rand) Roll_Dice(n, die_type int) int {
	sum := 0
	for range make([]int, n) {
		sum += m.Range(1, die_type+1)
	}
	return sum
}
