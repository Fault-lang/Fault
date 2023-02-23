package forks

import "sort"

// Key is the base variable name
type Fork map[string][]*Choice

type PhiState struct {
	levels int
}

func NewPhiState() *PhiState {
	return &PhiState{
		levels: 0,
	}
}

func (p *PhiState) Check() bool {
	return p.levels > 0
}

func (p *PhiState) Level() int {
	return p.levels
}

func (p *PhiState) In() {
	p.levels = p.levels + 1
}

func (p *PhiState) Out() {
	if p.levels != 0 {
		p.levels = p.levels - 1
	}
}

func GetForkEndPoints(c []*Choice) []int16 {
	var ends []int16
	for _, v := range c {
		e := v.Values[len(v.Values)-1]
		ends = append(ends, e)
	}
	return ends
}

type Choice struct {
	Base   string  // What variable?
	Branch string  // For conditionals, is this the true block or false block?
	Values []int16 // All the versions of this variable in this branch
}

func (c *Choice) AddChoiceValue(n int16) *Choice {
	c.Values = append(c.Values, n)
	sort.Slice(c.Values, func(i, j int) bool { return c.Values[i] < c.Values[j] })
	return c
}

func (c *Choice) GetEnd() int16 {
	return c.Values[len(c.Values)-1]
}
