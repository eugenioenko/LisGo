package lisgo

type LisgoReturn struct {
	From  string
	Value LisgoData
	DType int
}

func NewLisgoReturn(from string, value LisgoData) *LisgoReturn {
	return &LisgoReturn{DType: LisgoTypeReturn, From: from, Value: value}
}

func (data *LisgoReturn) ToString() string {
	return data.Value.ToString()
}

func (data *LisgoReturn) ToBoolean() bool {
	return data.Value.ToBoolean()
}

func (data *LisgoReturn) ToInteger() int64 {
	return data.Value.ToInteger()
}

func (data *LisgoReturn) ToFloat() float64 {
	return data.Value.ToFloat()
}

func (data *LisgoReturn) GetType() int {
	return data.DType
}

func (data *LisgoReturn) GetTypeName() string {
	return "return"
}

func (data *LisgoReturn) GetValue() interface{} {
	return data.Value
}

func (data *LisgoReturn) Equals(other LisgoData) bool {
	return other.GetType() == LisgoTypeFunction && data.GetValue() == other.GetValue()
}
