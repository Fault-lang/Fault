package solvers

type Solver struct {
	Command   string
	Arguments []string
}

func Z3() map[string]*Solver {
	s := make(map[string]*Solver)
	s["basic_run"] = &Solver{
		Command:   "execute/solvers/z3/bin/z3",
		Arguments: []string{"-in"}}
	return s
}
