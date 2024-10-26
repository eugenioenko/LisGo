package lisgo

type LisgoException struct {
	Value string
	DType int
}

func NewLisgoException(value string) *LisgoException {
	return &LisgoException{Value: value, DType: LisgoTypeException}
}
func (data *LisgoException) ToString() string {
	return data.Value
}

func (data *LisgoException) ToBoolean() bool {
	return true
}

func (data *LisgoException) ToInteger() int64 {
	return 0
}

func (data *LisgoException) ToFloat() float64 {
	return 0.0
}

func (data *LisgoException) GetType() int {
	return data.DType
}

func (data *LisgoException) GetTypeName() string {
	return "exception"
}

func (data *LisgoException) GetValue() interface{} {
	return data.Value
}

func (data *LisgoException) Equals(other LisgoData) bool {
	return other.GetType() == data.DType && other.ToString() == data.Value
}
