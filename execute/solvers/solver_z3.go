package solvers

//import "C"

type Solver struct {
	Command   string
	Arguments []string
}

/*func (s *Solver) Run(smt string) (string, error) {
	spec := C.CString(smt)
	config := C.Z3_mk_config()
	c := C.Z3_mk_context(config)
	results, err := C.Z3_eval_smtlib2_string(c, spec)
	if results == nil {
		return "", error(err)
	}
	ret := C.GoString(results)
	C.free(unsafe.Pointer(spec))
	return ret, nil
}*/

func Z3() map[string]*Solver {
	s := make(map[string]*Solver)
	s["basic_run"] = &Solver{
		Command: "z3",
		//Command:   "solvers/z3/bin/z3",
		Arguments: []string{"-in"}}
	return s
}
