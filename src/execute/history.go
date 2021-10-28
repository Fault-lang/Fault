package execute

type Scenario interface{}

type FloatTrace struct {
	Scenario
	results map[int64]float64
	weights map[int64]float64
}

func NewFloatTrace() *FloatTrace {
	return &FloatTrace{
		results: make(map[int64]float64),
		weights: make(map[int64]float64),
	}
}

func (ft *FloatTrace) Index(i int64) (float64, bool) {
	v, ok := ft.results[i]
	return v, ok
}

func (ft *FloatTrace) Get() map[int64]float64 {
	return ft.results
}

func (ft *FloatTrace) GetWeights() map[int64]float64 {
	return ft.weights
}

func (ft *FloatTrace) Add(i int64, f float64) {
	ft.results[i] = f
}

func (ft *FloatTrace) AddWeight(i int64, f float64) {
	ft.weights[i] = f
}

func (ft *FloatTrace) Remove(i int64) {
	delete(ft.results, i)
	delete(ft.weights, i)
}

type IntTrace struct {
	Scenario
	results map[int64]int64
	weights map[int64]float64
}

func NewIntTrace() *IntTrace {
	return &IntTrace{
		results: make(map[int64]int64),
		weights: make(map[int64]float64),
	}
}

func (it *IntTrace) Index(i int64) (int64, bool) {
	v, ok := it.results[i]
	return v, ok
}

func (it *IntTrace) Get() map[int64]int64 {
	return it.results
}

func (it *IntTrace) GetWeights() map[int64]float64 {
	return it.weights
}

func (it *IntTrace) Add(i int64, in int64) {
	it.results[i] = in
}

func (it *IntTrace) AddWeight(i int64, f float64) {
	it.weights[i] = f
}

func (it *IntTrace) Remove(i int64) {
	delete(it.results, i)
	delete(it.weights, i)
}

type BoolTrace struct {
	Scenario
	results map[int64]bool
	weights map[int64]float64
}

func NewBoolTrace() *BoolTrace {
	return &BoolTrace{
		results: make(map[int64]bool),
		weights: make(map[int64]float64),
	}
}

func (bt *BoolTrace) Index(i int64) (bool, bool) {
	v, ok := bt.results[i]
	return v, ok
}

func (bt *BoolTrace) Get() map[int64]bool {
	return bt.results
}

func (bt *BoolTrace) GetWeights() map[int64]float64 {
	return bt.weights
}

func (bt *BoolTrace) Add(i int64, b bool) {
	bt.results[i] = b
}

func (bt *BoolTrace) AddWeight(i int64, f float64) {
	bt.weights[i] = f
}

func (bt *BoolTrace) Remove(i int64) {
	delete(bt.results, i)
	delete(bt.weights, i)
}
