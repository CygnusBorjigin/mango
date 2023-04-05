package mangoDataType

type ConditionType int64

const (
	StringComparison ConditionType = iota
	NumericComparison
	Undefined
)

func (c ConditionType) String() string {
	switch c {
	case StringComparison:
		return "StringComparison"
	case NumericComparison:
		return "NumericComparison"
	case Undefined:
		return "Undefined"
	default:
		return "Not a conditionType"
	}
}
