package mangoDataType

type NumericRelation int64

const (
	Equal NumericRelation = iota
	Lesser
	LesserEqual
	Greater
	GreaterEqual
	NotEqual
)

func (c NumericRelation) String() string {
	switch c {
	case Equal:
		return "="
	case Lesser:
		return "<"
	case LesserEqual:
		return "<="
	case Greater:
		return ">"
	case GreaterEqual:
		return ">="
	case NotEqual:
		return "!="
	default:
		return "undefined"
	}
}
