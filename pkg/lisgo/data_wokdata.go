package lisgo

type LisgoData interface {
	ToString() string
	ToBoolean() bool
	ToInteger() int64
	ToFloat() float64
	Equals(other LisgoData) bool
	GetType() int
	GetTypeName() string
	GetValue() interface{}
}

const (
	LisgoTypeNull      = 0
	LisgoTypeBoolean   = 1
	LisgoTypeInteger   = 2
	LisgoTypeFloat     = 3
	LisgoTypeString    = 4
	LisgoTypeCallable  = 5
	LisgoTypeFunction  = 6
	LisgoTypeReturn    = 7
	LisgoTypeException = 8
)
