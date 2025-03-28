package builder

// Operator is an interface that should be implemented by structs used as operators.
type Operator interface {
	GetKey() string
	GetVal() any
}

// BaseOperator is a simple operator struct that implementes the Operator interface.
type BaseOperator struct {
	key string
	val any
}

// GetKey function returns the operator's key.
func (operator *BaseOperator) GetKey() string {
	return operator.key
}

// GetVal function returns the operator's value.
func (operator *BaseOperator) GetVal() any {
	return operator.val
}

// New function creates a new base operator with the specified key and value
func New(key string, val any) Operator {
	return &BaseOperator{
		key: key,
		val: val,
	}
}

// Ensure that the BaseOperator implements the Operator interace
var _ Operator = &BaseOperator{}
