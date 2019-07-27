// Code generated by "stringer -type=MetricSemantics"; DO NOT EDIT

package speed

import "fmt"

const (
	_MetricSemantics_name_0 = "NoSemanticsCounterSemantics"
	_MetricSemantics_name_1 = "InstantSemanticsDiscreteSemantics"
)

var (
	_MetricSemantics_index_0 = [...]uint8{0, 11, 27}
	_MetricSemantics_index_1 = [...]uint8{0, 16, 33}
)

func (i MetricSemantics) String() string {
	switch {
	case 0 <= i && i <= 1:
		return _MetricSemantics_name_0[_MetricSemantics_index_0[i]:_MetricSemantics_index_0[i+1]]
	case 3 <= i && i <= 4:
		i -= 3
		return _MetricSemantics_name_1[_MetricSemantics_index_1[i]:_MetricSemantics_index_1[i+1]]
	default:
		return fmt.Sprintf("MetricSemantics(%d)", i)
	}
}
