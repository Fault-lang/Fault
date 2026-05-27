package execute

type Branch struct {
	trail []int16
	phi   int16
	base  string
}

func (b *Branch) End() int16 {
	return b.trail[len(b.trail)-1]
}

func (b *Branch) InTrail(i int16) bool {
	for _, v := range b.trail {
		if i == v {
			return true
		}
	}
	return false
}

type Scenario interface{}

type FloatTrace struct {
	Scenario
	Base    string
	results map[int16]float64
}

func NewFloatTrace() *FloatTrace {
	return &FloatTrace{
		results: make(map[int16]float64),
	}
}

func (ft *FloatTrace) Index(i int16) (float64, bool) {
	v, ok := ft.results[i]
	return v, ok
}

func (ft *FloatTrace) Get() map[int16]float64 {
	return ft.results
}

func (ft *FloatTrace) Add(i int16, f float64) {
	ft.results[i] = f
}

func (ft *FloatTrace) Remove(i int16) {
	delete(ft.results, i)
}

type IntTrace struct {
	Scenario
	Base    string
	results map[int16]int64
}

func NewIntTrace() *IntTrace {
	return &IntTrace{
		results: make(map[int16]int64),
	}
}

func (it *IntTrace) Index(i int16) (int64, bool) {
	v, ok := it.results[i]
	return v, ok
}

func (it *IntTrace) Get() map[int16]int64 {
	return it.results
}

func (it *IntTrace) Add(i int16, in int64) {
	it.results[i] = in
}

func (it *IntTrace) Remove(i int16) {
	delete(it.results, i)
}

type BoolTrace struct {
	Scenario
	Base    string
	results map[int16]bool
}

func NewBoolTrace() *BoolTrace {
	return &BoolTrace{
		results: make(map[int16]bool),
	}
}

func (bt *BoolTrace) Index(i int16) (bool, bool) {
	v, ok := bt.results[i]
	return v, ok
}

func (bt *BoolTrace) Get() map[int16]bool {
	return bt.results
}

func (bt *BoolTrace) Add(i int16, b bool) {
	bt.results[i] = b
}

func (bt *BoolTrace) Remove(i int16) {
	delete(bt.results, i)
}
