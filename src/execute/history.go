package execute

type Scenario interface {
	IsNil() bool
}

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

func (fh *FloatTrace) Get() map[int64]float64 {
	return fh.results
}

func (fh *FloatTrace) GetWeights() map[int64]float64 {
	return fh.weights
}

func (fh *FloatTrace) Add(i int64, f float64) {
	fh.results[i] = f
}

func (fh *FloatTrace) AddWeight(i int64, f float64) {
	fh.weights[i] = f
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
