package lisgo

type LisgoFunction struct {
	name  string
	args  []string
	body  []Expression
	DType int
}

func NewLisgoFunction(name string, args []string, body []Expression) *LisgoFunction {
	return &LisgoFunction{DType: LisgoTypeFunction, name: name, args: args, body: body}
}

func (data *LisgoFunction) ToString() string {
	return data.name
}

func (data *LisgoFunction) ToBoolean() bool {
	return true
}

func (data *LisgoFunction) ToInteger() int64 {
	return 0
}

func (data *LisgoFunction) ToFloat() float64 {
	return 0
}

func (data *LisgoFunction) GetType() int {
	return data.DType
}

func (data *LisgoFunction) GetTypeName() string {
	return "function"
}

func (data *LisgoFunction) GetValue() interface{} {
	return data.body
}

func (data *LisgoFunction) Equals(other LisgoData) bool {
	return other.GetType() == LisgoTypeFunction && data.GetValue() == other.GetValue()
}
