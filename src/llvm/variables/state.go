package variables

type State interface{}

type globalSEntry struct {
	State
	state int16
}

func NewGlobalSEntry(st int16) *globalSEntry {
	return &globalSEntry{
		state: st,
	}
}

func (ge *globalSEntry) get(id []string) int16 {
	return ge.state
}

type structSEntry struct {
	State
	state map[string]map[string]int16
}

func NewStructSEntry(id []string, st int16) *structSEntry {
	s := &structSEntry{
		state: make(map[string]map[string]int16),
	}
	s.state[id[1]] = make(map[string]int16)
	s.state[id[1]][id[2]] = st
	return s
}

func (se *structSEntry) get(id []string) int16 {
	return se.state[id[1]][id[2]]
}

func (se *structSEntry) update(id []string, st int16) {
	se.state[id[1]][id[2]] = st
}

func (se *structSEntry) increment(id []string) {
	se.state[id[1]][id[2]] = se.state[id[1]][id[2]] + 1
}

type instanceSEntry struct {
	State
	state map[string]map[string]map[string]int16
}

func NewInstanceSEntry(id []string, st int16) *instanceSEntry {
	s := &instanceSEntry{
		state: make(map[string]map[string]map[string]int16),
	}
	s.state[id[1]] = make(map[string]map[string]int16)
	s.state[id[1]][id[2]] = make(map[string]int16)
	s.state[id[1]][id[2]][id[3]] = st
	return s
}

func (ie *instanceSEntry) get(id []string) int16 {
	return ie.state[id[1]][id[2]][id[3]]
}

func (ie *instanceSEntry) update(id []string, st int16) {
	ie.state[id[1]][id[2]][id[3]] = st
}

func (ie *instanceSEntry) increment(id []string) {
	ie.state[id[1]][id[2]][id[3]] = ie.state[id[1]][id[2]][id[3]] + 1
}
